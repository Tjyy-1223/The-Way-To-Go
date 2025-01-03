package greetings

import "time"

func GoodDay(name string) string {
	return "Good Day " + name
}

func GoodNight(name string) string {
	return "Good Night " + name
}

func IsAm() bool {
	localTime := time.Now()
	return localTime.Hour() <= 12
}

func IsPm() bool {
	localTime := time.Now()
	return localTime.Hour() <= 18
}

func IsEvening() bool {
	localTime := time.Now()
	return localTime.Hour() <= 22
}
