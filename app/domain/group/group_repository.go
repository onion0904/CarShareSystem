package group

import "context"

type GroupRepository interface {
	Update(ctx context.Context, group *Group) error
	Save(ctx context.Context ,group *Group) error
	AddGroupIDToUser(ctx context.Context ,groupID string, userID string) error
	Delete(ctx context.Context , groupID string) error
	FindAllUserID(ctx context.Context,groupID string) (userIDs []string,err error)
	FindGroup(ctx context.Context,groupID string) (group *Group,err error)
	FindGroupName(ctx context.Context,groupID string) (groupName string,err error)
}