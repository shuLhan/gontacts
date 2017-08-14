package microsoft

import (
	"io/ioutil"
	"reflect"
	"runtime/debug"
	"testing"

	"github.com/shuLhan/gontacts"
)

const (
	sampleContacts = "samples/contacts.json"
)

func assert(t *testing.T, exp, got interface{}, equal bool) {
	if reflect.DeepEqual(exp, got) == equal {
		return
	}

	debug.PrintStack()

	t.Fatalf("\n"+
		">>> Expecting '%v'\n"+
		"          got '%v'\n", exp, got)
}

func TestImportFromJSON(t *testing.T) {
	exp := &gontacts.Contact{
		Name: gontacts.Name{
			Given:  "First",
			Middle: "Middle",
			Family: "Tester",
			Prefix: "Prof.",
			Suffix: "AMD",
		},
		Birthday: &gontacts.Date{
			Year:  "1984",
			Month: "08",
			Day:   "14",
		},
		Addresses: []gontacts.Address{{
			Type:        gontacts.TypeHome,
			Street:      "Tubagus Ismail VI",
			City:        "Bandung",
			StateOrProv: "JABAR",
			PostalCode:  "40124",
			Country:     "Indonesia",
		}, {
			Type:   gontacts.TypeWork,
			Street: "Cikutra",
			City:   "Bandung",
		}},
		Emails: []gontacts.Email{{
			Type:    gontacts.TypeMain,
			Address: "first.tester@proofn.com",
		}, {
			Type:    gontacts.TypeHome,
			Address: "tester@proofn.com",
		}},
		Phones: []gontacts.Phone{{
			Type:   gontacts.TypeHome,
			Number: "+22808080",
		}, {
			Type:   gontacts.TypeMobile,
			Number: "+62856123456789",
		}, {
			Type:   gontacts.TypeWork,
			Number: "+22909090",
		}},
		Notes: []string{
			"This is a note.",
		},
		Company:  "Myabuy",
		JobTitle: "Tester",
	}

	jsonb, err := ioutil.ReadFile(sampleContacts)
	if err != nil {
		t.Fatal(err)
	}

	microsoftClient := NewClient("", "", "")

	contacts, err := microsoftClient.ImportFromJSON(jsonb)
	if err != nil {
		t.Fatal(err)
	}

	assert(t, 1, len(contacts), true)

	got := contacts[0]

	assert(t, exp, got, true)
}
