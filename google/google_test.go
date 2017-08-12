package google

import (
	"io/ioutil"
	"reflect"
	"runtime/debug"
	"testing"
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

func TestNewContacts(t *testing.T) {
	jsonb, err := ioutil.ReadFile(sampleContacts)
	if err != nil {
		t.Fatal(err)
	}

	contacts, err := NewContacts(jsonb)
	if err != nil {
		t.Fatal(err)
	}

	assert(t, 55, len(contacts), true)
}
