package group

import "context"

type GroupRepository interface {
	Save(ctx context.Context , userID string) error
	Delete(ctx context.Context , userID string) error
	FindAllID(ctx context.Context) (userID string,err error)
}