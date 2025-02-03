package user

import (
	"context"
	userDomain "github.com/onion0904/app/domain/user"
)

type FindGroupUseCase struct {
	userRepo userDomain.UserRepository
}

func NewFindGroupUseCase(
	userRepo userDomain.UserRepository,
) *FindGroupUseCase {
	return &FindGroupUseCase{
		userRepo: userRepo,
	}
}

type FindGroupUseCaseDto struct {
	GroupID []string
}

func (uc *FindGroupUseCase) Run(ctx context.Context, UserID string) (*FindGroupUseCaseDto, error) {
	groupID, err := uc.userRepo.FindAllGroupID(ctx, UserID)
	if err != nil {
		return nil, err
	}
	return &FindGroupUseCaseDto{
		GroupID: groupID,
    }, nil
}