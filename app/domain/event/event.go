package event

import(
	"unicode/utf8"
	"github.com/onion0904/go-pkg/ulid"
	"github.com/onion0904/go-pkg/ints"
	pkgTime "github.com/onion0904/go-pkg/time"
	errDomain "github.com/onion0904/app/domain/error"
	"time"
)


type Event struct {
	id string
	usersID string
	together bool
	description string
	year int32
	month int32
	day int32
	date time.Time
	startDate time.Time
	endDate time.Time
	important bool
}

func Reconstruct(
	id string,
	usersID string,
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
		usersID,
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
	usersID string,
	together bool,
	description string,
	important bool,
) (*Event, error) {
	return newEvent(
		ulid.NewUlid(),
		usersID,
		together,
		description,
		pkgTime.Year(),
        pkgTime.Month(),
		pkgTime.Day(),
		pkgTime.Now(),
		pkgTime.NextStartWeek(),
		pkgTime.NextEndWeek(),
		important,
	)
}

func newEvent(
	id string,
	usersID string,
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
	// ownerIDのバリデーション
	if !ulid.IsValid(id) {
		return nil, errDomain.NewError("オーナーIDの値が不正です。")
	}
	// 名前のバリデーション
	if utf8.RuneCountInString(description) < nameLengthMin || utf8.RuneCountInString(description) > nameLengthMax {
		return nil, errDomain.NewError("グループ名が不正です。")
	}

	if ints.Digit(year) != yearLength || ints.Digit(month) != monthLength || ints.Digit(day) != dayLength{
		return nil, errDomain.NewError("グループ名が不正です。")
	}


	return &Event{
		id:          id,
		usersID:        usersID,
		together:        together,
		description:   description,
		year:         year,
        month:        month,
		day:         day,
		date:         date,
		startDate:  startDate,
        endDate:    endDate,
		important: important,
	}, nil
}

func (c *Event) ID() string {
	return c.id
}

func (c *Event) UserID() string {
	return c.usersID
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
	nameLengthMin = 1
	nameLengthMax = 200
	yearLength = 4
	monthLength = 2
	dayLength = 2
)