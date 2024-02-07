package model

type DayOfWeek int

const (
	SUNDAY DayOfWeek = iota
	MONDAY
	TUESDAY
	WEDNESDAY
	THURSDAY
	FRIDAY
	SATURDAY
)

func DayOfWeekIntToString(d int) string {
	switch DayOfWeek(d) {
	case SUNDAY:
		return "SUNDAY"
	case MONDAY:
		return "MONDAY"
	case TUESDAY:
		return "TUESDAY"
	case WEDNESDAY:
		return "WEDNESDAY"
	case THURSDAY:
		return "THURSDAY"
	case FRIDAY:
		return "FRIDAY"
	case SATURDAY:
		return "SUSATURDAY"
	}
	return ""
}

func DayOfWeekStringToInt(d string) int {
	switch d {
	case "SUNDAY":
		return int(SUNDAY)
	case "MONDAY":
		return int(MONDAY)
	case "TUESDAY":
		return int(TUESDAY)
	case "WEDNESDAY":
		return int(WEDNESDAY)
	case "THURSDAY":
		return int(THURSDAY)
	case "FRIDAY":
		return int(FRIDAY)
	case "SATURDAY":
		return int(SATURDAY)
	}
	return 0
}
