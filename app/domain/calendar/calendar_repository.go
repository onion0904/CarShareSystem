package calendar

import "context"

type CalendarRepository interface {
	SaveCalendar(ctx context.Context , eventID string) error
	DeleteCalendar(ctx context.Context , eventID string) error
	FindMonthEventID(ctx context.Context, year int, month int) (eventID string,err error)
	FindAllUserID(ctx context.Context,eventID string) (userID string,err error)
}