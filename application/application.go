package main

import (
	"fmt"

	"github.com/YoeaKai/leet_code/topic/meeting_rooms_II"
)

func run() {
	println(meeting_rooms_II.MinMeetingRooms([]*meeting_rooms_II.Interval{
		{Start: 1, End: 4},
		{Start: 6, End: 12},
		{Start: 7, End: 9},
		{Start: 8, End: 14},
		{Start: 9, End: 12},
		{Start: 13, End: 15},
	}))
}

func main() {
	fmt.Println("----------Start----------")
	run()
	fmt.Println("----------Finish----------")
}
