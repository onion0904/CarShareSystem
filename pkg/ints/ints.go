package ints

import "strconv"

func Digit(num int) int {
	str := strconv.Itoa(num)
    return len(str)
}