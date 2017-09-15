package gigasecond

// import path for the time package from the standard library
import "time"

const testVersion = 4

func AddGigasecond(date time.Time) time.Time {
	return date.Add(1000000000 * time.Second)
}
