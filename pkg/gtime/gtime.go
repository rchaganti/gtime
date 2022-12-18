package gtime

import (
	"fmt"
	"os"
	"text/tabwriter"
	"time"
)

type tInfo struct {
	Location string
	Time     string
}

var result = []tInfo{}

func ConvertTime(t string, tz []string) (res []tInfo) {
	// Get local time
	localLoc, _ := time.LoadLocation("Local")
	normalizedT := "2006 " + t
	localTime, err := time.ParseInLocation("2006 15:04", normalizedT, localLoc)

	// Add local time to result map
	res = append(res, tInfo{"Local", localTime.Format("15:04")})

	if err != nil {
		fmt.Printf("Error parsing supplied time string %s", t)
	}

	for _, v := range tz {
		loc, err := time.LoadLocation(v)

		if err != nil {
			fmt.Printf("%s is not a valid timezone", loc)
		}

		targetTime := localTime.In(loc)
		res = append(res, tInfo{v, targetTime.Format("15:04")})
	}

	return
}

func PrettyPrint(r []tInfo) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.AlignRight|tabwriter.Debug)

	fmt.Fprintln(w, "Location\tTime\t")
	for _, v := range r {
		fmt.Fprintln(w, v.Location, "\t", v.Time, "\t")
	}
	w.Flush()
}
