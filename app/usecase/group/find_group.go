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
	Id string
	Name string
	UsersID []string
	Icon string
}

func (uc *FindGroupUseCase) Run(ctx context.Context, groupID string) (*FindGroupUseCaseDto, error) {
	group, err := uc.groupRepo.FindGroup(ctx, groupID)
	if err != nil {
		return nil, err
	}
	return &FindGroupUseCaseDto{
		Id:          group.ID(),
		Name:        group.Name(),
		UsersID:     group.UsersID(),
		Icon:        group.Icon(),
	}, nil
}