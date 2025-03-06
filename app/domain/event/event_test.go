package event

import (
	"reflect"
	"testing"
	"time"
)

func TestReconstruct(t *testing.T) {
	type args struct {
		id          string
		userID     string
		together    bool
		description string
		year        int32
		month       int32
		day         int32
		date        time.Time
		startDate   time.Time
		endDate     time.Time
		important   bool
	}
	tests := []struct {
		name    string
		args    args
		want    *Event
		wantErr bool
	}{
		{
			name: "ok case: Reconstruct & Getter",
			args: args{
				id:          "01F8B9Z6G9WBJK9XJH5M7RQK5X",  // 有効なULIDの例
                userID:      "user123",
                together:    true,
                description: "Test Event",
				year:        2000,
                month:       12,
                day:         12,
                date:        time.Time{},
                startDate:   time.Time{},
                endDate:     time.Time{},
				important:   true,
			},
			want: &Event{
				id:          "01F8B9Z6G9WBJK9XJH5M7RQK5X",
                userID:      "user123",
                together:    true,
                description: "Test Event",
                year:        2000,
                month:       12,
                day:         12,
                date:        time.Time{},
                startDate:   time.Time{},
                endDate:     time.Time{},
                important:   true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Reconstruct(tt.args.id, tt.args.userID, tt.args.together, tt.args.description, tt.args.year, tt.args.month, tt.args.day, tt.args.date, tt.args.startDate, tt.args.endDate, tt.args.important)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reconstruct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Reconstruct() = %v, want %v", got, tt.want)
			}
			if got.ID() != tt.args.id {
				t.Errorf("got.ID() = %v, want %v", got.ID(), tt.args.id)
			}
			if got.UserID() != tt.args.userID {
				t.Errorf("got.UserID() = %v, want %v", got.UserID(), tt.args.userID)
			}
			if got.Together() != tt.args.together {
				t.Errorf("got.Together() = %v, want %v", got.Together(), tt.args.together)
			}
			if got.Description() != tt.args.description {
				t.Errorf("got.Description() = %v, want %v", got.Description(), tt.args.description)
			}
			if got.Year() != tt.args.year {
				t.Errorf("got.Year() = %v, want %v", got.Year(), tt.args.year)
			}
			if got.Month() != tt.args.month {
				t.Errorf("got.Month() = %v, want %v", got.Month(), tt.args.month)
			}
			if got.Day() != tt.args.day {
				t.Errorf("got.Day() = %v, want %v", got.Day(), tt.args.day)
			}
			if got.Date() != tt.args.date {
				t.Errorf("got.Year() = %v, want %v", got.Date(), tt.args.date)
			}
			if got.StartDate() != tt.args.startDate {
				t.Errorf("got.StartDate() = %v, want %v", got.StartDate(), tt.args.startDate)
			}
			if got.EndDate() != tt.args.endDate {
				t.Errorf("got.EndDate() = %v, want %v", got.EndDate(), tt.args.endDate)
			}
			if got.Important() != tt.args.important {
				t.Errorf("got.Important() = %v, want %v", got.Important(), tt.args.important)
			}
		})
	}
}

func TestNewEvent(t *testing.T) {
	type args struct {
		userID     string
		together    bool
		description string
		important   bool
	}
	tests := []struct {
		name    string
		args    args
		want    *Event
		wantErr bool
	}{
		{
			name: "ok case: New",
			args: args{
                userID:      "user123",
                together:    true,
                description: "Test Event",
				important:   true,
			},
			want: &Event{
				id:          "01F8B9Z6G9WBJK9XJH5M7RQK5X",
                userID:      "user123",
                together:    true,
                description: "Test Event",
                important:   true,
			},
			wantErr: false,
		},
		{
			name: "error case: description",
			args: args{
				userID:      "user123",
                together:    true,
                description: "",
                important:   true,
			},
			wantErr: true,
		},
		{
			name: "error case: description",
			args: args{
				userID:      "user123",
                together:    true,
                description: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
                important:   true,
			},
			wantErr: true,
		},
		{
			name: "error case: year",
			args: args{
				userID:      "user123",
                together:    true,
                description: "Test year",
                important:   true,
			},
			wantErr: true,
		},
		{
			name: "error case: month",
			args: args{
				userID:      "user123",
                together:    true,
                description: "Test month",
                important:   true,
			},
			wantErr: true,
		},
		{
			name: "error case: day",
			args: args{
				userID:      "user123",
                together:    true,
                description: "Test day",
                important:   true,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewEvent(tt.args.userID, tt.args.together, tt.args.description, tt.args.important)
			if got != nil {
				got.id = "01F8B9Z6G9WBJK9XJH5M7RQK5X" //ulidがランダムで生成されるため
				// 現在の時刻になるため
				if tt.name == "ok case: New"{
					tt.want.year = got.year
                    tt.want.month = got.month
                    tt.want.day = got.day
					tt.want.date = got.date
					tt.want.startDate = got.startDate
					tt.want.endDate = got.endDate
				}
				if tt.name == "error case: year"{
					got.year = 2424142
				}
				if tt.name == "error case: month"{
					got.month = 412412414
				}
				if tt.name == "error case: day"{
					got.day = 41251324
				}
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("NewEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_newEvent(t *testing.T) {
	type args struct {
		id          string
		userID     string
		together    bool
		description string
		year        int32
		month       int32
		day         int32
		date        time.Time
		startDate   time.Time
		endDate     time.Time
		important   bool
	}
	tests := []struct {
		name    string
		args    args
		want    *Event
		wantErr bool
	}{
		{
			name: "ok case: new",
			args: args{
				id:          "01F8B9Z6G9WBJK9XJH5M7RQK5X",
                userID:      "user123",
                together:    true,
                description: "Test Event",
                year:        2000,
                month:       12,
                day:         12,
                date:        time.Time{},
                startDate:   time.Time{},
                endDate:     time.Time{},
                important:   true,
			},
			want: &Event{
				id:          "01F8B9Z6G9WBJK9XJH5M7RQK5X",
                userID:      "user123",
                together:    true,
                description: "Test Event",
				year:        2000,
                month:       12,
                day:         12,
                date:        time.Time{},
                startDate:   time.Time{},
                endDate:     time.Time{},
                important:   true,
			},
			wantErr: false,
		},
		{
			name: "error case: id",
			args: args{
				id:          "id",
                userID:      "user123",
                together:    true,
                description: "Test Event",
                year:        2000,
                month:       12,
                day:         12,
                date:        time.Time{},
                startDate:   time.Time{},
                endDate:     time.Time{},
                important:   true,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newEvent(tt.args.id, tt.args.userID, tt.args.together, tt.args.description, tt.args.year, tt.args.month, tt.args.day, tt.args.date, tt.args.startDate, tt.args.endDate, tt.args.important)
			if (err != nil) != tt.wantErr {
				t.Errorf("newEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUser_SetCreatedAt(t *testing.T) {
	event := &Event{}
	now := time.Now()
	event.SetCreatedAt(now)
	if !event.CreatedAt().Equal(now) {
		t.Errorf("SetCreatedAt() failed: got %v, want %v", event.CreatedAt(), now)
	}
}

func TestUser_SetUpdatedAt(t *testing.T) {
	event := &Event{}
	now := time.Now()
	event.SetUpdatedAt(now)
	if !event.UpdatedAt().Equal(now) {
		t.Errorf("SetUpdatedAt() failed: got %v, want %v", event.UpdatedAt(), now)
	}
}