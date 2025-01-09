package group

import (
	"context"
	groupDomain "github.com/onion0904/app/domain/group"
)

type SaveUseCase struct {
	groupRepo groupDomain.GroupRepository
}

func NewSaveUserUseCase(
	groupRepo groupDomain.GroupRepository,
) *SaveUseCase {
	return &SaveUseCase{
		groupRepo: groupRepo,
	}
}

type SaveUseCaseDto struct {
	name string
	usersID []string
	icon string
}

func (uc *SaveUseCase) Run(ctx context.Context, dto SaveUseCaseDto) error {
	// dtoからuserへ変換
	group, err := groupDomain.NewGroup(dto.name, dto.usersID, dto.icon)
	if err != nil {
		return err
	}
	return uc.groupRepo.Save(ctx, group)
}