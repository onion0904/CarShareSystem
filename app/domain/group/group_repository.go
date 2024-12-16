package group

import "context"

type GroupRepository interface {
	SaveGroup(ctx context.Context ,group *Group) error
	DeleteGroup(ctx context.Context , groupID string) error
	FindAllGroupUserID(ctx context.Context,groupID string) (usersID []string,err error)
	FindAllGroupName(ctx context.Context,groupID string) (groupName string,err error)
	FindGroupID(ctx context.Context,userID string) (groupID string,err error)
}