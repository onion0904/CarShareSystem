package calendar

import(
	"unicode/utf8"
	"github.com/onion0904/go-pkg/ulid"
	"github.com/onion0904/go-pkg/ints"
	"github.com/onion0904/go-pkg/time"
	errDomain "github.com/onion0904/app/domain/error"
)


type Calendar struct {
	id string
	usersID string
	together bool
	description string
	year int
	month int
	day int
}

func Reconstruct(
	id string,
	usersID string,
	together bool,
	description string,
	year int,
	month int,
	day int,
) (*Calendar, error) {
	return newCalendar(
		id,
		usersID,
		together,
		description,
		year,
        month,
		day,
	)
}

func NewProduct(
	id string,
	usersID string,
	together bool,
	description string,
	year int,
    month int,
	day int,
) (*Calendar, error) {
	return newCalendar(
		ulid.NewUlid(),
		usersID,
		together,
		description,
		time.Year(),
        time.Month(),
		time.Day(),
	)
}

func newCalendar(
	id string,
	usersID string,
	together bool,
	description string,
	year int,
    month int,
	day int,
) (*Calendar, error) {
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

	return &Calendar{
		id:          id,
		usersID:        usersID,
		together:        together,
		description:   description,
		year:         year,
        month:        month,
		day:         day,
	}, nil
}

func (c *Calendar) ID() string {
	return c.id
}

func (c *Calendar) UserID() string {
	return c.usersID
}

func (c *Calendar) Together() bool {
	return c.together
}

func (c *Calendar) Description() string {
	return c.description
}

const (
	nameLengthMin = 1
	nameLengthMax = 200
	yearLength = 4
	monthLength = 2
	dayLength = 2
)