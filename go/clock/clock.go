package clock

import (
	"math"
	"strconv"
)

// The value of testVersion here must match `targetTestVersion` in the file
// clock_test.go.
const testVersion = 4

// Clock type represents...a clock.
type Clock struct {
	Minute int
	Hour   int
}

// New will create a clock with the given hour and minute
func New(hour, minute int) Clock {
	newHour, newMinute := calcRolloverNew(hour, minute)
	clock := Clock{newMinute, newHour}
	return clock
}

// String turns a clock into a string representation
func (clock Clock) String() string {
	hour := padLeftZeroes(strconv.FormatInt(int64(clock.Hour), 10), 2)
	minute := padLeftZeroes(strconv.FormatInt(int64(clock.Minute), 10), 2)
	return hour + ":" + minute
}

// Add will add a given amount of minutes to the clock
func (clock Clock) Add(minutes int) Clock {
	newHour, newMinute := calcRollover(0, minutes, clock.Hour, clock.Minute)
	clock.Hour = newHour
	clock.Minute = newMinute
	return clock
}

// helper function to handle minute and hour rollover for add
func calcRollover(newHour, newMinute, oldHour, oldMinute int) (hour, minute int) {

	if oldMinute != 0 && 60%oldMinute < newMinute {

	} else {
		minute = oldMinute + newMinute
	}
	// if oldHour
	return hour, minute
}

// helper function to handle minute and hour rollover for new
func calcRolloverNew(newHour, newMinute int) (hour, minute int) {
	hour = newHour
	minute = newMinute
	absNewMinute := int(math.Abs(float64(newMinute)))
	absNewHour := int(math.Abs(float64(newHour)))

	// handle minute rollover
	if newMinute == 60 {
		hour++
		minute = 0
	} else if newMinute > 60 {
		numHours := newMinute / 60
		hour += numHours
		numMinutes := newMinute % 60
		minute = numMinutes
	} else if newMinute < 0 {
		if newMinute > -60 {
			numHours := int(absNewMinute) / 60
			hour -= numHours
			numMinutes := newMinute % 60
			minute = numMinutes
		}
		else if newMinute < -60 {
			
		}
	} else {
		minute = newMinute
	}

	// handle hour rollover
	if hour >= 24 {
		hour = hour % 24
	} else if hour == -24 {
		hour = 0
	} else if hour > -24 && hour < 0 {
		hour = 24 + hour
	} else if hour < -24 {
		hour = 24 - absNewHour%24
	}

	return hour, minute
}

// helper function to pad a string with zeroes because apparently that's not built in
func padLeftZeroes(str string, num int) string {
	for i := 1; i <= num-len(str); i++ {
		str = "0" + str
	}
	return str
}
