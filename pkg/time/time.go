package time

import "time"

func Now() time.Time {
	return time.Now()
}
func Year() int {
	year := time.Now().Year()
	return year
}

func Month() int {
	month:= time.Now().Month()
	return int(month)
}

func Day() int {
	day := time.Now().Day()
	return day
}

//次の週の日曜日の日を出力
func NextEndWeek() time.Time {
	now := time.Now()
	// 現在の日付から次の日曜日までの日数を計算
	daysUntilSunday := (7 - int(now.Weekday())) % 7
	// 次の週の日曜日までさらに7日を加算
	return now.AddDate(0, 0, daysUntilSunday+7)
}

func NextStartWeek() time.Time {
	now := time.Now()
    daysUntilSunday := (7 - int(now.Weekday())) % 7
    return now.AddDate(0, 0, daysUntilSunday+1)
}