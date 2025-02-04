package repository

import (
	"context"
	"database/sql"
	"errors"

	errDomain "github.com/onion0904/app/domain/error"
    "github.com/onion0904/app/domain/group"
	"github.com/onion0904/app/infrastructure/db"
	dbgen "github.com/onion0904/app/infrastructure/db/sqlc/dbgen"
)

type groupRepository struct {}

func NewGroupRepository() group.GroupRepository {
	return &groupRepository{}
}

func (gr *groupRepository)Update(ctx context.Context, group *group.Group) error {
	query := db.GetQuery(ctx)

	
	err := query.UpsertGroup(ctx,dbgen.UpsertGroupParams{
		Name: group.Name(), 
		Icon: sql.NullString{String: group.Icon(), Valid: group.Icon()!= ""}, 
		ID: group.ID(),
	})
	if err!= nil {
        return err
    }
	return nil
}
	
func (gr *groupRepository)Save(ctx context.Context ,group *group.Group) error {
	query := db.GetQuery(ctx)

	err := query.UpsertGroup(ctx, dbgen.UpsertGroupParams{
        ID:   group.ID(),
        Name: group.Name(), 
        Icon: sql.NullString{String: group.Icon(), Valid: group.Icon()!= ""},
    })
	if err!= nil {
        return err
    }
	return nil
}
	
func (gr *groupRepository)Delete(ctx context.Context , groupID string) error {
	query := db.GetQuery(ctx)

	err := query.DeleteGroup(ctx, groupID)
	if err!= nil {
        return err
    }
	return nil
}
	
func (gr *groupRepository)FindAllUserID(ctx context.Context,groupID string) ([]string,error) {
	query := db.GetQuery(ctx)

	userIDs, err := query.FindAllUserID(ctx, groupID)
	if err!= nil {
        return nil, err
    }
	return userIDs, nil
}
	
func (gr *groupRepository)FindGroup(ctx context.Context,groupID string) (*group.Group, error) {
	query := db.GetQuery(ctx)
	

	g, err := query.FindGroup(ctx, groupID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errDomain.NewError("Group not found")
		}
		return nil, err
	}

	icon := ""
	if g.Icon.Valid {
		icon = g.Icon.String
	}

	userIDs, err := query.GetUserIDsByGroupID(ctx, groupID)
	if err != nil {
		return nil, err
	}

	eventIDs, err := query.GetEventIDsByGroupID(ctx, groupID)
	if err!= nil {
        return nil, err
    }

	ng, err := group.Reconstruct(
		g.ID,
		g.Name,
		userIDs,
		eventIDs,
		icon,
	)
	if err != nil {
		return nil, err
	}
	return ng, nil
}

func (gr *groupRepository)AddUserToGroup(ctx context.Context ,groupID string, userID string) error {
	query := db.GetQuery(ctx)

    err := query.AddUserToGroup(ctx, dbgen.AddUserToGroupParams{
        Groupid: groupID,
        Userid: userID,
    })
	if err!= nil {
        return err
    }
	return nil
}

func (gr *groupRepository)AddEventToGroup(ctx context.Context ,groupID string, eventID string) error {
	query := db.GetQuery(ctx)

    err := query.AddEventToGroup(ctx, dbgen.AddEventToGroupParams{
        Groupid: groupID,
        Eventid: eventID,
    })
	if err!= nil {
        return err
    }
	return nil
}

func (gr *groupRepository)RemoveUserFromGroup(ctx context.Context ,groupID string, userID string) error {
	query := db.GetQuery(ctx)

    err := query.RemoveUserFromGroup(ctx, dbgen.RemoveUserFromGroupParams{
        GroupID: groupID,
        UserID: userID,
    })
	if err!= nil {
        return err
    }
	return nil
}