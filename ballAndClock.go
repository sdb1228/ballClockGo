package main

import "fmt"
import "os"
import "log"
import "time"
import "reflect"

func main() {
	start := time.Now()
	if len(os.Args) > 2 {
		balls, minutes, err := parseBallsAndMinutes(os.Args[1], os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		calculateClockWithMinutes(balls, minutes)
	} else {
		balls, _, err := parseBallsAndMinutes(os.Args[1], "0")
		if err != nil {
			log.Fatal(err)
		}
		calculateClockWithoutMinutes(balls)
		elapsed := time.Since(start)
		fmt.Printf("Completed in %v milliseconds (%v seconds)\n", elapsed, elapsed.Seconds())
	}
}

func calculateClockWithMinutes(balls int, minutes int) {
	queue := makeRange(1, balls)
	var oneMinuteShaft []int
	var fiveMinuteShaft []int
	var oneHourShaft []int
	for i := 0; i < minutes; i++ {
		if len(oneMinuteShaft) != 4 {
			oneMinuteShaft = append(oneMinuteShaft, queue[0])
			queue = queue[1:]
			continue
		} else if len(oneMinuteShaft) == 4 && len(fiveMinuteShaft) != 11 {
			queue = append(queue, reverseArray(oneMinuteShaft)...)
			oneMinuteShaft = nil
			fiveMinuteShaft = append(fiveMinuteShaft, queue[0])
			queue = queue[1:]
			continue
		} else if len(oneMinuteShaft) == 4 && len(fiveMinuteShaft) == 11 && len(oneHourShaft) != 11 {
			queue = append(queue, reverseArray(oneMinuteShaft)...)
			queue = append(queue, reverseArray(fiveMinuteShaft)...)
			oneMinuteShaft = nil
			fiveMinuteShaft = nil
			oneHourShaft = append(oneHourShaft, queue[0])
			queue = queue[1:]
			continue
		} else {
			queue = append(queue, reverseArray(oneMinuteShaft)...)
			queue = append(queue, reverseArray(fiveMinuteShaft)...)
			queue = append(queue, reverseArray(oneHourShaft)...)
			lastBall := queue[0]
			queue = queue[1:]
			queue = append(queue, lastBall)
			oneMinuteShaft = nil
			fiveMinuteShaft = nil
			oneHourShaft = nil
		}
	}
	fmt.Printf("Min: %v FiveMin: %v Hour: %v Queue: %v \n", oneMinuteShaft, fiveMinuteShaft, oneHourShaft, queue)
}

func calculateClockWithoutMinutes(balls int) {
	queue := makeRange(1, balls)
	start := makeRange(1, balls)
	var oneMinuteShaft []int
	var fiveMinuteShaft []int
	var oneHourShaft []int
	count := 0
	firstIterator := true
	for !reflect.DeepEqual(queue, start) || firstIterator {
		if len(oneMinuteShaft) != 4 {
			oneMinuteShaft = append(oneMinuteShaft, queue[0])
			queue = queue[1:]
			continue
		} else if len(oneMinuteShaft) == 4 && len(fiveMinuteShaft) != 11 {
			queue = append(queue, reverseArray(oneMinuteShaft)...)
			oneMinuteShaft = nil
			fiveMinuteShaft = append(fiveMinuteShaft, queue[0])
			queue = queue[1:]
			continue
		} else if len(oneMinuteShaft) == 4 && len(fiveMinuteShaft) == 11 && len(oneHourShaft) != 11 {
			queue = append(queue, reverseArray(oneMinuteShaft)...)
			queue = append(queue, reverseArray(fiveMinuteShaft)...)
			oneMinuteShaft = nil
			fiveMinuteShaft = nil
			oneHourShaft = append(oneHourShaft, queue[0])
			queue = queue[1:]
			continue
		} else {
			count++
			queue = append(queue, reverseArray(oneMinuteShaft)...)
			queue = append(queue, reverseArray(fiveMinuteShaft)...)
			queue = append(queue, reverseArray(oneHourShaft)...)
			lastBall := queue[0]
			queue = queue[1:]
			queue = append(queue, lastBall)
			oneMinuteShaft = nil
			fiveMinuteShaft = nil
			oneHourShaft = nil
		}
		firstIterator = false
	}
	fmt.Println("Number of Days until the queue is the same as when it started: ")
	fmt.Println(count / 2)

}
