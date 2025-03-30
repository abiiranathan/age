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

func TestAgeString(t *testing.T) {
	tests := []struct {
		name     string
		age      Age
		expected string
	}{
		// Years only cases
		{
			name:     "1 year",
			age:      Age{Years: 1, Months: 0, Days: 0},
			expected: "1 year",
		},
		{
			name:     "multiple years",
			age:      Age{Years: 5, Months: 0, Days: 0},
			expected: "5 years",
		},
		{
			name:     "years with months (under 2 years)",
			age:      Age{Years: 1, Months: 6, Days: 0},
			expected: "1 year, 6 months",
		},
		{
			name:     "years without months (2+ years)",
			age:      Age{Years: 2, Months: 3, Days: 0},
			expected: "2 years",
		},

		// Months only cases
		{
			name:     "1 month",
			age:      Age{Years: 0, Months: 1, Days: 0},
			expected: "1 month",
		},
		{
			name:     "multiple months",
			age:      Age{Years: 0, Months: 5, Days: 0},
			expected: "5 months",
		},
		{
			name:     "months with days (under 2 months)",
			age:      Age{Years: 0, Months: 1, Days: 15},
			expected: "1 month, 15 days",
		},
		{
			name:     "months without days (2+ months)",
			age:      Age{Years: 0, Months: 3, Days: 10},
			expected: "3 months",
		},

		// Days only cases
		{
			name:     "0 days (newborn)",
			age:      Age{Years: 0, Months: 0, Days: 0},
			expected: "1 day",
		},
		{
			name:     "1 day",
			age:      Age{Years: 0, Months: 0, Days: 1},
			expected: "1 day",
		},
		{
			name:     "multiple days",
			age:      Age{Years: 0, Months: 0, Days: 15},
			expected: "15 days",
		},

		// Mixed cases
		{
			name:     "1 year, 1 month, 1 day",
			age:      Age{Years: 1, Months: 1, Days: 1},
			expected: "1 year, 1 month",
		},
		{
			name:     "1 month, 1 day",
			age:      Age{Years: 0, Months: 1, Days: 1},
			expected: "1 month, 1 day",
		},
		{
			name:     "complex case with all components",
			age:      Age{Years: 5, Months: 11, Days: 23},
			expected: "5 years",
		},

		// Edge cases
		{
			name:     "maximum values",
			age:      Age{Years: 999, Months: 11, Days: 30},
			expected: "999 years",
		},
		{
			name:     "minimum non-zero values",
			age:      Age{Years: 0, Months: 0, Days: 1},
			expected: "1 day",
		},
		{
			name:     "29 days (almost a month)",
			age:      Age{Years: 0, Months: 0, Days: 29},
			expected: "29 days",
		},
		{
			name:     "11 months (almost a year)",
			age:      Age{Years: 0, Months: 11, Days: 0},
			expected: "11 months",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.age.String(); got != tt.expected {
				t.Errorf("Age.String() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestAgeComparison(t *testing.T) {
	tests := []struct {
		name    string
		age1    Age
		age2    Age
		compare int  // expected Compare() result
		equals  bool // expected Equals() result
		older   bool // expected OlderThan() result
		younger bool // expected YoungerThan() result
	}{
		{
			name:    "equal ages",
			age1:    Age{Years: 5, Months: 3, Days: 10},
			age2:    Age{Years: 5, Months: 3, Days: 10},
			compare: 0,
			equals:  true,
			older:   false,
			younger: false,
		},
		{
			name:    "age1 older by years",
			age1:    Age{Years: 10, Months: 2, Days: 5},
			age2:    Age{Years: 5, Months: 11, Days: 30},
			compare: 1,
			equals:  false,
			older:   true,
			younger: false,
		},
		{
			name:    "age1 older by months",
			age1:    Age{Years: 5, Months: 5, Days: 1},
			age2:    Age{Years: 5, Months: 3, Days: 30},
			compare: 1,
			equals:  false,
			older:   true,
			younger: false,
		},
		{
			name:    "age1 older by days",
			age1:    Age{Years: 5, Months: 3, Days: 15},
			age2:    Age{Years: 5, Months: 3, Days: 10},
			compare: 1,
			equals:  false,
			older:   true,
			younger: false,
		},
		{
			name:    "age1 younger by years",
			age1:    Age{Years: 2, Months: 11, Days: 30},
			age2:    Age{Years: 5, Months: 1, Days: 1},
			compare: -1,
			equals:  false,
			older:   false,
			younger: true,
		},
		{
			name:    "age1 younger by months",
			age1:    Age{Years: 5, Months: 1, Days: 30},
			age2:    Age{Years: 5, Months: 3, Days: 1},
			compare: -1,
			equals:  false,
			older:   false,
			younger: true,
		},
		{
			name:    "age1 younger by days",
			age1:    Age{Years: 5, Months: 3, Days: 5},
			age2:    Age{Years: 5, Months: 3, Days: 10},
			compare: -1,
			equals:  false,
			older:   false,
			younger: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.age1.Compare(tt.age2); got != tt.compare {
				t.Errorf("Compare() = %v, want %v", got, tt.compare)
			}
			if got := tt.age1.Equals(tt.age2); got != tt.equals {
				t.Errorf("Equals() = %v, want %v", got, tt.equals)
			}
			if got := tt.age1.OlderThan(tt.age2); got != tt.older {
				t.Errorf("OlderThan() = %v, want %v", got, tt.older)
			}
			if got := tt.age1.YoungerThan(tt.age2); got != tt.younger {
				t.Errorf("YoungerThan() = %v, want %v", got, tt.younger)
			}
		})
	}
}

func TestThresholdComparisons(t *testing.T) {
	tests := []struct {
		name    string
		age     Age
		years   int
		months  int
		days    int
		atLeast bool
		atMost  bool
	}{
		{
			name:    "exactly at threshold",
			age:     Age{Years: 5, Months: 3, Days: 10},
			years:   5,
			months:  3,
			days:    10,
			atLeast: true,
			atMost:  true,
		},
		{
			name:    "above threshold",
			age:     Age{Years: 6, Months: 0, Days: 0},
			years:   5,
			months:  0,
			days:    0,
			atLeast: true,
			atMost:  false,
		},
		{
			name:    "below threshold",
			age:     Age{Years: 4, Months: 11, Days: 30},
			years:   5,
			months:  0,
			days:    0,
			atLeast: false,
			atMost:  true,
		},
		{
			name:    "complex threshold comparison",
			age:     Age{Years: 2, Months: 8, Days: 15},
			years:   2,
			months:  6,
			days:    0,
			atLeast: true,
			atMost:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.age.AtLeast(tt.years, tt.months, tt.days); got != tt.atLeast {
				t.Errorf("AtLeast() = %v, want %v", got, tt.atLeast)
			}
			if got := tt.age.AtMost(tt.years, tt.months, tt.days); got != tt.atMost {
				t.Errorf("AtMost() = %v, want %v", got, tt.atMost)
			}
		})
	}
}

func TestBetween(t *testing.T) {
	tests := []struct {
		name    string
		age     Age
		younger Age
		older   Age
		want    bool
	}{
		{
			name:    "within range",
			age:     Age{Years: 5, Months: 0, Days: 0},
			younger: Age{Years: 4, Months: 0, Days: 0},
			older:   Age{Years: 6, Months: 0, Days: 0},
			want:    true,
		},
		{
			name:    "at lower boundary",
			age:     Age{Years: 4, Months: 0, Days: 0},
			younger: Age{Years: 4, Months: 0, Days: 0},
			older:   Age{Years: 6, Months: 0, Days: 0},
			want:    true,
		},
		{
			name:    "at upper boundary",
			age:     Age{Years: 6, Months: 0, Days: 0},
			younger: Age{Years: 4, Months: 0, Days: 0},
			older:   Age{Years: 6, Months: 0, Days: 0},
			want:    true,
		},
		{
			name:    "below range",
			age:     Age{Years: 3, Months: 0, Days: 0},
			younger: Age{Years: 4, Months: 0, Days: 0},
			older:   Age{Years: 6, Months: 0, Days: 0},
			want:    false,
		},
		{
			name:    "above range",
			age:     Age{Years: 7, Months: 0, Days: 0},
			younger: Age{Years: 4, Months: 0, Days: 0},
			older:   Age{Years: 6, Months: 0, Days: 0},
			want:    false,
		},
		{
			name:    "complex range check",
			age:     Age{Years: 2, Months: 8, Days: 15},
			younger: Age{Years: 2, Months: 6, Days: 0},
			older:   Age{Years: 3, Months: 0, Days: 0},
			want:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.age.Between(tt.younger, tt.older); got != tt.want {
				t.Errorf("Between() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTotalDays(t *testing.T) {
	tests := []struct {
		name string
		age  Age
		want int
	}{
		{
			name: "years only",
			age:  Age{Years: 5, Months: 0, Days: 0},
			want: 5 * 365,
		},
		{
			name: "months only",
			age:  Age{Years: 0, Months: 6, Days: 0},
			want: 6 * 30,
		},
		{
			name: "days only",
			age:  Age{Years: 0, Months: 0, Days: 15},
			want: 15,
		},
		{
			name: "mixed",
			age:  Age{Years: 2, Months: 3, Days: 10},
			want: 2*365 + 3*30 + 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.age.TotalDays(); got != tt.want {
				t.Errorf("TotalDays() = %v, want %v", got, tt.want)
			}
		})
	}
}
