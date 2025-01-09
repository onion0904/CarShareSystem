package group

import (
	"context"
	groupDomain "github.com/onion0904/app/domain/group"
)

type FindUserUseCase struct {
	groupRepo groupDomain.GroupRepository
}

func NewFindUserUseCase(
	groupRepo groupDomain.GroupRepository,
) *FindUserUseCase {
	return &FindUserUseCase{
		groupRepo: groupRepo,
	}
}

type FindUserUseCaseDto struct {
	userID []string
}

func (uc *FindUserUseCase) Run(ctx context.Context, groupID string) (*FindUserUseCaseDto, error) {
	userID, err := uc.groupRepo.FindAllUserID(ctx, groupID)
	if err != nil {
		return nil, err
	}
	return &FindUserUseCaseDto{
		userID: userID,
    }, nil
}