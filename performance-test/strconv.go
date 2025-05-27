package main

import (
	"fmt"
	"strconv"
	"time"
)

func TestNumbersToString_1(iterations int) {
	defer PrintDurationWithTime("string + strconv int64", iterations, time.Now())

	var number int64 = 100
	for i := 0; i < iterations; i++ {
		_ = "message:" + strconv.FormatInt(number, 10)
	}
}

func TestNumbersToString_2(iterations int) {
	defer PrintDurationWithTime("fmt string int64", iterations, time.Now())

	var number int64 = 100
	for i := 0; i < iterations; i++ {
		_ = fmt.Sprintf("message: %d", number)
	}
}
