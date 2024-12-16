package user

import "context"

type UserRepository interface {
    Save(ctx context.Context, user *User) error
    FindUser(ctx context.Context, id string) (*User, error)
	FindUserName(ctx context.Context, id string) (string, error)
    FindUserIcon(ctx context.Context, id string) (string, error)
	Delete(ctx context.Context, id string) error
    ExistUser(email string, password string) bool
}