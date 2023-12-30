package time

import (
	"time"
)

var jst *time.Location

func init() {
	var err error
	jst, err = time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
}

type Clock interface {
	Now() time.Time
	Parse(layout, value string) (time.Time, error)
	TimeToString(layout string) string
	StringToTime(str string, layout string) time.Time
}

type realClock struct {
	location *time.Location
}

func NewRealClock(loc time.Location) Clock {
	return &realClock{location: &loc}
}

func NewClockJST() (Clock, error) {
	return NewRealClock(*jst), nil
}

func (c realClock) Now() time.Time {
	return time.Now().In(c.location)
}

func (c realClock) Parse(layout string, value string) (time.Time, error) {
	return time.ParseInLocation(layout, value, c.location)
}

func (c realClock) TimeToString(layout string) string {
	return timeToString(c.Now(), layout)
}

func (c realClock) StringToTime(str string, layout string) time.Time {
	return stringToTime(str, layout)
}

type dummyClock struct {
	now func() time.Time
}

func (c dummyClock) Now() time.Time {
	return c.now()
}

func (c dummyClock) Parse(layout string, value string) (time.Time, error) {
	return time.ParseInLocation(layout, value, c.Now().Location())
}

func (c dummyClock) TimeToString(layout string) string {
	return timeToString(c.Now(), layout)
}

func (c dummyClock) StringToTime(str string, layout string) time.Time {
	return stringToTime(str, layout)
}

func NewDummyClock(f func() time.Time) Clock {
	return &dummyClock{now: f}
}

func NewFixedClock(t time.Time) Clock {
	return NewDummyClock(
		func() time.Time { return t },
	)
}
func stringToTime(str string, layout string) time.Time {
	t, _ := time.Parse(layout, str)
	return t
}

func timeToString(t time.Time, layout string) string {
	str := t.Format(layout)
	return str
}
