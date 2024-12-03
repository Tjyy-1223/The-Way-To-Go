package main

import "fmt"

var Days = map[int]string{
	1: "monday",
	2: "tuesday",
	3: "wednesday",
	4: "thursday",
	5: "friday",
	6: "saturday",
	7: "sunday",
}

func main() {
	fmt.Println(Days)

	flag := false
	for k, v := range Days {
		if v == "thursday" || v == "holliday" {
			fmt.Println(v, "is the ", k, "th day in the week")
			if v == "holliday" {
				flag = true
			}
		}
	}
	if !flag {
		fmt.Println("holliday is not a day!")
	}

}
