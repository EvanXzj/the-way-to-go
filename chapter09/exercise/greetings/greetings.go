package greetings

import "time"

func GoodDay(name string) string {
	return "Good Day " + name
}

func GoodNight(name string) string {
	return "Good Night " + name
}

func IsAM() bool {
	t := time.Now()

	return t.Hour() <= 12
}

func IsAfternoon() bool {
	t := time.Now()

	return t.Hour() > 12 && t.Hour() <= 18
}

func IsEvening() bool {
	t := time.Now()
	return t.Hour() <= 22
}
