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
	Name string
	UsersID []string
	Icon string
}

func (uc *SaveUseCase) Run(ctx context.Context, dto SaveUseCaseDto) error {
	// dtoからuserへ変換
	group, err := groupDomain.NewGroup(dto.Name, dto.UsersID, dto.Icon)
	if err != nil {
		return err
	}
	return uc.groupRepo.Save(ctx, group)
}