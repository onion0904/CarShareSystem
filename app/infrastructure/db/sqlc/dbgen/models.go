// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sqlc

import (
	"time"
)

type Event struct {
	ID          string
	UserID      string
	Together    bool
	Description string
	Year        int32
	Month       int32
	Day         int32
	Date        time.Time
	CreatedAt   time.Time
	UpdatedAt   time.Time
	StartDate   time.Time
	EndDate     time.Time
	Important   bool
}

type Group struct {
	ID        string
	Name      string
	Icon      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type GroupEvent struct {
	GroupID string
	EventID string
}

type GroupUser struct {
	GroupID string
	UserID  string
}

type User struct {
	ID        string
	LastName  string
	FirstName string
	Email     string
	Password  string
	Icon      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserEvent struct {
	UserID  string
	EventID string
}
