package yahoo

import (
	"io/ioutil"
	"testing"
)

const (
	sampleContacts = "samples/contacts.json"
)

func TestImportFromJSON(t *testing.T) {
	contactsb, err := ioutil.ReadFile(sampleContacts)
	if err != nil {
		t.Fatal(err)
	}

	contacts, err := ImportFromJSON(contactsb)
	if err != nil {
		t.Fatal(err)
	}

	assert(t, 54, len(contacts), true)
}
