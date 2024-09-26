package controllers

import (
	"api_gateway/dto"
	"api_gateway/helpers"
	"api_gateway/models"
	"api_gateway/pb/userpb"
	"api_gateway/utils"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserController struct {
	client userpb.UserClient
}

func NewUserController(client userpb.UserClient) *UserController {
	return &UserController{client}
}

func (u *UserController) UserRegister(c echo.Context) error {
	return u.Register(c, utils.UserRoleID, utils.UserRole)
}

func (u *UserController) WasherRegister(c echo.Context) error {
	return u.Register(c, utils.WasherRoleID, utils.WasherRole)
}

func (u *UserController) AdminRegister(c echo.Context) error {
	return u.Register(c, utils.AdminRoleID, utils.AdminRole)
}

func (u *UserController) Register(c echo.Context, roleID uint, roleName string) error {
	register := new(dto.UserRegister)
	if err := c.Bind(register); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}

	if err := c.Validate(register); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(register.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}

	registerData := &userpb.RegisterRequest{
		FirstName: register.FirstName,
		LastName:  register.LastName,
		Email:     register.Email,
		Password:  string(hashedPassword),
		RoleId:    uint32(roleID),
	}

	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return err
	}
	defer cancel()

	responseGrpc, err := u.client.Register(ctx, registerData)
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.AlreadyExists:
				return echo.NewHTTPError(utils.ErrConflict.EchoFormatDetails(err.Error()))
			case codes.Internal:
				return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
			}
		}

		return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
	}

	if roleName == utils.WasherRole {
		ctx, cancel, err := helpers.NewServiceContext()
		if err != nil {
			return err
		}
		defer cancel()

		if _, err := u.client.CreateWasher(ctx, &userpb.WasherID{Id: responseGrpc.UserId}); err != nil {
			return utils.AssertGrpcStatus(err)
		}
	}

	responseData := models.User{
		ID:        uint(responseGrpc.UserId),
		FirstName: register.FirstName,
		LastName:  register.LastName,
		Email:     register.Email,
		Role:      roleName,
		CreatedAt: responseGrpc.CreatedAt,
	}

	return c.JSON(http.StatusCreated, dto.Response{
		Message: "Registered successfully",
		Data:    responseData,
	})
}

func (u *UserController) VerifyUser(c echo.Context) error {
	token := c.Param("token")
	userIdTmp := c.Param("user_id")
	userID, err := strconv.Atoi(userIdTmp)
	if err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}

	pbUserData := &userpb.UserCredential{
		Id:    uint32(userID),
		Token: token,
	}

	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return err
	}
	defer cancel()

	if _, err := u.client.VerifyNewUser(ctx, pbUserData); err != nil {
		return utils.AssertGrpcStatus(err)
	}

	return c.Redirect(http.StatusTemporaryRedirect, "/users/verified")
}

func (u *UserController) Login(c echo.Context) error {
	loginReq := new(dto.UserLogin)
	if err := c.Bind(loginReq); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}

	if err := c.Validate(loginReq); err != nil {
		return echo.NewHTTPError(utils.ErrBadRequest.EchoFormatDetails(err.Error()))
	}

	emailRequest := &userpb.EmailRequest{
		Email: loginReq.Email,
	}

	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return err
	}
	defer cancel()

	userDataTmp, err := u.client.GetUser(ctx, emailRequest)
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				return echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Invalid username/password"))
			default:
				return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(e.Message()))
			}
		}
	}

	if !userDataTmp.IsVerified {
		return echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Please do an email verfication"))
	}

	userData := models.User{
		ID:        uint(userDataTmp.Id),
		FirstName: userDataTmp.FirstName,
		LastName:  userDataTmp.LastName,
		Email:     userDataTmp.Email,
		Password:  userDataTmp.Password,
		Role:      userDataTmp.Role,
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(loginReq.Password)); err != nil {
		return echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Invalid username/password"))
	}

	if userData.Role == utils.WasherRole {
		ctx, cancel, err := helpers.NewServiceContext()
		if err != nil {
			return err
		}
		defer cancel()

		washerData, err := u.client.GetWasher(ctx, &userpb.WasherID{Id: uint32(userData.ID)})
		if err != nil {
			return echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Invalid username/password"))
		}

		if !washerData.IsActive {
			return echo.NewHTTPError(utils.ErrForbidden.EchoFormatDetails("Your account is still being reviewed by Fox Wash Team"))
		}

		if _, err := u.client.SetWasherStatusOnline(ctx, &userpb.WasherID{Id: washerData.UserId}); err != nil {
			return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
		}
	}

	if err := helpers.SignNewJWT(c, userData); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: "Login successfully",
		Data:    "Authorization is stored in cookie",
	})
}

func (u *UserController) WasherActivation(c echo.Context) error {
	user, err := helpers.GetClaims(c)
	if err != nil {
		return err
	}

	if user.Role != utils.AdminRole {
		return echo.NewHTTPError(utils.ErrForbidden.EchoFormatDetails("Access permission"))
	}

	email := c.Param("email")

	ctx, cancel, err := helpers.NewServiceContext()
	if err != nil {
		return err
	}
	defer cancel()

	if _, err := u.client.WasherActivation(ctx, &userpb.EmailRequest{Email: email}); err != nil {
		return utils.AssertGrpcStatus(err)
	}

	return c.JSON(http.StatusOK, dto.Response{
		Message: "Washer has been activated!",
		Data:    email,
	})
}

func (u *UserController) Logout(c echo.Context) error {
	user, err := helpers.GetClaims(c)
	fmt.Println(user)
	if err != nil {
		return err
	}

	if user.Role == utils.WasherRole {
		ctx, cancel, err := helpers.NewServiceContext()
		if err != nil {
			return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
		}
		defer cancel()

		washerData, err := u.client.GetWasher(ctx, &userpb.WasherID{Id: uint32(user.ID)})
		if err != nil {
			return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
		}
		if _, err := u.client.SetWasherStatusOffline(ctx, &userpb.WasherID{Id: washerData.UserId}); err != nil {
			return echo.NewHTTPError(utils.ErrInternalServer.EchoFormatDetails(err.Error()))
		}
	}

	c.SetCookie(&http.Cookie{
		Name:     "Authorization",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0)})

	return c.JSON(http.StatusOK, dto.Response{
		Message: "Logout successfully",
		Data:    "Authorization in cookie is deleted",
	})
}
