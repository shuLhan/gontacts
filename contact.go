package gontacts

import (
	"strings"
)

//
// Contact define a single contact entity with sane format.
//
type Contact struct {
	Name        Name
	Birthday    *Date
	Anniversary *Date
	Addresses   []Address
	Emails      []Email
	Phones      []Phone
	Links       []string
	Notes       []string
	Company     string
	JobTitle    string
}

//
// SetBirthday will set contact birthday from string format "YYYY-MM-DD" or
// "YYYY-MM-DDTHH:MM:SSZ".
//
// (1) Split by zone first, and then
// (2) split the date.
//
func (contact *Contact) SetBirthday(dateStr string) {
	if dateStr == "" {
		return
	}

	// (1)
	dateTimeZone := strings.Split(dateStr, "T")

	// (2)
	dates := strings.Split(dateTimeZone[0], "-")
	if len(dates) != 3 {
		return
	}

	contact.Birthday = &Date{
		Year:  dates[0],
		Month: dates[1],
		Day:   dates[2],
	}
}

//
// SetAnniversary will set contact annivery from string format "YYYY-MM-DD".
//
func (contact *Contact) SetAnniversary(dateStr string) {
	if dateStr == "" {
		return
	}

	dates := strings.Split(dateStr, "-")
	if len(dates) != 3 {
		return
	}

	contact.Anniversary = &Date{
		Year:  dates[0],
		Month: dates[1],
		Day:   dates[2],
	}
}
