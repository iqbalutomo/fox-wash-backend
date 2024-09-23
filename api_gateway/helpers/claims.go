package helpers

import (
	"api_gateway/dto"
	"api_gateway/utils"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func GetClaims(c echo.Context) (dto.Claims, error) {
	claimTmp := c.Get("user")
	if claimTmp == nil {
		return dto.Claims{}, echo.NewHTTPError(utils.ErrUnauthorized.EchoFormatDetails("Failed to fetch user claims from JWT"))
	}

	claims := claimTmp.(jwt.MapClaims)
	return dto.Claims{
		ID:    uint(claims["id"].(float64)),
		Name:  claims["name"].(string),
		Email: claims["email"].(string),
		Role:  claims["role"].(string),
	}, nil
}
