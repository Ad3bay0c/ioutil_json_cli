package dateandtime

import (
	"fmt"
	"time"
)

func TimeTesting() {
	t1 := time.Now()
	t := time.Now().Format(time.RFC850)
	fmt.Println(t)

	d, _ := time.Parse(time.RFC1123, "2 days")
	fmt.Println(d)
	t2 := time.Now()
	// diff := t2-t1
	fmt.Printf("The start Duration: %v, the end: %v\nThe Difference is: %v\n", t1, t2, t2.Sub(t1))

	fmt.Println(time.Since(t1))
	// fmt.Println(time.Until(t1))
	tString, _ := time.ParseDuration("1h")
	fmt.Println(tString)

	t3, _ := time.ParseDuration("14h12m5s")
	fmt.Println(t3.Truncate(time.Hour)) // Truncate return multiple of m
	//_ = time.Now().Date() // returns current year, month, and day

	currentDate := time.Date(2019, time.November, 18, 01, 0, 0, 0, time.UTC)

	fmt.Println(currentDate.Add(10 * time.Hour))
	fmt.Println(currentDate.Add(-10 * time.Hour))
	fmt.Println(currentDate.AddDate(0, -10, -5))

	tm := []byte("Time: ")
	text := currentDate.AppendFormat(tm, time.Kitchen)
	fmt.Println(string(text))

	year, week := time.Now().ISOWeek()
	fmt.Println(year, week)
}
