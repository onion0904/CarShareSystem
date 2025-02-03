package group

import (
	"context"
	userDomain "github.com/onion0904/app/domain/user"
)

type AddGroupIDToUserUseCase struct {
	userRepo userDomain.UserRepository
}

func NewAddGroupIDToUserUseCase(
	userRepo userDomain.UserRepository,
) *AddGroupIDToUserUseCase {
	return &AddGroupIDToUserUseCase{
		userRepo: userRepo,
	}
}

type AddGroupIDToUserUseCaseDto struct {
	UserID  string
	GroupID string
}

//
func (uc *AddGroupIDToUserUseCase) Run(ctx context.Context, dto AddGroupIDToUserUseCaseDto) error {
	user, err := uc.userRepo.FindUser(ctx, dto.UserID)
	if err != nil {
        return err
    }
	groupIDs := user.GroupID()
	groupIDs = append(groupIDs, dto.GroupID)
	user, err = userDomain.Reconstruct(user.ID(), user.LastName(), user.FirstName(), user.Email(),user.Password(), user.Icon(), groupIDs)
	if err != nil {
		return err
	}
	return uc.userRepo.Update(ctx, user)
}