package dates

import (
	"testing"
	"time"
)

func TestLayoutConstantsFormatKnownTime(t *testing.T) {
	known := time.Date(2024, time.February, 29, 15, 4, 5, 123456789, time.UTC)

	if got := known.Format(DateNumLayout); got != "20240229" {
		t.Fatalf("DateNumLayout format = %q, want %q", got, "20240229")
	}
	if got := known.Format(DateLayout); got != "2024-02-29" {
		t.Fatalf("DateLayout format = %q, want %q", got, "2024-02-29")
	}
	if got := known.Format(DateTimeLayout); got != "2024-02-29 15:04:05" {
		t.Fatalf("DateTimeLayout format = %q, want %q", got, "2024-02-29 15:04:05")
	}
	if got := known.Format(TimeLayout); got != "15:04:05" {
		t.Fatalf("TimeLayout format = %q, want %q", got, "15:04:05")
	}
}

func TestDayBoundariesPreserveLocation(t *testing.T) {
	loc := time.FixedZone("TEST", 8*60*60)
	input := time.Date(2024, time.May, 17, 13, 14, 15, 16, loc)

	begin := BeginOfDay(input)
	wantBegin := time.Date(2024, time.May, 17, 0, 0, 0, 0, loc)
	if !begin.Equal(wantBegin) || begin.Location() != loc {
		t.Fatalf("BeginOfDay = %v (%p), want %v (%p)", begin, begin.Location(), wantBegin, loc)
	}

	end := EndOfDay(input)
	wantEnd := time.Date(2024, time.May, 17, 23, 59, 59, int(time.Second-time.Nanosecond), loc)
	if !end.Equal(wantEnd) || end.Location() != loc {
		t.Fatalf("EndOfDay = %v (%p), want %v (%p)", end, end.Location(), wantEnd, loc)
	}
}

func TestMonthBoundariesPreserveLocation(t *testing.T) {
	loc := time.FixedZone("TEST", -5*60*60)
	input := time.Date(2024, time.May, 17, 13, 14, 15, 16, loc)

	begin := BeginOfMonth(input)
	wantBegin := time.Date(2024, time.May, 1, 0, 0, 0, 0, loc)
	if !begin.Equal(wantBegin) || begin.Location() != loc {
		t.Fatalf("BeginOfMonth = %v (%p), want %v (%p)", begin, begin.Location(), wantBegin, loc)
	}

	end := EndOfMonth(input)
	wantEnd := time.Date(2024, time.May, 31, 23, 59, 59, int(time.Second-time.Nanosecond), loc)
	if !end.Equal(wantEnd) || end.Location() != loc {
		t.Fatalf("EndOfMonth = %v (%p), want %v (%p)", end, end.Location(), wantEnd, loc)
	}
}

func TestYearBoundariesPreserveLocation(t *testing.T) {
	loc := time.FixedZone("TEST", 90*60)
	input := time.Date(2024, time.May, 17, 13, 14, 15, 16, loc)

	begin := BeginOfYear(input)
	wantBegin := time.Date(2024, time.January, 1, 0, 0, 0, 0, loc)
	if !begin.Equal(wantBegin) || begin.Location() != loc {
		t.Fatalf("BeginOfYear = %v (%p), want %v (%p)", begin, begin.Location(), wantBegin, loc)
	}

	end := EndOfYear(input)
	wantEnd := time.Date(2024, time.December, 31, 23, 59, 59, int(time.Second-time.Nanosecond), loc)
	if !end.Equal(wantEnd) || end.Location() != loc {
		t.Fatalf("EndOfYear = %v (%p), want %v (%p)", end, end.Location(), wantEnd, loc)
	}
}

func TestEndOfMonthLeapYearFebruary(t *testing.T) {
	loc := time.FixedZone("TEST", 0)
	input := time.Date(2024, time.February, 10, 9, 8, 7, 6, loc)

	want := time.Date(2024, time.February, 29, 23, 59, 59, int(time.Second-time.Nanosecond), loc)
	if got := EndOfMonth(input); !got.Equal(want) || got.Location() != loc {
		t.Fatalf("EndOfMonth(leap February) = %v (%p), want %v (%p)", got, got.Location(), want, loc)
	}
}

func TestNextDayPrevDayAroundFixedDate(t *testing.T) {
	loc := time.FixedZone("TEST", 2*60*60)
	input := time.Date(2024, time.March, 1, 22, 30, 45, 99, loc)

	if got, want := NextDay(input), time.Date(2024, time.March, 2, 0, 0, 0, 0, loc); !got.Equal(want) || got.Location() != loc {
		t.Fatalf("NextDay = %v (%p), want %v (%p)", got, got.Location(), want, loc)
	}
	if got, want := PrevDay(input), time.Date(2024, time.February, 29, 0, 0, 0, 0, loc); !got.Equal(want) || got.Location() != loc {
		t.Fatalf("PrevDay = %v (%p), want %v (%p)", got, got.Location(), want, loc)
	}
}

func TestParseAndFormatDate(t *testing.T) {
	date, err := ParseDate("2024-02-29")
	if err != nil {
		t.Fatalf("ParseDate returned error: %v", err)
	}
	wantDate := time.Date(2024, time.February, 29, 0, 0, 0, 0, time.Local)
	if !date.Equal(wantDate) || date.Location() != time.Local {
		t.Fatalf("ParseDate = %v (%p), want %v (%p)", date, date.Location(), wantDate, time.Local)
	}
	if got := FormatDate(date); got != "2024-02-29" {
		t.Fatalf("FormatDate = %q, want %q", got, "2024-02-29")
	}

	dateTime, err := ParseDateTime("2024-02-29 15:04:05")
	if err != nil {
		t.Fatalf("ParseDateTime returned error: %v", err)
	}
	wantDateTime := time.Date(2024, time.February, 29, 15, 4, 5, 0, time.Local)
	if !dateTime.Equal(wantDateTime) || dateTime.Location() != time.Local {
		t.Fatalf("ParseDateTime = %v (%p), want %v (%p)", dateTime, dateTime.Location(), wantDateTime, time.Local)
	}
	if got := FormatDateTime(dateTime); got != "2024-02-29 15:04:05" {
		t.Fatalf("FormatDateTime = %q, want %q", got, "2024-02-29 15:04:05")
	}
}

func TestTodayYesterdayTomorrowRelations(t *testing.T) {
	today := BeginOfDay(Today())
	yesterday := BeginOfDay(Yesterday())
	tomorrow := BeginOfDay(Tomorrow())

	if want := today.AddDate(0, 0, -1); !yesterday.Equal(want) {
		t.Fatalf("BeginOfDay(Yesterday()) = %v, want %v", yesterday, want)
	}
	if want := today.AddDate(0, 0, 1); !tomorrow.Equal(want) {
		t.Fatalf("BeginOfDay(Tomorrow()) = %v, want %v", tomorrow, want)
	}
}
