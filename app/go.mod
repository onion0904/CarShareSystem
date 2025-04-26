module github.com/onion0904/app

go 1.23.8

toolchain go1.24.2

require (
	github.com/99designs/gqlgen v0.17.72
	github.com/go-sql-driver/mysql v1.8.1
	github.com/golang-migrate/migrate/v4 v4.18.2
	github.com/mailgun/mailgun-go/v4 v4.22.1
	github.com/onion0904/go-pkg v0.0.0-00010101000000-000000000000
	github.com/skip2/go-qrcode v0.0.0-20200617195104-da1b6568686e
	github.com/vektah/gqlparser/v2 v2.5.25
	go.uber.org/mock v0.5.0
	golang.org/x/crypto v0.37.0
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/agnivade/levenshtein v1.2.1 // indirect
	github.com/go-chi/chi/v5 v5.2.0 // indirect
	github.com/go-viper/mapstructure/v2 v2.2.1 // indirect
	github.com/golang-jwt/jwt/v5 v5.2.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mailgun/errors v0.4.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/oklog/ulid/v2 v2.1.0 // indirect
	github.com/sosodev/duration v1.3.1 // indirect
	go.uber.org/atomic v1.7.0 // indirect
)

replace github.com/onion0904/go-pkg => ../pkg
