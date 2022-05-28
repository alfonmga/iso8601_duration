// Package iso8601_duration provides ISO 8601 duration utilities.
package iso8601_duration

import (
	"regexp"
	"strconv"
)

// DurationStrToSeconds ISO 8601 duration string to seconds
func DurationStrToSeconds(duration string) uint32 {
	r, _ := regexp.Compile(`^P(?:(\d+)Y)?(?:(\d+)M)?(?:(\d+)W)?(?:(\d+)D)?(?:T(?:(\d+)H)?(?:(\d+)M)?(?:(\d+)S)?)?$`)
	rMatches := r.FindAllSubmatch([]byte(duration), -1)

	matchToIntV := func(v []byte) uint32 {
		if len(v) > 0 {
			vInt, err := strconv.Atoi(string(v))
			checkErr(err)
			return uint32(vInt)
		}
		return 0
	}

	result := uint32(0)

	result += matchToIntV(rMatches[0][1]) * 31536000 // years
	result += matchToIntV(rMatches[0][2]) * 18748800 // months
	result += matchToIntV(rMatches[0][3]) * 604800   // weeks
	result += matchToIntV(rMatches[0][4]) * 86400    // days
	result += matchToIntV(rMatches[0][5]) * 3600     // hours
	result += matchToIntV(rMatches[0][6]) * 60       // minutes
	result += matchToIntV(rMatches[0][7])            // seconds

	return result
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
