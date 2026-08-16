package dates

import "time"

const (
	DateNumLayout  = "20060102"
	DateLayout     = "2006-01-02"
	DateTimeLayout = "2006-01-02 15:04:05"
	TimeLayout     = "15:04:05"
)

// Today returns the current local time.
func Today() time.Time {
	return time.Now()
}

// Yesterday returns the current local time shifted to the previous calendar day.
func Yesterday() time.Time {
	return time.Now().AddDate(0, 0, -1)
}

// Tomorrow returns the current local time shifted to the next calendar day.
func Tomorrow() time.Time {
	return time.Now().AddDate(0, 0, 1)
}

// BeginOfDay returns the beginning of the day containing t.
func BeginOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

// EndOfDay returns the end of the day containing t.
func EndOfDay(t time.Time) time.Time {
	return NextDay(t).Add(-time.Nanosecond)
}

// NextDay returns midnight of the next calendar day after t.
func NextDay(t time.Time) time.Time {
	return BeginOfDay(t).AddDate(0, 0, 1)
}

// PrevDay returns midnight of the previous calendar day before t.
func PrevDay(t time.Time) time.Time {
	return BeginOfDay(t).AddDate(0, 0, -1)
}

// BeginOfMonth returns the beginning of the month containing t.
func BeginOfMonth(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, t.Location())
}

// EndOfMonth returns the end of the month containing t.
func EndOfMonth(t time.Time) time.Time {
	return BeginOfMonth(t).AddDate(0, 1, 0).Add(-time.Nanosecond)
}

// BeginOfYear returns the beginning of the year containing t.
func BeginOfYear(t time.Time) time.Time {
	return time.Date(t.Year(), time.January, 1, 0, 0, 0, 0, t.Location())
}

// EndOfYear returns the end of the year containing t.
func EndOfYear(t time.Time) time.Time {
	return BeginOfYear(t).AddDate(1, 0, 0).Add(-time.Nanosecond)
}

// ParseDate parses s in DateLayout using time.Local.
func ParseDate(s string) (time.Time, error) {
	return time.ParseInLocation(DateLayout, s, time.Local)
}

// ParseDateTime parses s in DateTimeLayout using time.Local.
func ParseDateTime(s string) (time.Time, error) {
	return time.ParseInLocation(DateTimeLayout, s, time.Local)
}

// FormatDate formats t using DateLayout.
func FormatDate(t time.Time) string {
	return t.Format(DateLayout)
}

// FormatDateTime formats t using DateTimeLayout.
func FormatDateTime(t time.Time) string {
	return t.Format(DateTimeLayout)
}
