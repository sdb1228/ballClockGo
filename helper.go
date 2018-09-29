package main

import "strconv"
import "errors"

const MAXBALLS = 127
const MINBALLS = 27

func parseBallsAndMinutes(ballsInput string, minutesInput string) (balls int, minutes int, err error) {
	balls, ballErr := strconv.Atoi(ballsInput)
	minutes, minuteErr := strconv.Atoi(minutesInput)
	if ballErr != nil {
		return 0, 0, errors.New("invalid balls provided")
	}
	if minuteErr != nil {
		return 0, 0, errors.New("invalid minutes input provided")
	}
	if minutes < 0 {
		return 0, 0, errors.New("minutes must be positive")
	}
	if balls < MINBALLS || balls > MAXBALLS {
		return 0, 0, errors.New("ensure that you have less than 27 balls and greater than 127 minutes")
	}
	return balls, minutes, nil
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func reverseArray(array []int) []int {
	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		array[i], array[j] = array[j], array[i]
	}
	return array
}
