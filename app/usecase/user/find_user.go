package user

import (
	"context"
	userDomain "github.com/onion0904/app/domain/user"
)

type FindUserUseCase struct {
	userRepo userDomain.UserRepository
}

func NewFindUserUseCase(
	userRepo userDomain.UserRepository,
) *FindUserUseCase {
	return &FindUserUseCase{
		userRepo: userRepo,
	}
}

type FindUserUseCaseDto struct {
	ID          string
	LastName    string
	FirstName   string
	Email       string
	Icon        string
	groupID      []string
}

func (uc *FindUserUseCase) Run(ctx context.Context, id string) (*FindUserUseCaseDto, error) {
	user, err := uc.userRepo.FindUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return &FindUserUseCaseDto{
		ID:          user.ID(),
		LastName:    user.LastName(),
		FirstName:   user.FirstName(),
		Email:       user.Email(),
		Icon:        user.Icon(),
		groupID:      user.GroupID(),
	}, nil
}