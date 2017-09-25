package datetimeManipulation

import (
	"fmt"
	"time"
)

func DatetimeManipulationExamples() {
	t := time.Now()
	fmt.Println("Location : ", t.Location(), " Time : ", t) // local time

	location, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Location : ", location, " Time : ", t.In(location)) // America/New_York

	loc, _ := time.LoadLocation("Asia/Shanghai")
	now := time.Now().In(loc)
	fmt.Println("Location : ", loc, " Time : ", now) // Asia/Shanghai
	return
}

func ZoneTime() {

	t := time.Now()
	z, _ := t.Zone()
	fmt.Println("ZONE : ", z, " Time : ", t) // local time

	location, err := time.LoadLocation("EST")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ZONE : ", location, " Time : ", t.In(location)) // EST

	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)
	fmt.Println("ZONE : ", loc, " Time : ", now) // UTC

	loc, _ = time.LoadLocation("MST")
	now = time.Now().In(loc)
	fmt.Println("ZONE : ", loc, " Time : ", now) // MST

}

func TimeDiff() {
	loc, _ := time.LoadLocation("UTC")
	now := time.Now().In(loc)
	fmt.Println("\nToday : ", loc, " Time : ", now)

	pastDate := time.Date(2017, time.August, 10, 23, 10, 52, 211, time.UTC)
	fmt.Println("\n\nPast  : ", loc, " Time : ", pastDate) //
	fmt.Printf("###############################################################\n")
	diff := now.Sub(pastDate)

	hrs := int(diff.Hours())
	fmt.Printf("Diffrence in Hours : %d Hours\n", hrs)

	mins := int(diff.Minutes())
	fmt.Printf("Diffrence in Minutes : %d Minutes\n", mins)

	second := int(diff.Seconds())
	fmt.Printf("Diffrence in Seconds : %d Seconds\n", second)

	days := int(diff.Hours() / 24)
	fmt.Printf("Diffrence in days : %d days\n", days)

	fmt.Printf("###############################################################\n\n\n")

	futureDate := time.Date(2019, time.August, 10, 23, 10, 52, 211, time.UTC)
	fmt.Println("Future  : ", loc, " Time : ", futureDate) //
	fmt.Printf("###############################################################\n")
	diff = futureDate.Sub(now)

	hrs = int(diff.Hours())
	fmt.Printf("Diffrence in Hours : %d Hours\n", hrs)

	mins = int(diff.Minutes())
	fmt.Printf("Diffrence in Minutes : %d Minutes\n", mins)

	second = int(diff.Seconds())
	fmt.Printf("Diffrence in Seconds : %d Seconds\n", second)

	days = int(diff.Hours() / 24)
	fmt.Printf("Diffrence in days : %d days\n", days)

}

func DateAddition() {
	now := time.Now()
	fmt.Println("\nToday:", now)

	after := now.AddDate(1, 0, 0)
	fmt.Println("\nAdd 1 Year:", after)

	after = now.AddDate(0, 1, 0)
	fmt.Println("\nAdd 1 Month:", after)

	after = now.AddDate(0, 0, 1)
	fmt.Println("\nAdd 1 Day:", after)

	after = now.AddDate(2, 2, 5)
	fmt.Println("\nAdd multiple values:", after)

	after = now.Add(10 * time.Minute)
	fmt.Println("\nAdd 10 Minutes:", after)

	after = now.Add(10 * time.Second)
	fmt.Println("\nAdd 10 Second:", after)

	after = now.Add(10 * time.Hour)
	fmt.Println("\nAdd 10 Hour:", after)

	after = now.Add(10 * time.Millisecond)
	fmt.Println("\nAdd 10 Millisecond:", after)

	after = now.Add(10 * time.Microsecond)
	fmt.Println("\nAdd 10 Microsecond:", after)

	after = now.Add(10 * time.Nanosecond)
	fmt.Println("\nAdd 10 Nanosecond:", after)

	after = now.AddDate(0, 0, 115)
	fmt.Println("Currently breeding for:", after)
}
