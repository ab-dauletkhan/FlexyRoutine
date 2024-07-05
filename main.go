package main

import (
	"fmt"

	"github.com/ab-dauletkhan/flexy/cmd/prayer"
)

func main() {
	fmt.Println("Welcome to Flexy Routine App")
	fmt.Println("This program let's you add prayer times for your Google Calendar")
	fmt.Println("First let's choose the location from the below list")
	prayer.ReceiveCity()
	prayer.GetPrayerTimes(2024)
}
