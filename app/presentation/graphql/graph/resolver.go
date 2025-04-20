package graph

import (
	"database/sql"
	usecase_mail "github.com/onion0904/app/usecase/mail"
)


// Resolver はアプリケーションの依存関係を管理する
// Resolverのメソッドでsql.DBを使えるようにするため
type Resolver struct{
	DB *sql.DB
	EmailUseCase  *usecase_mail.SendEmailUseCase
}
