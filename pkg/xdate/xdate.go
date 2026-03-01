package xdate

import "time"

const (
	location          = "America/Sao_Paulo"
	DefaultDateLayout = "2006-01-02 15:04:05"
)

func GetLocation() *time.Location {
	loc, _ := time.LoadLocation(location)
	return loc
}

func ParseDate(dateStr string, layout *string, loc *time.Location) (time.Time, error) {
	var defaultLayout = DefaultDateLayout
	if layout == nil {
		layout = &defaultLayout
	}

	if loc == nil {
		loc = GetLocation()
	}

	return time.ParseInLocation(*layout, dateStr, loc)
}
