package event

import (
	"context"
	eventDomain "github.com/onion0904/app/domain/event"
)

type DeleteEventUseCase struct {
	eventRepo eventDomain.EventRepository
}

func NewDeleteEventUseCase(
	eventRepo eventDomain.EventRepository,
) *DeleteEventUseCase {
	return &DeleteEventUseCase{
		eventRepo: eventRepo,
	}
}

func (uc *DeleteEventUseCase) Run(ctx context.Context, eventID string) error {
	err := uc.eventRepo.DeleteEvent(ctx, eventID)
	if err != nil {
		return err
	}
	return nil
}