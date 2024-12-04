package main

import (
	"exercises/chapter_9/greetings"
	"fmt"
)

func mainGreetings() {
	name := "James"
	fmt.Println(greetings.GoodDay(name))
	fmt.Println(greetings.GoodNight(name))

	if greetings.IsAm() {
		fmt.Println("Good morning", name)
	} else if greetings.IsPm() {
		fmt.Println("Good afternoon", name)
	} else if greetings.IsEvening() {
		fmt.Println("Good night", name)
	} else {
		fmt.Println("Good night", name)
	}
}
