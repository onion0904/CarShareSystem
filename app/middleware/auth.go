// emailを使うやつの実装を加える。カスタムクレームもsegon_pixを参考に変更

package middleware

import (
    "context"
    "net/http"
    "os"
    "strings"

    "github.com/golang-jwt/jwt/v5"
    "github.com/labstack/echo/v4"
)

// 認証情報を格納するキー
type contextKey string

const UserContextKey contextKey = "authUser"

// JWT のカスタムクレーム
type MyCustomClaims struct {
    ID    string `json:"id"`
    Email string `json:"email"`
    jwt.RegisteredClaims
}

// 認証ミドルウェア
func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        authHeader := c.Request().Header.Get("Authorization")
        if authHeader == "" {
            return next(c) // トークンがない場合は認証なしのまま次へ
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        secret := os.Getenv("JWT_SECRET_KEY")

        token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
            return []byte(secret), nil
        })

        if err != nil || !token.Valid {
            return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
        }

        claims, ok := token.Claims.(*MyCustomClaims)
        if !ok {
            return echo.NewHTTPError(http.StatusUnauthorized, "Invalid claims")
        }

        // 認証情報を Context に格納
        ctx := context.WithValue(c.Request().Context(), UserContextKey, claims)
        c.SetRequest(c.Request().WithContext(ctx))

        return next(c)
    }
}

// Context から認証ユーザーを取得
func GetAuthUser(ctx context.Context) *MyCustomClaims {
    authUser, ok := ctx.Value(UserContextKey).(*MyCustomClaims)
    if !ok {
        return nil
    }
    return authUser
}
