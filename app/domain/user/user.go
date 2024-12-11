package user

import (
	"net/mail"
	"unicode/utf8"
	"github.com/onion0904/go-pkg/ulid"
	errDomain "github.com/onion0904/app/domain/error"
)


type User struct {
	id string
	email string
	lastName string
	firstName string
	emailAddress string
}

func Reconstruct(
	id string,
    email string,
    lastName string,
    firstName string,
    emailAddress string,
) (*User, error) {
	return newUser(
		id,
        email,
        lastName,
        firstName,
        emailAddress,
	)
}

func newUser(
	id string,
    email string,
    lastName string,
    firstName string,
    emailAddress string,
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
        email:        email,
        lastName:     lastName,
        firstName:    firstName,
        emailAddress: emailAddress,
    }, nil
}

func (u *User) ID() string {
	return u.id
}

func (u *User) Email() string {
	return u.email
}

func (u *User) LastName() string {
	return u.lastName
}

func (u *User) FirstName() string {
	return u.firstName
}

const (
	nameLengthMin = 1
    nameLengthMax = 50
)