package calendar

import (
	"context"
	calendarDomain "github.com/onion0904/app/domain/calendar"
	"time"
)

type FindCalendarUseCase struct {
	calendarRepo calendarDomain.CalendarRepository
}

func NewFindCalendarUseCase(
	calendarRepo calendarDomain.CalendarRepository,
) *FindCalendarUseCase {
	return &FindCalendarUseCase{
		calendarRepo: calendarRepo,
	}
}

type FindCalendarUseCaseDto struct {
	id          string
	usersID     string
	together    bool
	description string
	year        int32
	month       int32
	day         int32
	date        time.Time
	startDate   time.Time
	endDate     time.Time
	important   bool
}

func (uc *FindCalendarUseCase) Run(ctx context.Context, eventID string) (*FindCalendarUseCaseDto, error) {
	event, err := uc.calendarRepo.FindEvent(ctx, eventID)
	if err != nil {
		return nil, err
	}
	return &FindCalendarUseCaseDto{
		id:          event.ID(),
		usersID:     event.UserID(),
		together:    event.Together(),
		description: event.Description(),
		year:        event.Year(),
		month:       event.Month(),
		day:         event.Day(),
		date:        event.Date(),
		startDate:   event.StartDate(),
		endDate:     event.EndDate(),
		important:   event.Important(),
	}, nil
}
