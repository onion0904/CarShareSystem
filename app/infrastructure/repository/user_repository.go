package repository

import (
	"context"
	"database/sql"
	"errors"

	errDomain "github.com/onion0904/app/domain/error"
    "github.com/onion0904/app/domain/user"
	"github.com/onion0904/app/infrastructure/db"
	dbgen "github.com/onion0904/app/infrastructure/db/sqlc/dbgen"
)

type userRepository struct {}

func NewUserRepository() user.UserRepository {
	return &userRepository{}
}

func (ur *userRepository)Update(ctx context.Context, user *user.User) error {
	query := db.GetQuery(ctx)
	if err := query.UpsertUser(ctx, dbgen.UpsertUserParams{
		ID:        user.ID(),
        LastName:  user.LastName(),
        FirstName: user.FirstName(),
        Email:     user.Email(),
		// iconがある場合urlを代入,ない場合はnullを代入代入
		Icon:      sql.NullString{String: user.Icon(), Valid: user.Icon()!= ""},
		}); err != nil {
		return err
	}
	return nil
}

func (ur *userRepository)Save(ctx context.Context, user *user.User) error {
	query := db.GetQuery(ctx)
	if err := query.UpsertUser(ctx, dbgen.UpsertUserParams{
		ID:        user.ID(),
        LastName:  user.LastName(),
        FirstName: user.FirstName(),
        Email:     user.Email(),
		// iconがある場合urlを代入,ない場合はnullを代入代入
		Icon:      sql.NullString{String: user.Icon(), Valid: user.Icon()!= ""},
		}); err != nil {
		return err
	}
	return nil
}
    
func (ur *userRepository)FindUser(ctx context.Context, UserID string) (*user.User, error) {
	query := db.GetQuery(ctx)

	u, err := query.FindUser(ctx, UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errDomain.NewError("User not found")
		}
		return nil, err
	}

	//Reconstructではstringしか使えないため
	icon := ""
	if u.Icon.Valid {
		icon = u.Icon.String
	}

	groupIDs, err := query.GetGroupIDsByUserID(ctx, UserID)
	if err != nil {
		return nil, err
	}

	eventIDs, err := query.GetEventIDsByUserID(ctx, UserID)
	if err!= nil {
        return nil, err
    }

	nu, err := user.Reconstruct(
		u.ID,
        u.LastName,
        u.FirstName,
        u.Email,
		u.Password,
        icon,
		groupIDs,
		eventIDs,
	)
	if err != nil {
		return nil, err
	}
	return nu, nil
}
	
func (ur *userRepository)Delete(ctx context.Context, UserID string) error {
	query := db.GetQuery(ctx)

	if err := query.DeleteUser(ctx, UserID); err!= nil {
        return err
    }
	return nil
}
    
func (ur *userRepository)ExistUser(ctx context.Context, email string, password string) (bool,error) {
	query := db.GetQuery(ctx)
	
	exist,err := query.ExistUser(ctx, dbgen.ExistUserParams{Email: email, Password: password})
	if err!= nil {
        return false, err
    }
	return exist, nil
}