package user

import (
	"net/mail"
	"unicode/utf8"
	"github.com/onion0904/go-pkg/ulid"
	errDomain "github.com/onion0904/app/domain/error"
)


type User struct {
	id string
	lastName string
	firstName string
	email string
	icon string
}

func Reconstruct(
	id string,
    lastName string,
    firstName string,
    email string,
	icon string,
) (*User, error) {
	return newUser(
		id,
        lastName,
        firstName,
        email,
		icon,
	)
}

func NewUser(
	lastName string,
	firstName string,
	email string,
	icon string,
) (*User, error) {
	return newUser(
		ulid.NewUlid(),
		lastName,
		firstName,
		email,
		icon,
	)
}

func newUser(
	id string,
    lastName string,
    firstName string,
    email string,
	icon string,
) (*User, error) {
	// 名前のバリデーション
	if utf8.RuneCountInString(lastName) < nameLengthMin || utf8.RuneCountInString(lastName) > nameLengthMax {
		return nil, errDomain.NewError("名前（姓）の値が不正です。")
	}
	if utf8.RuneCountInString(firstName) < nameLengthMin || utf8.RuneCountInString(firstName) > nameLengthMax {
		return nil, errDomain.NewError("名前（名）の値が不正です。")
	}

	// メールアドレスのバリデーション
	if _, err := mail.ParseAddress(email); err != nil {
		return nil, errDomain.NewError("メールアドレスの値が不正です。")
	}

    return &User{
        id:          id,
        lastName:     lastName,
        firstName:    firstName,
        email:        email,
		icon:         icon,
    }, nil
}

func (u *User) ID() string {
	return u.id
}

func (u *User) LastName() string {
	return u.lastName
}

func (u *User) FirstName() string {
	return u.firstName
}

func (u *User) Email() string {
	return u.email
}

func (u *User) Icon() string {
	return u.icon
}

const (
	nameLengthMin = 1
    nameLengthMax = 50
)