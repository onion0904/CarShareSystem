package user

import (
	"context"
	"fmt"

	userDomain "github.com/onion0904/app/domain/user"
)

type FindUserNameUseCase struct {
	userRepo userDomain.UserRepository
}

func NewFindUserNameUseCase(
	userRepo userDomain.UserRepository,
) *FindUserNameUseCase {
	return &FindUserNameUseCase{
		userRepo: userRepo,
	}
}

type FindUserNameUseCaseDto struct {
	Name string
}

func (uc *FindUserNameUseCase) Run(ctx context.Context, id string) (*FindUserNameUseCaseDto, error) {
	user, err := uc.userRepo.FindUserName(ctx, id)
	if err != nil {
		return nil, err
	}
	Name := fmt.Sprintf(user.LastName()+user.FirstName())
	return &FindUserNameUseCaseDto{
        Name: Name,
    }, nil
}