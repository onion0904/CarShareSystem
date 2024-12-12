package user

import "context"

type UserRepository interface {
	//ユーザーの保存
    Save(ctx context.Context, user *User) error
	//IDでユーザーを取得
    FindById(ctx context.Context, id string) (*User, error)
	//IDでユーザーを削除
	Delete(ctx context.Context, id string) error
}