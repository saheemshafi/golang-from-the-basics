package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Times and dates")

	currentDate := time.Now()
	fmt.Println("The date now is: ", currentDate)

	// Formatting the time
	/* The datetime layout used in go is:
	Month 1
	Date 2
	Hour 3
	Minute 4
	Second 5
	Year 6

	 Month - Date - Year  Hour : Minute : Second Day
	"01-02-2006 15:04:05 Monday"
	*/

	fmt.Println("The formatted date is: ", currentDate.Format("01-02-2006 15:04:05 Monday"))
	fmt.Println("The formatted date is: ", currentDate.Format("01/02/2006 15:04 Monday"))
	fmt.Println("The formatted date is: ", currentDate.Format("Mon 15:04 01:02:2006"))
	fmt.Println("The formatted date is: ", currentDate.Format("Mon 2 Jan 2006"))
	fmt.Println("The formatted date is: ", currentDate.Format("2 Jan 2006 Monday"))

	createdDate := time.Date(2006, time.July, 26, 0, 0, 0, 0, time.UTC)
	fmt.Println("My birthday is on: ", createdDate.Format("02 January 2006"))
}
