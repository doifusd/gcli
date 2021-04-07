package timer

import (
	"time"
)

//const (
//	Nanosecond  Duration = 1
//	Microsecond          = 1000 * Nanosecond
//	Millisecond          = 1000 * Microsecond
//	Second               = 1000 * Millisecond
//	Minute               = 60 * Second
//	Hour                 = 60 * Minute
//)

func GetNowTime() time.Time {
	return time.Now()
}

func GetCalculateTime(cTime time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}
	return cTime.Add(duration), nil
	// return GetNowTime().Add(time.Second*60)
}
