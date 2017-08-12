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
	Company     string
	JobTitle    string
}

//
// SetBirthday will set contact birthday from string format "YYYY-MM-DD".
//
func (contact *Contact) SetBirthday(dateStr string) {
	if dateStr == "" {
		return
	}

	dates := strings.Split(dateStr, "-")
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
