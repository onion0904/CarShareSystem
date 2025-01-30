package repository

import (
	"context"
	"database/sql"
	"errors"

	errDomain "github.com/onion0904/app/domain/error"
    "github.com/onion0904/app/domain/event"
	"github.com/onion0904/app/infrastructure/db"
	dbgen "github.com/onion0904/app/infrastructure/db/sqlc/dbgen"
)

type eventRepository struct {}

func NewEventRepository() event.EventRepository {
	return &eventRepository{}
}

func (er *eventRepository)SaveEvent(ctx context.Context, event *event.Event) error {
	query := db.GetQuery(ctx)

	err := query.SaveEvent(ctx ,dbgen.SaveEventParams{
		ID:          event.ID(),
		UsersID:     event.UserID(),
		Together:    event.Together(),
		Description: event.Description(),
		Year:        event.Year(),
		Month:       event.Month(),
		Day:         event.Day(),
		StartDate:   event.StartDate(),
		EndDate:     event.EndDate(),
	})
	if err!= nil {
		return err
	}
	return nil
}
	
func (er *eventRepository)DeleteEvent(ctx context.Context , eventID string) error {
	query := db.GetQuery(ctx)

	err := query.DeleteEvent(ctx, eventID)
	if err!= nil {
        return err
    }
	return nil
}
	
func (er *eventRepository)FindEvent(ctx context.Context, eventID string) (*event.Event, error) {
	query := db.GetQuery(ctx)

	e, err := query.FindEvent(ctx, eventID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errDomain.NewError("Event not found")
		}
		return nil, err
	}

	ne, err := event.Reconstruct(
		e.ID,
        e.UsersID,
        e.Together,
        e.Description,
        e.Year,
        e.Month,
        e.Day,
		e.Date,
        e.StartDate,
        e.EndDate,
		e.Important,
	)
	if err != nil {
		return nil, err
	}
	return ne, nil
}
	
func (er *eventRepository)FindMonthEventID(ctx context.Context, year int32, month int32) ([]string, error) {
	query := db.GetQuery(ctx)

	eventIDs, err := query.FindMonthEventID(ctx, dbgen.FindMonthEventIDParams{
		Year:  year,
        Month: month,
	})
	if err!= nil {
		return nil, err
	}
	return eventIDs, nil
}