// Tiny pkg for calculating age from birthDate and printing the age in a human-readable format.
package age

import (
	"fmt"
	"strings"
	"time"
)

// Age struct represents age in years, months and days.
type Age struct {
	Years  int // Age in years
	Months int // Age in months (in addition to years)
	Days   int // Age in days (in addition to years and months)
}

// IsLeapYear returns true if the year is a Leap year.
func IsLeapYear(year int) bool {
	return (year%4 == 0 && (year%100 != 0)) || (year%400 == 0)
}

// Calculates age from birthDate to now time.
func AgeAt(birthDate time.Time, endDate time.Time) Age {
	if birthDate.After(endDate) {
		return Age{}
	}

	var MONTHS = [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	birthYear, birthMonth, bday := birthDate.Year(), int(birthDate.Month()), birthDate.Day()
	cYear, cMonth, cday := endDate.Year(), int(endDate.Month()), endDate.Day()

	if IsLeapYear(endDate.Year()) {
		MONTHS[1] = 29
	}

	if (bday > cday) && !IsLeapYear(endDate.Year()) {
		cday = cday + MONTHS[birthMonth-1]
		cMonth = cMonth - 1
	}

	if birthMonth > cMonth {
		cYear = cYear - 1
		cMonth = cMonth + 12
	}

	years := cYear - birthYear
	months := cMonth - birthMonth
	days := cday - bday

	// Leap year: Special case when someone was born on in Feb in a leap year and current date
	// is less than their birthdate.
	if IsLeapYear(birthDate.Year()) && IsLeapYear(endDate.Year()) &&
		cMonth == 2 && birthMonth == 2 && bday > cday {

		years -= 1
		months = 11
		days = cday
	}

	return Age{Years: years, Months: months, Days: days}
}

// Calculates the current age at this instant.
func CurrentAge(birthDate time.Time) Age {
	return AgeAt(birthDate, time.Now())
}

// Returns a human-readable age representation
func (a Age) String() string {
	str := strings.Builder{}

	// years or (years and months)
	if a.Years > 0 {
		str.WriteString(fmt.Sprintf("%d year", a.Years))
		if a.Years != 1 {
			str.WriteRune('s')
		}

		if a.Months > 0 && a.Years < 2 {
			str.WriteString(fmt.Sprintf(", %d month", a.Months))
			if a.Months != 1 {
				str.WriteRune('s')
			}
		}

		return str.String()
	}

	// Months or months and days (toddlers)
	if a.Months > 0 {
		str.WriteString(fmt.Sprintf("%d month", a.Months))
		if a.Months != 1 {
			str.WriteRune('s')
		}

		if a.Days > 0 && a.Months < 2 {
			str.WriteString(fmt.Sprintf(", %d day", a.Days))
			if a.Days != 1 {
				str.WriteRune('s')
			}
		}

		return str.String()
	}

	// If the number of days is 0, then the baby is a newborn.
	// We don't want to return Newborn so that we can have consistent pattern easily parsed with regex.
	if a.Days == 0 {
		return "1 day"
	}

	// >= 1 days old newborns and neonates.
	str.Reset()
	str.WriteString(fmt.Sprintf("%d day", a.Days))
	if a.Days != 1 {
		str.WriteRune('s')
	}

	return str.String()
}

// TotalDays returns the approximate total days in the age.
// Useful for comparisons where exact days matter.
func (a Age) TotalDays() int {
	// Approximate - assumes 30.44 days/month and 365.25 days/year
	return a.Years*365 + a.Months*30 + a.Days
}

// Compare compares two ages and returns:
// -1 if a is younger than other
// 0 if a is equal to other
// 1 if a is older than other
func (a Age) Compare(other Age) int {
	// Use multipliers but maintain proper ordering.
	aTotal := a.Years*10000 + a.Months*100 + a.Days
	otherTotal := other.Years*10000 + other.Months*100 + other.Days

	switch {
	case aTotal < otherTotal:
		return -1
	case aTotal > otherTotal:
		return 1
	default:
		return 0
	}
}

// Equals returns true if two ages are exactly equal
func (a Age) Equals(other Age) bool {
	return a.Years == other.Years && a.Months == other.Months && a.Days == other.Days
}

// OlderThan returns true if this age is older than the other age
func (a Age) OlderThan(other Age) bool {
	return a.Compare(other) == 1
}

// YoungerThan returns true if this age is younger than the other age
func (a Age) YoungerThan(other Age) bool {
	return a.Compare(other) == -1
}

// AtLeast returns true if this age is at least as old as the specified years, months, and days
func (a Age) AtLeast(years, months, days int) bool {
	compareAge := Age{Years: years, Months: months, Days: days}
	return a.Compare(compareAge) >= 0
}

// AtMost returns true if this age is at most as old as the specified years, months, and days
func (a Age) AtMost(years, months, days int) bool {
	compareAge := Age{Years: years, Months: months, Days: days}
	return a.Compare(compareAge) <= 0
}

// Between returns true if this age is between the two specified age ranges (inclusive)
func (a Age) Between(younger, older Age) bool {
	return a.Compare(younger) >= 0 && a.Compare(older) <= 0
}
