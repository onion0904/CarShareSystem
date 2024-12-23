package calendar

import (
	"context"
	"errors"
	pkgTime "github.com/onion0904/go-pkg/time"
	"time"
)

type CalendarDomainService interface {
	AddEventToCalendar(ctx context.Context, event *Event) error
}

type calendarDomainService struct {
	CalendarRepo CalendarRepository
	startDate    time.Time
	endDate      time.Time
}

func NewCalendarDomainService(
	CalendarRepo CalendarRepository,
	startDate time.Time,
	endDate time.Time,
) CalendarDomainService {
	return &calendarDomainService{
		CalendarRepo: CalendarRepo,
		startDate:     pkgTime.NextStartWeek(),
		endDate:      pkgTime.NextEndWeek(),
	}
}

func ReloadCalendar (c *calendarDomainService) {
	c.startDate = pkgTime.NextStartWeek()
	c.endDate = pkgTime.NextEndWeek()
}

func (c *calendarDomainService) AddEventToCalendar(ctx context.Context, event *Event) error {
	// イベント期間の制約を確認
	if event.date.Before(c.startDate) || event.date.After(c.endDate){
		return errors.New("イベントが登録可能期間外です")
	}

	// 現在のイベント数を取得
	events, err := c.CalendarRepo.FindMonthEventID(ctx, event.year, event.month)
	if err != nil {
		return err
	}

	// イベント数の制約を確認
	if !c.validNumEvents(ctx,events) {
		return errors.New("イベントの最大数を超えています")
	}

	err = c.CalendarRepo.SaveEvent(ctx, event)
	if err != nil {
		return err
	}

	return nil
}

func (c *calendarDomainService) validNumEvents (ctx context.Context , eventIDs []string) bool {
	var importantEvent int
	var nimportantEvent int
	for _, eventID := range eventIDs {
		event, err := c.CalendarRepo.FindEvent(ctx, eventID)
		if err!= nil {
            return false
        }
		if event.important{
			importantEvent++
		}else{
			nimportantEvent++
		}
	}
	if importantEvent < MaxImportantEvents && (nimportantEvent+importantEvent) < MaxEvents{
		return true
	}else if importantEvent >= MaxImportantEvents && (nimportantEvent+importantEvent) >= MaxEvents{
		return false
	}else {
		return false
	}
}

const (
	MaxEvents int = 4
	MaxImportantEvents int = 2
)