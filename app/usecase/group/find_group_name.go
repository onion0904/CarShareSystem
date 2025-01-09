package group

import (
	"context"
	groupDomain "github.com/onion0904/app/domain/group"
)

type FindGroupNameUseCase struct {
	groupRepo groupDomain.GroupRepository
}

func NewFindGroupNameUseCase(
	groupRepo groupDomain.GroupRepository,
) *FindGroupNameUseCase {
	return &FindGroupNameUseCase{
		groupRepo: groupRepo,
	}
}

type FindGroupNameUseCaseDto struct {
	Name string
}

func (uc *FindGroupNameUseCase) Run(ctx context.Context, groupID string) (*FindGroupNameUseCaseDto, error) {
	groupName, err := uc.groupRepo.FindGroupName(ctx, groupID)
	if err != nil {
		return nil, err
	}
	return &FindGroupNameUseCaseDto{
        Name: groupName,
    }, nil
}