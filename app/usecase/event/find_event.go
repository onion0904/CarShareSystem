package event

import (
	"context"
	eventDomain "github.com/onion0904/app/domain/event"
	"time"
)

type FindEventUseCase struct {
	eventRepo eventDomain.EventRepository
}

func NewFindEventUseCase(
	eventRepo eventDomain.EventRepository,
) *FindEventUseCase {
	return &FindEventUseCase{
		eventRepo: eventRepo,
	}
}

type FindEventUseCaseDto struct {
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

func (uc *FindEventUseCase) Run(ctx context.Context, eventID string) (*FindEventUseCaseDto, error) {
	event, err := uc.eventRepo.FindEvent(ctx, eventID)
	if err != nil {
		return nil, err
	}
	return &FindEventUseCaseDto{
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
