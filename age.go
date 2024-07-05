// Tiny pkg for calculating age from birthDate and printing the age in a human-readable format.
package age

import (
	"fmt"
	"strings"
	"time"
)

type Age struct {
	Years  int
	Months int
	Days   int
}

func isLeap(date time.Time) bool {
	year := date.Year()
	return (year%4 == 0 && (year%100 != 0)) || (year%400 == 0)
}

// Calculates age from birthDate to now time.
func AgeAt(birthDate time.Time, endDate time.Time) Age {
	var MONTHS = [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

	birthYear, birthMonth, bday := birthDate.Year(), int(birthDate.Month()), birthDate.Day()
	cYear, cMonth, cday := endDate.Year(), int(endDate.Month()), endDate.Day()

	if isLeap(endDate) {
		MONTHS[1] = 29
	}

	if (bday > cday) && !isLeap(endDate) {
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
	if isLeap(birthDate) && isLeap(endDate) && cMonth == 2 && birthMonth == 2 && bday > cday {
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
			if a.Months != 1 {
				str.WriteRune('s')
			}
		}

		return str.String()
	}

	// If the number of days is 0, then the baby is a newborn.
	if a.Days == 0 {
		return "1 day"
	}

	// Newborns(0 days) && Days old (neonates)
	str.WriteString(fmt.Sprintf("%d day", a.Days))
	if a.Days != 1 {
		str.WriteRune('s')
	}

	return str.String()
}
