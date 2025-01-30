package calendar

import "context"


type CalendarRepository interface {
	//CalendarDomainService経由で使用してください (domain/calendar/calendar_domain_service.go)
	SaveEvent(ctx context.Context, event *Event) error
	//以下はCalendarDomainService経由でなくてOKです
	DeleteEvent(ctx context.Context , eventID string) error
	FindEvent(ctx context.Context, eventID string) (*Event, error)
	FindMonthEventID(ctx context.Context, year int32, month int32) (eventID []string,err error)
}