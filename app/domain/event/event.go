package event

import (
	errDomain "github.com/onion0904/app/domain/error"
	pkgTime "github.com/onion0904/go-pkg/time"
	"github.com/onion0904/go-pkg/ulid"
	"time"
	"unicode/utf8"
)

type Event struct {
	id          string
	userID      string
	together    bool
	description string
	year        int32
	month       int32
	day         int32
	date        time.Time
	createdAt   time.Time
	updatedAt   time.Time
	startDate   time.Time
	endDate     time.Time
	important   bool
}

func Reconstruct(
	id string,
	userID string,
	together bool,
	description string,
	year int32,
	month int32,
	day int32,
	date time.Time,
	startDate time.Time,
	endDate time.Time,
	important bool,
) (*Event, error) {
	return newEvent(
		id,
		userID,
		together,
		description,
		year,
		month,
		day,
		date,
		startDate,
		endDate,
		important,
	)
}

func NewEvent(
	userID string,
	together bool,
	description string,
	year int32,
    month int32,
    day int32,
	important bool,
) (*Event, error) {
	return newEvent(
		ulid.NewUlid(),
		userID,
		together,
		description,
		year,
        month,
        day,
		pkgTime.Now(),
		pkgTime.NextStartWeek(),
		pkgTime.NextEndWeek(),
		important,
	)
}

func newEvent(
	id string,
	userID string,
	together bool,
	description string,
	year int32,
	month int32,
	day int32,
	date time.Time,
	startDate time.Time,
	endDate time.Time,
	important bool,
) (*Event, error) {
	// IDのバリデーション
	if !ulid.IsValid(id) {
		return nil, errDomain.NewError("IDの値が不正です。")
	}
	// descriptionのバリデーション
	if utf8.RuneCountInString(description) < descriptionLengthMin || utf8.RuneCountInString(description) > descriptionLengthMax {
		return nil, errDomain.NewError("descriptionが不正です。")
	}

	// 値の範囲チェック
	if year < 1000 || year > 9999 || month < 1 || month > 12 || day < 1 || day > 31 {
		return nil, errDomain.NewError("year,month,dayの値が範囲外です")
	}

	// 月ごとの日数チェック（より厳密にする場合）
	var daysInMonth int32 = 31
	if month == 4 || month == 6 || month == 9 || month == 11 {
		daysInMonth = 30
	} else if month == 2 {
		// 閏年チェック
		if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
			daysInMonth = 29
		} else {
			daysInMonth = 28
		}
	}

	if day > daysInMonth {
		return nil, errDomain.NewError("指定された月の日数が不正です")
	}

	return &Event{
		id:          id,
		userID:      userID,
		together:    together,
		description: description,
		year:        year,
		month:       month,
		day:         day,
		date:        date,
		startDate:   startDate,
		endDate:     endDate,
		important:   important,
	}, nil
}

func (c *Event) ID() string {
	return c.id
}

func (c *Event) UserID() string {
	return c.userID
}

func (c *Event) Together() bool {
	return c.together
}

func (c *Event) Description() string {
	return c.description
}

func (c *Event) Year() int32 {
	return c.year
}

func (c *Event) Month() int32 {
	return c.month
}

func (c *Event) Day() int32 {
	return c.day
}

func (c *Event) Date() time.Time {
	return c.date
}

func (c *Event) CreatedAt() time.Time {
	return c.createdAt
}

func (c *Event) UpdatedAt() time.Time {
	return c.updatedAt
}

func (u *Event) SetCreatedAt(t time.Time) {
	u.createdAt = t
}

func (u *Event) SetUpdatedAt(t time.Time) {
	u.updatedAt = t
}

func (c *Event) StartDate() time.Time {
	return c.startDate
}

func (c *Event) EndDate() time.Time {
	return c.endDate
}

func (c *Event) Important() bool {
	return c.important
}

const (
	descriptionLengthMin = 1
	descriptionLengthMax = 200
)
