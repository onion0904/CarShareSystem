package jwt

import(
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func EchoJwtInit (jwtSecret []byte) *echojwt.Config {
	return &echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(MyCustomClaims)
		},
		SigningKey: jwtSecret,
	}
}