package calendar

type Period string

const (
	Day   = Period("day")
	Week  = Period("week")
	Month = Period("month")
	Year  = Period("year")
)

func GetPeriodMap() map[Period]bool {
	return map[Period]bool{
		Day:   true,
		Week:  true,
		Month: true,
		Year:  true,
	}
}
