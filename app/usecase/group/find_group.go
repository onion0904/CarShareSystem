package group

import (
	"context"
	groupDomain "github.com/onion0904/app/domain/group"
)

type FindGroupUseCase struct {
	groupRepo groupDomain.GroupRepository
}

func NewFindGroupUseCase(
	groupRepo groupDomain.GroupRepository,
) *FindGroupUseCase {
	return &FindGroupUseCase{
		groupRepo: groupRepo,
	}
}

type FindGroupUseCaseDto struct {
	id string
	name string
	usersID []string
	icon string
}

func (uc *FindGroupUseCase) Run(ctx context.Context, groupID string) (*FindGroupUseCaseDto, error) {
	group, err := uc.groupRepo.FindGroup(ctx, groupID)
	if err != nil {
		return nil, err
	}
	return &FindGroupUseCaseDto{
		id:          group.ID(),
		name:        group.Name(),
		usersID:     group.UsersID(),
		icon:        group.Icon(),
	}, nil
}