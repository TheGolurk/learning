package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()

	writeTime(start, "started")

	first := time.Now()
	fmt.Println("Currently ", first.Format("15:04:05"))
	time.Sleep(2 * time.Second)
	second := time.Now()
	fmt.Println("currently", second.Format("15:04:05"))

	writeTime(start, "With Two Seconds")

	startDate := time.Date(2018, 07, 04, 7, 00, 00, 0, time.UTC)
	elapsed := time.Since(startDate)
	fmt.Printf("%v has passed from %v\n", elapsed, startDate.Format("Monday, Jan 2 2000"))
	fmt.Printf("Hours %.0f Minutes %.0f Seconds %.0f  \n", elapsed.Hours(), elapsed.Minutes(), elapsed.Seconds())

	today := time.Now()
	future := today.AddDate(0, 6, 0)
	// past := today.AddDate(0, -6, 0)
	// hours := today.AddDate(6 * time.Hours())
	fmt.Println("It will be ", future.Format("Monday, Jan 2 "))

	bedtime := time.Date(2020, 6, 6, 23, 0, 0, 0, time.Local)
	fmt.Printf("hours %.0f until bed time \n", time.Until(bedtime).Hours())

	writeTime(start, "AT FINALLY")
	defer writeTime(start, "FINISHED")
}

func writeTime(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
