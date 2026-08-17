package dates_test

import (
	"fmt"
	"time"

	"github.com/yuvv/candy/dates"
)

func Example() {
	t := time.Date(2024, time.March, 14, 15, 9, 26, 0, time.UTC)

	fmt.Println(dates.FormatDate(dates.BeginOfDay(t)))
	fmt.Println(dates.FormatDateTime(t))
	// Output:
	// 2024-03-14
	// 2024-03-14 15:09:26
}
