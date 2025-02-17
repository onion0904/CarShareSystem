package graph

import "database/sql"


// Resolver はアプリケーションの依存関係を管理する
// Resolverのメソッドでsql.DBを使えるようにするため
type Resolver struct{
	DB *sql.DB
}
