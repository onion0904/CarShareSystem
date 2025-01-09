package calendar

import (
	"context"
	calendarDomain "github.com/onion0904/app/domain/calendar"
)

type FindMonthCalendarUseCase struct {
	calendarRepo calendarDomain.CalendarRepository
}

func NewFindMonthCalendarUseCase(
	calendarRepo calendarDomain.CalendarRepository,
) *FindMonthCalendarUseCase {
	return &FindMonthCalendarUseCase{
		calendarRepo: calendarRepo,
	}
}

type FindMonthCalendarUseCaseDto struct {
	eventIDs []string
}

func (uc *FindMonthCalendarUseCase) Run(ctx context.Context, year int, month int) (*FindMonthCalendarUseCaseDto, error) {
	eventIDs, err := uc.calendarRepo.FindMonthEventID(ctx, year, month)
	if err != nil {
		return nil, err
	}
	return &FindMonthCalendarUseCaseDto{
        eventIDs: eventIDs,
    }, nil
}
