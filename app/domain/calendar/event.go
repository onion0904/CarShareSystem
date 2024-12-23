package calendar

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
	year int
	month int
	day int
	date time.Time
	important bool
}

func Reconstruct(
	id string,
	usersID string,
	together bool,
	description string,
	year int,
	month int,
	day int,
	date time.Time,
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
		important,
	)
}

func NewEvent(
	id string,
	usersID string,
	together bool,
	description string,
	year int,
    month int,
	day int,
	date time.Time,
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
		important,
	)
}

func newEvent(
	id string,
	usersID string,
	together bool,
	description string,
	year int,
    month int,
	day int,
	date time.Time,
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

func (c *Event) Year() int {
    return c.year
}

func (c *Event) Month() int {
    return c.month
}

func (c *Event) Day() int {
    return c.day
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