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

// @Summary     Register a new user
// @Description Register a new user with the role 'user'
// @Tags        customer
// @Accept      json
// @Produce     json
// @Param       request body dto.UserRegister true "User registration details"
// @Success     201 {object} dto.SwaggerResponseRegister
// @Failure     400 {object} utils.ErrResponse
// @Failure     409 {object} utils.ErrResponse
// @Failure     500 {object} utils.ErrResponse
// @Router      /users/register/user [post]
func (u *UserController) UserRegister(c echo.Context) error {
	return u.Register(c, utils.UserRoleID, utils.UserRole)
}

// @Summary     Register a new washer
// @Description Register a new user with the role 'washer'
// @Tags        washer
// @Accept      json
// @Produce     json
// @Param       request body dto.UserRegister true "Washer registration details"
// @Success     201 {object} dto.SwaggerResponseRegister
// @Failure     400 {object} utils.ErrResponse
// @Failure     409 {object} utils.ErrResponse
// @Failure     500 {object} utils.ErrResponse
// @Router      /users/register/washer [post]
func (u *UserController) WasherRegister(c echo.Context) error {
	return u.Register(c, utils.WasherRoleID, utils.WasherRole)
}

// @Summary     Register a new admin
// @Description Register a new user with the role 'admin'
// @Tags        admin
// @Accept      json
// @Produce     json
// @Param       request body dto.UserRegister true "Admin registration details"
// @Success     201 {object} dto.SwaggerResponseRegister
// @Failure     400 {object} utils.ErrResponse
// @Failure     409 {object} utils.ErrResponse
// @Failure     500 {object} utils.ErrResponse
// @Router      /users/register/admin [post]
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

// @Summary 	Verify user credentials
// @Description Verify the user registration using unique token sent to the registered email
// @Tags        all user
// @Accept 		json
// @Produce 	json
// @Param 		userid path integer true "User ID"
// @Param 		token path string true "Verification token"
// @Success 	200 {object} dto.Response
// @Failure 	400 {object} utils.ErrResponse
// @Failure 	500 {object} utils.ErrResponse
// @Router 		/users/verify/{userid}/{token} [get]
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

// @Summary 	Log in
// @Description Login users and embeds a JWt-Auth in cookie
// @Tags        all user
// @Accept 		json
// @Produce 	json
// @Param 		request body dto.UserLogin true "User login details"
// @Success     200  {object}  dto.Response
// @Failure     400  {object}  utils.ErrResponse
// @Failure     401  {object}  utils.ErrResponse
// @Failure     500  {object}  utils.ErrResponse
// @Router      /users/login [post]
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

// @Summary 	Washer Activation
// @Description Activation washer by admin for hired washer as team
// @Tags        admin
// @Accept 		json
// @Produce 	json
// @Param 		email path string true "Email Washer"
// @Success     200  {object}  dto.Response
// @Failure     400  {object}  utils.ErrResponse
// @Failure     403  {object}  utils.ErrResponse
// @Failure     500  {object}  utils.ErrResponse
// @Router      /admins/washer-activation/{email} [post]
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

// @Summary 	Logout the user
// @Description Logout the currently authenticated user and clears the authorization cookie
// @Tags        all user
// @Accept 		json
// @Produce 	json
// @Success 	200 {object} dto.Response
// @Failure 	500 {object} utils.ErrResponse
// @Router 		/users/logout [get]
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
