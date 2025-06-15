package main

import (
	"fmt"
	"time"
)

func main() {
	redisTimestamp := "2025-06-15T17:31:42+07:00"

	layout := "2006-01-02T15:04:05.999999999Z07:00"
	
	parsedTime, err := time.Parse(layout, redisTimestamp)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return
	}

	now := time.Now()

	diff := parsedTime.Sub(now)

	diffInSeconds := diff.Seconds()

	diffInSecondsInt := int(diffInSeconds)

	fmt.Println("Timestamp Redis:", parsedTime)
	fmt.Println("Waktu sekarang:", now)
	fmt.Println("Selisih waktu (detik float):", diffInSeconds)
	fmt.Println("Selisih waktu (detik int):", diffInSecondsInt)
}
