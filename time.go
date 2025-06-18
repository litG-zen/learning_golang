// Writing a basic time-handling code to build understanding of how the Time module works in GoLang.

package main

import (
	"fmt"
	"time"
)

func main() {

	// Getting current time
	now := time.Now()

	fmt.Printf("Now: %v", now)

	// timeZone details
	zone, offset := now.Zone()

	fmt.Printf("Current time details !\nTime: %v,\nDay:%v\nMonth: %v,\nYear: %v,\nTimeZone: %v,\nTime-Offset :%v",
		now, now.Day(), now.Month(), now.Year(), zone, offset)

	// Parsing current date from the `now` variable
	year, month, day := now.Date()
	fmt.Printf("\nCurrent Date: %d-%02d-%04d", day, month, year)

	// Getting current weekday
	fmt.Printf("\nCurrent weekday: %v", now.Weekday())

	// Lets try to iterate over next 30 days
	new_day := now
	for i := 0; i < 30; i++ {
		oneDay := time.Hour * 24
		new_day = new_day.Add(oneDay)
		fmt.Printf("\n %v", new_day.Format("2006_01_02"))
	}
}
