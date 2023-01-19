package age

import (
	"testing"
	"time"
)

type AgeTable struct {
	Name          string
	BirthDate     time.Time
	FromDate      time.Time
	ExpectedYears int
	ExpectedMonth int
	ExpectedDays  int
}

func TestAge(t *testing.T) {
	tt := []AgeTable{
		{
			Name:          "normal",
			BirthDate:     time.Date(1989, 1, 1, 0, 0, 0, 0, time.Local),
			FromDate:      time.Date(2022, 8, 25, 0, 0, 0, 0, time.Local),
			ExpectedYears: 33,
			ExpectedMonth: 7,
			ExpectedDays:  24,
		},
		{
			Name:          "normal_same_year",
			BirthDate:     time.Date(2022, 2, 6, 0, 0, 0, 0, time.Local),
			FromDate:      time.Date(2022, 8, 25, 0, 0, 0, 0, time.Local),
			ExpectedYears: 0,
			ExpectedMonth: 6,
			ExpectedDays:  19,
		},
		{
			Name:          "leap_year_to_leap_year",
			BirthDate:     time.Date(2000, 2, 6, 0, 0, 0, 0, time.Local),  // leap year
			FromDate:      time.Date(2016, 2, 29, 0, 0, 0, 0, time.Local), // leap year
			ExpectedYears: 16,
			ExpectedMonth: 0,
			ExpectedDays:  23,
		},
		{
			Name:          "leap_year_to_leap_year_boundaries",
			BirthDate:     time.Date(2000, 2, 29, 0, 0, 0, 0, time.Local), // leap year
			FromDate:      time.Date(2020, 2, 28, 0, 0, 0, 0, time.Local), // leap year
			ExpectedYears: 19,
			ExpectedMonth: 11,
			ExpectedDays:  28,
		},
		{
			Name:          "leap_year_to_leap_year_boundaries_two",
			BirthDate:     time.Date(2000, 2, 28, 0, 0, 0, 0, time.Local), // leap year
			FromDate:      time.Date(2020, 2, 29, 0, 0, 0, 0, time.Local), // leap year
			ExpectedYears: 20,
			ExpectedMonth: 0,
			ExpectedDays:  1,
		},
		{
			Name:          "leap_year_to_leap_year_boundaries_three",
			BirthDate:     time.Date(2000, 2, 13, 0, 0, 0, 0, time.Local), // leap year
			FromDate:      time.Date(2020, 2, 04, 0, 0, 0, 0, time.Local), // leap year
			ExpectedYears: 19,
			ExpectedMonth: 11,
			ExpectedDays:  4,
		},
		{
			Name:          "normal_months",
			BirthDate:     time.Date(2022, 2, 13, 0, 0, 0, 0, time.Local), // normal year
			FromDate:      time.Date(2022, 6, 04, 0, 0, 0, 0, time.Local), // normal year
			ExpectedYears: 0,
			ExpectedMonth: 3,
			ExpectedDays:  19,
		},
		{
			Name:          "multiple_days",
			BirthDate:     time.Date(2022, 6, 13, 0, 0, 0, 0, time.Local), // normal year
			FromDate:      time.Date(2022, 6, 20, 0, 0, 0, 0, time.Local), // normal year
			ExpectedYears: 0,
			ExpectedMonth: 0,
			ExpectedDays:  7,
		},
		{
			Name:          "new_born",
			BirthDate:     time.Date(2022, 6, 13, 12, 30, 0, 0, time.Local),     // normal year
			FromDate:      time.Date(2022, 6, 13, 14, 38, 12, 3455, time.Local), // normal year
			ExpectedYears: 0,
			ExpectedMonth: 0,
			ExpectedDays:  0,
		},
	}

	for _, item := range tt {
		t.Run(item.Name, func(t *testing.T) {
			age := AgeAt(item.BirthDate, item.FromDate)
			if age.Years != item.ExpectedYears {
				t.Errorf("Expected age to be %d. Got %d", item.ExpectedYears, age.Years)
			}

			if age.Months != item.ExpectedMonth {
				t.Errorf("Expected months to be %d. Got %d", item.ExpectedMonth, age.Months)
			}

			if age.Days != item.ExpectedDays {
				t.Errorf("Expected days to be %d. Got %d", item.ExpectedDays, age.Days)
			}

		})
	}
}
