package server

import (
	"log"
	"net/http"
	"github.com/onion0904/app/config"
	"github.com/onion0904/app/infrastructure/db"
	mymiddleware "github.com/onion0904/app/middleware"
	// "github.com/onion0904/go-pkg/jwt"
	
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/onion0904/app/presentation/graphql/graph"
	// "github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
	// echojwt "github.com/labstack/echo-jwt/v4"
)

// func Start() {
// 	cfg := config.GetConfig()
// 	DB := db.NewMainDB(cfg.DB)

// 	Port := cfg.Server.Port

// 	// Echoのインスタンスを作成
// 	e := echo.New()

// 	// ミドルウェアの設定
// 	e.Use(middleware.Logger())
// 	e.Use(middleware.Recover())
	
// 	schema := graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{DB: DB}})
// 	if schema == nil {
// 		log.Fatal("スキーマ生成に失敗しました")
// 	}
// 	graphqlHandler := handler.New(schema)


// 	publicGroup := e.Group("/public")
// 	privateGroup := e.Group("/private")

// 	// Recover ミドルウェア適用
// 	publicGroup.Use(mymiddleware.RecoverMiddleware)
// 	privateGroup.Use(mymiddleware.RecoverMiddleware)

// 	playgroundHandler := playground.Handler("GraphQL", "/query")

// 	publicGroup.POST("/query", func(c echo.Context) error {
// 		graphqlHandler.ServeHTTP(c.Response(), c.Request())
// 		return nil
// 	})

// 	publicGroup.GET("/playground", func(c echo.Context) error {
// 		playgroundHandler.ServeHTTP(c.Response(), c.Request())
// 		return nil
// 	})

// 	secret := cfg.JWT.Secret
// 	jwtSecret := []byte(secret)
// 	echojwt_config := jwt.EchoJwtInit(jwtSecret)
// 	privateGroup.Use(echojwt.WithConfig(*echojwt_config))

// 	// GraphQLでは通常クエリのルートは統一して、認証の有無でアクセスを制御する
// 	privateGroup.POST("/query", func(c echo.Context) error {
// 		graphqlHandler.ServeHTTP(c.Response(), c.Request())
// 		return nil
// 	})

// 	privateGroup.GET("/playground", func(c echo.Context) error {
// 		playgroundHandler.ServeHTTP(c.Response(), c.Request())
// 		return nil
// 	})

// 	e.Logger.Fatal(e.Start(":" + Port))
// }

func Start() {
	cfg := config.GetConfig()
	DB := db.NewMainDB(cfg.DB)

	Port := cfg.Server.Port
	
	srv := handler.New(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{DB: DB},
		Directives: graph.Directive,
	}))

	// CORS対応。
	srv.AddTransport(transport.Options{})
	// GET / POST：GraphQLクエリを HTTP 経由で受け取るため。
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	
	// GraphQLクエリのキャッシュを有効化。LRU（最近使っていないものから削除）
	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))
	
	// GraphQLのスキーマをクエリで確認できる「introspection」機能をオンにしてる。GraphQL Playgroundなどでスキーマの構造が見えるのはこれのおかげ。
	srv.Use(extension.Introspection{})
	// クライアントが「ハッシュ化されたクエリ」を使って通信することをサポート（帯域の節約になる）。
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})
	
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", mymiddleware.AuthMiddleware(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", Port)
	log.Fatal(http.ListenAndServe(":"+Port, nil))
}