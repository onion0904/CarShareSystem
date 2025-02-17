package middleware

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// RecoverMiddleware はパニック発生時にエラーレスポンスを返す
func RecoverMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered from panic:", r) // ログ出力
				c.JSON(http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
			}
		}()
		return next(c)
	}
}
