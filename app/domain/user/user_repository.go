package user

import "context"

type UserRepository interface {
	Update(ctx context.Context, user *User) error
    Save(ctx context.Context, user *User) error
    FindUser(ctx context.Context, UserID string) (*User, error)
	FindUserName(ctx context.Context, UserID string) (string, error)
	Delete(ctx context.Context, UserID string) error
    ExistUser(ctx context.Context,email string, password string) (bool,error)
	FindAllGroupID(ctx context.Context,UserID string) (groupID []string,err error)
}