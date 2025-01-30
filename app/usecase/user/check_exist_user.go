package user

import (
	"context"
	userDomain "github.com/onion0904/app/domain/user"
)

type CheckExistUserUseCase struct {
	userRepo userDomain.UserRepository
}

func NewCheckExistUserUseCase(
	userRepo userDomain.UserRepository,
) *CheckExistUserUseCase {
	return &CheckExistUserUseCase{
		userRepo: userRepo,
	}
}

type CheckExistUserUseCaseDto struct {
	exist bool
}

func (uc *CheckExistUserUseCase) Run(ctx context.Context, email string, password string) (*CheckExistUserUseCaseDto,error) {
	exist,err := uc.userRepo.ExistUser(ctx,email, password)
	if err != nil {
        return nil, err
    }
	return &CheckExistUserUseCaseDto{
		exist: exist,
	},nil
}