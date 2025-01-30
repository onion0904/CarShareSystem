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
	lastName string
	firstName string
	email string
	password string
	icon string
}

func (uc *SaveUseCase) Run(ctx context.Context, dto SaveUseCaseDto) error {
	// dtoからuserへ変換
	user, err := userDomain.NewUser(dto.lastName, dto.firstName, dto.email, dto.password, dto.icon)
	if err != nil {
		return err
	}
	return uc.userRepo.Save(ctx, user)
}