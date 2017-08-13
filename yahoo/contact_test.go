package yahoo

import (
	"io/ioutil"
	"testing"

	"github.com/shuLhan/gontacts"
)

const (
	sampleContact = "samples/contact.json"
)

var (
	gotContact *gontacts.Contact
)

func parseSampleJSON(t *testing.T, input string) (contact *gontacts.Contact) {
	jsonb, err := ioutil.ReadFile(input)
	if err != nil {
		t.Fatal(err)
	}

	contact, err = ParseJSON(jsonb)
	if err != nil {
		t.Fatal(err)
	}

	return
}

func TestParseJSON(t *testing.T) {
	exp := &gontacts.Contact{
		Name: gontacts.Name{
			Given:  "Test",
			Middle: "Middle",
			Family: "Proofn",
		},
		Birthday: &gontacts.Date{
			Day:   "24",
			Month: "1",
			Year:  "1980",
		},
		Emails: []gontacts.Email{{
			Address: "test@proofn.com",
		}},
		Phones: []gontacts.Phone{{
			Type:   "home",
			Number: "084-563-21",
		}, {
			Type:   "mobile",
			Number: "084-563-20",
		}, {
			Type:   "work",
			Number: "084-563-23",
		}},
		Links: []string{
			"www.proofn.com",
		},
		Company:  "Myabuy",
		JobTitle: "Tester",
	}

	gotContact = parseSampleJSON(t, sampleContact)

	assert(t, exp.Name, gotContact.Name, true)
	assert(t, exp.Birthday, gotContact.Birthday, true)
	assert(t, exp.Addresses, gotContact.Addresses, true)
	assert(t, exp.Anniversary, gotContact.Anniversary, true)
	assert(t, exp.Emails, gotContact.Emails, true)
	assert(t, exp.Phones, gotContact.Phones, true)
	assert(t, exp.Links, gotContact.Links, true)
	assert(t, exp.Company, gotContact.Company, true)
	assert(t, exp.JobTitle, gotContact.JobTitle, true)
}
