package server

import (
	_ "log"
	_ "net/http"
	_ "os"
	"github.com/onion0904/app/config"
	"github.com/onion0904/app/infrastructure/db"
	mymiddleware "github.com/onion0904/app/middleware"
	"github.com/onion0904/go-pkg/jwt"
	
	"github.com/99designs/gqlgen/graphql/handler"
	_ "github.com/99designs/gqlgen/graphql/handler/extension"
	_ "github.com/99designs/gqlgen/graphql/handler/lru"
	_ "github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/onion0904/app/presentation/graphql/graph"
	_ "github.com/vektah/gqlparser/v2/ast"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echojwt "github.com/labstack/echo-jwt/v4"
)

const defaultPort = "8080"

func Start() {
	cfg := config.GetConfig()
	DB := db.NewMainDB(cfg.DB)

	// Echoのインスタンスを作成
	e := echo.New()

	// ミドルウェアの設定
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	publicGroup := e.Group("/public")
	privateGroup := e.Group("/private")

	// Recover ミドルウェア適用
	publicGroup.Use(mymiddleware.RecoverMiddleware)
	privateGroup.Use(mymiddleware.RecoverMiddleware)

    graphqlHandler := handler.NewDefaultServer(
		graph.NewExecutableSchema(
			graph.Config{Resolvers: &graph.Resolver{DB: DB}},
		),
	)
	playgroundHandler := playground.Handler("GraphQL", "/query")

	publicGroup.POST("/query", func(c echo.Context) error {
		graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	publicGroup.GET("/playground", func(c echo.Context) error {
		playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	secret := cfg.JWT.Secret
	jwtSecret := []byte(secret)
	echojwt_config := jwt.EchoJwtInit(jwtSecret)
	privateGroup.Use(echojwt.WithConfig(*echojwt_config))

	// GraphQLでは通常クエリのルートは統一して、認証の有無でアクセスを制御する
	privateGroup.POST("/query", func(c echo.Context) error {
		graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	privateGroup.GET("/playground", func(c echo.Context) error {
		playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.Logger.Fatal(e.Start(":" + defaultPort))
}
