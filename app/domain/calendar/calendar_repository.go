package calendar

import "context"

type CalendarRepository interface {
	SaveEvent(ctx context.Context, event *Event) error//AddEventToCalendarで使用してください
	DeleteEvent(ctx context.Context , eventID string) error
	FindEvent(ctx context.Context, eventID string) (*Event, error)
	FindMonthEventID(ctx context.Context, year int, month int) (eventID []string,err error)
	FindDayEventID(ctx context.Context, year int, month int, day int) (eventID []string,err error)
}