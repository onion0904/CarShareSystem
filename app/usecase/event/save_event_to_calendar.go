package event

import (
	"context"
	eventDomain "github.com/onion0904/app/domain/event"
)
// eventUsecase 構造体
type SaveEventUsecase struct {
	eventService eventDomain.EventDomainService
}

// NewCalendarUsecase ファクトリ関数
func NewEventUsecase(
	eventService eventDomain.EventDomainService,
) *SaveEventUsecase {
	return &SaveEventUsecase{
		eventService: eventService,
	}
}

// AddEventUseCaseDTO ユースケース層で使用する入力データ
type AddEventUseCaseDTO struct {
	UsersID string
	Together bool
	Description string
	Important bool
}


// AddEvent メソッド: イベントを追加する
func (uc *SaveEventUsecase) AddEvent(ctx context.Context, dto AddEventUseCaseDTO) error {	
	event, err := eventDomain.NewEvent(dto.UsersID, dto.Together, dto.Description, dto.Important)
	if err != nil {
		return err
	}
	// ドメイン層のサービスを呼び出し
	err = uc.eventService.SaveEventService(ctx, event)
	if err != nil {
		return err
	}

	return nil
}
