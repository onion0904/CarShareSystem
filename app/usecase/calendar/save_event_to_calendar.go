package calendar

import (
	"context"
	calendarDomain "github.com/onion0904/app/domain/calendar"
)
// calendarUsecase 構造体
type SaveCalendarUsecase struct {
	calendarService calendarDomain.CalendarDomainService
}

// NewCalendarUsecase ファクトリ関数
func NewCalendarUsecase(
	calendarService calendarDomain.CalendarDomainService,
) *SaveCalendarUsecase {
	return &SaveCalendarUsecase{
		calendarService: calendarService,
	}
}

// AddEventUseCaseDTO ユースケース層で使用する入力データ
type AddEventUseCaseDTO struct {
	usersID string
	together bool
	description string
	important bool
}


// AddEvent メソッド: イベントを追加する
func (uc *SaveCalendarUsecase) AddEvent(ctx context.Context, dto AddEventUseCaseDTO) error {	
	event, err := calendarDomain.NewEvent(dto.usersID, dto.together, dto.description, dto.important)
	if err != nil {
		return err
	}
	// ドメイン層のサービスを呼び出し
	err = uc.calendarService.AddEventToCalendar(ctx, event)
	if err != nil {
		return err
	}

	return nil
}
