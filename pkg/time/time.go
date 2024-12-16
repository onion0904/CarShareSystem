package time

import "time"

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