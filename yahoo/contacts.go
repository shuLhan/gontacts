package yahoo

import (
	"github.com/json-iterator/go"

	"github.com/shuLhan/gontacts/proofn"
)

//
// Contacts define the holder for root of contacts response.
//
type Contacts struct {
	Contact []Contact `json:"contact"`
	Start   int       `json:"start"`
	Count   int       `json:"count"`
	Total   int       `json:"total"`
	URI     string    `json:"uri"`
}

type _Root struct {
	Contacts Contacts `json:"contacts"`
}

//
// NewContacts will parse JSON input and return Contacts object on success.
//
// On fail it will return nil and error.
//
func NewContacts(jsonb []byte) (contacts *Contacts, err error) {
	root := &_Root{}

	err = jsoniter.Unmarshal(jsonb, root)
	if err != nil {
		return
	}

	contacts = &root.Contacts

	for x := range contacts.Contact {
		contacts.Contact[x].decodeFields()
	}

	return
}

//
// ToProofn will convert contacts structure to Proofn format.
//
func (contacts *Contacts) ToProofn() (proofns []proofn.Contact) {
	for _, contact := range contacts.Contact {
		proofn := contact.ToProofn()
		proofns = append(proofns, proofn)
	}

	return
}
