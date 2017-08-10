package yahoo

import (
	"github.com/json-iterator/go"

	"github.com/shuLhan/gontacts"
	"github.com/shuLhan/gontacts/proofn"
)

//
// Contact define the contact item in response.
//
type Contact struct {
	Fields []Field `json:"fields"`

	// Decoded fields
	Name        *gontacts.Name
	Birthday    *gontacts.Date
	Anniversary *gontacts.Date
	Address     []gontacts.Address
	Email       []gontacts.Email
	Phone       []gontacts.Phone
	Link        []string
	Company     string
	JobTitle    string

	// Ignored fields for speedup.
	//ID           int        `json:"id"`
	//IsConnection bool       `json:"isConnection"`
	//Error        int        `json:"error"`
	//RestoredID   int        `json:"restoredId"`
	//Categories []Category `json:"categories"`
	//Meta
}

func (c *Contact) decodeAddress(typ string, vmap map[string]interface{}) (
	adr gontacts.Address,
) {
	adr = gontacts.Address{
		Type:        typ,
		Street:      vmap["street"].(string),
		City:        vmap["city"].(string),
		StateOrProv: vmap["stateOrProvince"].(string),
		PostalCode:  vmap["postalCode"].(string),
	}
	if vmap["country"] != nil {
		adr.Country = vmap["country"].(string)
	}

	return
}

//
// decodeField will transform Field Value into a struct or string, based on
// Field Type.
//
func (c *Contact) decodeField(field *Field) {
	vmap, vstr, ok := field.getValueType()

	if !ok {
		return
	}

	typ := field.getFlag()

	switch field.Type {
	case "address":
		adr := c.decodeAddress(typ, vmap)
		c.Address = append(c.Address, adr)

	case "birthday":
		c.Birthday = &gontacts.Date{
			Day:   vmap["day"].(string),
			Month: vmap["month"].(string),
			Year:  vmap["year"].(string),
		}

	case "company":
		c.Company = vstr

	case "email":
		c.Email = append(c.Email, gontacts.Email{
			Type:    typ,
			Address: vstr,
		})

	case "jobTitle":
		c.JobTitle = vstr

	case "link":
		c.Link = append(c.Link, vstr)

	case "name":
		c.Name = &gontacts.Name{
			Given:       vmap["givenName"].(string),
			Middle:      vmap["middleName"].(string),
			Family:      vmap["familyName"].(string),
			Prefix:      vmap["prefix"].(string),
			Suffix:      vmap["suffix"].(string),
			GivenSound:  vmap["givenNameSound"].(string),
			FamilySound: vmap["familyNameSound"].(string),
		}

	case "phone":
		c.Phone = append(c.Phone, gontacts.Phone{
			Type:   typ,
			Number: vstr,
		})
	}
}

//
// decodeFields will convert the interface value in each field into its struct
// representation.
//
func (c *Contact) decodeFields() {
	for x, field := range c.Fields {
		c.decodeField(&field)

		// Clear the Value to minimize memory usage.
		c.Fields[x].Value = nil
	}
}

//
// NewContact will parse JSON input and return Contact object on success.
//
// On fail it will return nil and error.
//
func NewContact(jsonb []byte) (contact *Contact, err error) {
	contact = &Contact{}

	err = jsoniter.Unmarshal(jsonb, contact)
	if err != nil {
		return
	}

	contact.decodeFields()

	return
}

//
// ToProofn will convert Yahoo contact structure to Proofn contact structure.
//
func (c *Contact) ToProofn() (proofnContact proofn.Contact) {
	for _, field := range c.Fields {
		switch field.Type {
		case "address":
			proofnContact.SetAddresses(c.Address)

		case "birthday":
			if c.Birthday != nil {
				proofnContact.Birthday = c.Birthday.String()
			}

		case "company":
			proofnContact.Office = c.Company

		case "email":
			proofnContact.SetEmails(c.Email)

		case "jobTitle":
			proofnContact.JobTitle = c.JobTitle

		case "link":

		case "name":
			proofnContact.SetName(c.Name)

		case "phone":
			proofnContact.SetPhones(c.Phone)
		}
	}

	return
}
