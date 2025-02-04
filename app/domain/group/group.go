package group

import (
	"unicode/utf8"
	"github.com/onion0904/go-pkg/ulid"
	errDomain "github.com/onion0904/app/domain/error"
)


type Group struct {
	id string
	name string
	userIDs []string
	eventIDs []string
	icon string
}


func Reconstruct(
	id string,
	name string,
	userIDs []string,
	eventIDs []string,
	icon string,
) (*Group, error) {
	return newGroup(
		id,
		name,
		userIDs,
		eventIDs,
		icon,
	)
}

func NewGroup(
	name string,
	userIDs []string,
	eventIDs []string,
	icon string,
) (*Group, error) {
	return newGroup(
		ulid.NewUlid(),
		name,
		userIDs,
		eventIDs,
		icon,
	)
}

func newGroup(
	id string,
	name string,
	userIDs []string,
	eventIDs []string,
	icon string,
) (*Group, error) {
	// ownerIDのバリデーション
	if !ulid.IsValid(id) {
		return nil, errDomain.NewError("オーナーIDの値が不正です。")
	}
	// 名前のバリデーション
	if utf8.RuneCountInString(name) < nameLengthMin || utf8.RuneCountInString(name) > nameLengthMax {
		return nil, errDomain.NewError("グループ名が不正です。")
	}
	// アイコンのバリデーションとデフォルトの設定
	if icon == "" {
		icon = "defaultIcon"
	}

	return &Group{
		id:          id,
		name:        name,
		userIDs:     userIDs,
		eventIDs:    eventIDs,
		icon:        icon,
	}, nil
}

func (p *Group) ID() string {
	return p.id
}

func (p *Group) Name() string {
	return p.name
}

func (p *Group) UsersID() []string {
	return p.userIDs
}

func (p *Group) EventsID() []string {
    return p.eventIDs
}

func (p *Group) Icon() string {
    return p.icon
}


const (
	nameLengthMin = 1
	nameLengthMax = 100
)