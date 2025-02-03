package user

import (
	"context"
	userDomain "github.com/onion0904/app/domain/user"
)

type SaveUseCase struct {
	userRepo userDomain.UserRepository
}

func NewSaveUserUseCase(
	userRepo userDomain.UserRepository,
) *SaveUseCase {
	return &SaveUseCase{
		userRepo: userRepo,
	}
}

type SaveUseCaseDto struct {
	LastName string
	FirstName string
	Email string
	Password string
	Icon string
}

func (uc *SaveUseCase) Run(ctx context.Context, dto SaveUseCaseDto) error {
	// dtoからuserへ変換
	user, err := userDomain.NewUser(dto.LastName, dto.FirstName, dto.Email, dto.Password, dto.Icon)
	if err != nil {
		return err
	}
	return uc.userRepo.Save(ctx, user)
}