package calendar

import (
	"context"
	calendarDomain "github.com/onion0904/app/domain/calendar"
)

type DeleteEventUseCase struct {
	calendarRepo calendarDomain.CalendarRepository
}

func NewDeleteEventUseCase(
	calendarRepo calendarDomain.CalendarRepository,
) *DeleteEventUseCase {
	return &DeleteEventUseCase{
		calendarRepo: calendarRepo,
	}
}

func (uc *DeleteEventUseCase) Run(ctx context.Context, eventID string) error {
	err := uc.calendarRepo.DeleteEvent(ctx, eventID)
	if err != nil {
		return err
	}
	return nil
}