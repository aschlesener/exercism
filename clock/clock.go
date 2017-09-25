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
	newHour, newMinute := calcRollover(hour, minute, 0, 0)
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

// helper function to handle minute and hour rollover
func calcRollover(newHour, newMinute, oldHour, oldMinute int) (hour, minute int) {
	if oldHour == 0 && oldMinute == 0 {
		hour = newHour
		minute = newMinute
	} else {
		hour = oldHour
		minute = oldMinute
	}

	absNewMinute := int(math.Abs(float64(newMinute)))

	// handle minute rollover
	if newMinute+oldMinute == 60 {
		// 60 mins, just add an hour
		hour++
		minute = 0
	} else if newMinute+oldMinute > 60 {
		// more than 60 minutes being added
		numHours := (oldMinute + newMinute) / 60
		hour += numHours
		numMinutes := (newMinute + oldMinute) % 60
		minute = numMinutes
	} else if newMinute+oldMinute < 0 {
		// negative minutes
		if newMinute+oldMinute > -60 {
			// less than an hour being subtracted
			hour--
			minute = 60 + (oldMinute - absNewMinute)
		} else if newMinute+oldMinute < -60 {
			// more than one hour being subtracted
			numHours := (absNewMinute + oldMinute) / 60
			hour -= numHours
			var numMinutes int
			if (newMinute % 60) == 0 {
				numMinutes = oldMinute
			} else {
				numMinutes = (oldMinute + newMinute) % 60
			}

			minute = numMinutes
			if minute < 0 {
				hour--
				minute = 60 + (oldMinute - int(math.Abs(float64(minute))))
			}
			if (newMinute % 60) != 0 {
				minute = int(math.Abs(float64(oldMinute - minute)))
			}

		}
	} else {
		minute = oldMinute + newMinute
	}

	// handle hour rollover
	if hour >= 24 {
		hour = hour % 24
	} else if int(math.Abs(float64(hour))) == 24 {
		hour = 0
	} else if hour > -24 && hour < 0 {
		hour = 24 + hour
	} else if hour < -24 {
		hour = 24 - int(math.Abs(float64(hour)))%24
		if hour == 24 {
			hour = 0
		}
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
