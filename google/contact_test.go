package google

import (
	"io/ioutil"
	"testing"

	"github.com/json-iterator/go"

	"github.com/shuLhan/gontacts"
)

const (
	sampleContact = "samples/contact.json"
)

var (
	gotContact *gontacts.Contact
)

func parseContact(t *testing.T) (contact *gontacts.Contact) {
	googleContact := &Contact{}

	jsonb, err := ioutil.ReadFile(sampleContact)
	if err != nil {
		t.Fatal(err)
	}

	err = jsoniter.Unmarshal(jsonb, googleContact)
	if err != nil {
		t.Fatal(err)
	}

	return googleContact.Decode()
}

func TestDecode(t *testing.T) {
	exp := &gontacts.Contact{
		Name: gontacts.Name{
			Given:  "Test",
			Middle: "Middle",
			Family: "Last",
			Prefix: "Prefix",
			Suffix: "Suffix",
		},
		Birthday: &gontacts.Date{
			Day:   "30",
			Month: "01",
			Year:  "1980",
		},
		Anniversary: &gontacts.Date{
			Day:   "20",
			Month: "11",
			Year:  "2016",
		},
		Addresses: []gontacts.Address{
			gontacts.Address{
				Type:        "home",
				POBox:       "40124",
				Street:      "Jl. Tubagus Ismail VI",
				City:        "Bandung",
				StateOrProv: "Jabar",
				PostalCode:  "40124",
				Country:     "Indonesia",
			},
			gontacts.Address{
				Type:   "work",
				Street: "Perumahan Delima Cikutra",
			},
		},
		Emails: []gontacts.Email{{
			Type:    "home",
			Address: "first.tester@proofn.com",
		}, {
			Type:    "work",
			Address: "work@proofn.com",
		}},
		Phones: []gontacts.Phone{{
			Type:   "mobile",
			Number: "856123456789",
		}, {
			Type:   "work",
			Number: "2233445566",
		}, {
			Type:   "home",
			Number: "9999999",
		}, {
			Type:   "main",
			Number: "8888888",
		}},
		Links: []string{
			"https://www.proofn.com",
		},
		Company:  "Myabuy",
		JobTitle: "Devops",
	}

	gotContact = parseContact(t)

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
