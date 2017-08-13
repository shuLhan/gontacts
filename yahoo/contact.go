package yahoo

import (
	"github.com/json-iterator/go"

	"github.com/shuLhan/gontacts"
)

//
// Contact define the contact item in response.
//
type Contact struct {
	Fields []Field `json:"fields"`

	// Ignored fields for speedup.
	//ID           int        `json:"id"`
	//IsConnection bool       `json:"isConnection"`
	//Error        int        `json:"error"`
	//RestoredID   int        `json:"restoredId"`
	//Categories []Category `json:"categories"`
	//Meta
}

//
// Decode will convert the interface value in each field into its struct
// representation.
//
func (c *Contact) Decode() (to *gontacts.Contact) {
	to = &gontacts.Contact{}

	for x, field := range c.Fields {
		field.Decode(to)

		// Clear the Value to minimize memory usage.
		c.Fields[x].Value = nil
	}

	return
}

//
// ParseJSON will parse JSON input and return gontacts.Contact object on
// success.
//
// On fail it will return nil and error.
//
func ParseJSON(jsonb []byte) (to *gontacts.Contact, err error) {
	ycontact := &Contact{}

	err = jsoniter.Unmarshal(jsonb, ycontact)
	if err != nil {
		return
	}

	to = ycontact.Decode()

	return
}
