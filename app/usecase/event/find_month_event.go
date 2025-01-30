package event

import (
	"context"
	eventDomain "github.com/onion0904/app/domain/event"
)

type FindMonthEventUseCase struct {
	eventRepo eventDomain.EventRepository
}

func NewFindMonthEventUseCase(
	eventRepo eventDomain.EventRepository,
) *FindMonthEventUseCase {
	return &FindMonthEventUseCase{
		eventRepo: eventRepo,
	}
}

type FindMonthEventUseCaseDto struct {
	eventIDs []string
}

func (uc *FindMonthEventUseCase) Run(ctx context.Context, year int32, month int32) (*FindMonthEventUseCaseDto, error) {
	eventIDs, err := uc.eventRepo.FindMonthEventID(ctx, year, month)
	if err != nil {
		return nil, err
	}
	return &FindMonthEventUseCaseDto{
        eventIDs: eventIDs,
    }, nil
}
