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

func (uc *SaveUseCase) Run(ctx context.Context, dto SaveUseCaseDto) (*FindUserUseCaseDto,error) {
	// dtoからuserへ変換
	nuser, err := userDomain.NewUser(dto.LastName, dto.FirstName, dto.Email, dto.Password, dto.Icon)
	if err != nil {
		return nil,err
	}
	err = uc.userRepo.Save(ctx, nuser)
	if err != nil {
		return nil,err
	}
	user,err := uc.userRepo.FindUser(ctx,nuser.ID())
	if err != nil {
		return nil,err
	}
	return &FindUserUseCaseDto{
		ID:          user.ID(),
		LastName:    user.LastName(),
		FirstName:   user.FirstName(),
		Email:       user.Email(),
		Password:    user.Password(),
		Icon:        user.Icon(),
		GroupIDs:    user.GroupIDs(),
		EventIDs:    user.EventIDs(),
		CreatedAt:   user.CreatedAt(),
        UpdatedAt:   user.UpdatedAt(),
	}, nil
}