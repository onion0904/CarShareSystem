package graph

import (
	"github.com/onion0904/go-pkg/ulid"
	"time"
)

// Resolver はアプリケーションの依存関係を管理する
type Resolver struct{}

// スカラー型のマッピング

// ID (ULID) の生成関数
func (r *Resolver) ID() string {
	return ulid.NewUlid()
}

// DateTime の生成関数
func (r *Resolver) DateTime() time.Time {
	return time.Now()
}
