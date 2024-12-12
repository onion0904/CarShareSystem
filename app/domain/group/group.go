package group

import (
	"unicode/utf8"
	"github.com/onion0904/go-pkg/ulid"
	errDomain "github.com/onion0904/app/domain/error"
)


type Group struct {
	id string
	name string
	membersID []string
}


func Reconstruct(
	id string,
	name string,
	membersID []string,
) (*Group, error) {
	return newGroup(
		id,
		name,
		membersID,
	)
}

func NewProduct(
	id string,
	name string,
	membersID []string,
) (*Group, error) {
	return newGroup(
		ulid.NewUlid(),
		name,
		membersID,
	)
}

func newGroup(
	id string,
	name string,
	membersID []string,
) (*Group, error) {
	// ownerIDのバリデーション
	if !ulid.IsValid(id) {
		return nil, errDomain.NewError("オーナーIDの値が不正です。")
	}
	// 名前のバリデーション
	if utf8.RuneCountInString(name) < nameLengthMin || utf8.RuneCountInString(name) > nameLengthMax {
		return nil, errDomain.NewError("グループ名が不正です。")
	}

	return &Group{
		id:          id,
		name:        name,
		membersID:   membersID,
	}, nil
}

func (p *Group) ID() string {
	return p.id
}

func (p *Group) Name() string {
	return p.name
}

func (p *Group) MemebersID() []string {
	return p.membersID
}


const (
	nameLengthMin = 1
	nameLengthMax = 100
)