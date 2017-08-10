package yahoo

import (
	"io/ioutil"
	"testing"
)

const (
	sampleContacts = "samples/contacts.json"
)

var contacts *Contacts

func TestNewContacts(t *testing.T) {
	var err error

	contactsb, err := ioutil.ReadFile(sampleContacts)
	if err != nil {
		t.Fatal(err)
	}

	contacts, err = NewContacts(contactsb)
	if err != nil {
		t.Fatal(err)
	}

	assert(t, 0, contacts.Start, true)
	assert(t, 54, contacts.Count, true)
	assert(t, 54, contacts.Total, true)

	contacts.ToProofn()
}
