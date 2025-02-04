package user

import (
	"context"
	userDomain "github.com/onion0904/app/domain/user"
)

type UpdateUseCase struct {
	userRepo userDomain.UserRepository
}

func NewUpdateUserUseCase(
	userRepo userDomain.UserRepository,
) *UpdateUseCase {
	return &UpdateUseCase{
		userRepo: userRepo,
	}
}

type UpdateUseCaseDto struct {
	LastName string
	FirstName string
	Email string
	Icon string
}

func (uc *UpdateUseCase) Run(ctx context.Context, id string, dto UpdateUseCaseDto) error {
	// dtoからuserへ変換
	user ,err := uc.userRepo.FindUser(ctx,id)
	if err != nil {
        return err
    }
	nuser, err := userDomain.Reconstruct(id,dto.LastName, dto.FirstName, dto.Email, user.Password(), dto.Icon,user.GroupID())
	if err != nil {
		return err
	}
	return uc.userRepo.Update(ctx, nuser)
}