package yahoo

import (
	"io/ioutil"
	"testing"

	"github.com/shuLhan/gontacts"
	"github.com/shuLhan/gontacts/proofn"
)

const (
	sampleContact = "samples/contact.json"
)

var (
	gotContact *Contact
)

func parseContact(t *testing.T) (contact *Contact) {
	jsonb, err := ioutil.ReadFile(sampleContact)
	if err != nil {
		t.Fatal(err)
	}

	contact, err = NewContact(jsonb)
	if err != nil {
		t.Fatal(err)
	}

	return
}

func TestNewContact(t *testing.T) {
	exp := &Contact{
		Name: &gontacts.Name{
			Given:  "Test",
			Middle: "Middle",
			Family: "Proofn",
		},
		Birthday: &gontacts.Date{
			Day:   "24",
			Month: "1",
			Year:  "1980",
		},
		Email: []gontacts.Email{{
			Address: "test@proofn.com",
		}},
		Phone: []gontacts.Phone{{
			Type:   "home",
			Number: "084-563-21",
		}, {
			Type:   "mobile",
			Number: "084-563-20",
		}, {
			Type:   "work",
			Number: "084-563-23",
		}},
		Link: []string{
			"www.proofn.com",
		},
		Company:  "Myabuy",
		JobTitle: "Tester",
	}

	gotContact = parseContact(t)

	assert(t, exp.Name, gotContact.Name, true)
	assert(t, exp.Birthday, gotContact.Birthday, true)
	assert(t, exp.Address, gotContact.Address, true)
	assert(t, exp.Anniversary, gotContact.Anniversary, true)
	assert(t, exp.Email, gotContact.Email, true)
	assert(t, exp.Phone, gotContact.Phone, true)
	assert(t, exp.Link, gotContact.Link, true)
	assert(t, exp.Company, gotContact.Company, true)
	assert(t, exp.JobTitle, gotContact.JobTitle, true)
}

func TestToProofn(t *testing.T) {
	exp := proofn.Contact{
		FirstName:   "Test",
		MiddleName:  "Middle",
		LastName:    "Proofn",
		FullName:    "Test Middle Proofn",
		Birthday:    "1980-1-24",
		Email:       "test@proofn.com",
		PhoneNumber: "084-563-21",
		WorkPhone:   "084-563-23",
		MobilePhone: "084-563-20",
		Office:      "Myabuy",
		JobTitle:    "Tester",
	}

	if gotContact == nil {
		gotContact = parseContact(t)
	}

	proofnContact := gotContact.ToProofn()

	assert(t, exp, proofnContact, true)
}
