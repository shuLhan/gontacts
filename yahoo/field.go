package yahoo

import (
	"strings"

	"github.com/shuLhan/gontacts"
)

const (
	// List of field type.
	// nolint: golint
	FieldTypeAddress  = "address"
	FieldTypeBirthday = "birthday"
	FieldTypeCompany  = "company"
	FieldTypeEmail    = "email"
	FieldTypeJobTitle = "jobTitle"
	FieldTypeLink     = "link"
	FieldTypeName     = "name"
	FieldTypePhone    = "phone"
)

//
// Field define a composite attribute in Contact.
// Known value for Type: "phone", "name", "address".
//
type Field struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
	Flags []string    `json:"flags"`

	// Ignored fields for speedup.
	//ID         int         `json:"id"`
	//EditedBy   string      `json:"editedBy"`
	//Categories []string    `json:"categories"`
	//Meta
}

//
// getValueType will return the Go type of field's Value.
//
func (field *Field) getValueType() (
	vmap map[string]interface{},
	vstr string,
	ok bool,
) {
	ok = true

	switch v := field.Value.(type) {
	case map[string]interface{}:
		vmap = v
	case string:
		vstr = v
	default:
		ok = false
	}

	return
}

//
// getFlag will return the first flag or empty string if flags fields is empty.
//
func (field *Field) getFlag() string {
	if len(field.Flags) > 0 {
		return strings.ToLower(field.Flags[0])
	}
	return ""
}

//
// decodeAddress will convert Yahoo address format to gontacts address format.
//
func (field *Field) decodeAddress(flag string, vmap map[string]interface{}) (
	adr gontacts.Address,
) {
	adr = gontacts.Address{
		Type:        flag,
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
// Decode will convert Yahoo' contact field value and save it to gontacts
// Contact format.
//
func (field *Field) Decode(to *gontacts.Contact) {
	vmap, vstr, ok := field.getValueType()

	if !ok {
		return
	}

	flag := field.getFlag()

	switch field.Type {
	case FieldTypeAddress:
		adr := field.decodeAddress(flag, vmap)
		to.Addresses = append(to.Addresses, adr)

	case FieldTypeBirthday:
		to.Birthday = &gontacts.Date{
			Day:   vmap["day"].(string),
			Month: vmap["month"].(string),
			Year:  vmap["year"].(string),
		}

	case FieldTypeCompany:
		to.Company = vstr

	case FieldTypeEmail:
		to.Emails = append(to.Emails, gontacts.Email{
			Type:    flag,
			Address: vstr,
		})

	case FieldTypeJobTitle:
		to.JobTitle = vstr

	case FieldTypeLink:
		to.Links = append(to.Links, vstr)

	case FieldTypeName:
		to.Name = gontacts.Name{
			Given:       vmap["givenName"].(string),
			Middle:      vmap["middleName"].(string),
			Family:      vmap["familyName"].(string),
			Prefix:      vmap["prefix"].(string),
			Suffix:      vmap["suffix"].(string),
			GivenSound:  vmap["givenNameSound"].(string),
			FamilySound: vmap["familyNameSound"].(string),
		}

	case FieldTypePhone:
		to.Phones = append(to.Phones, gontacts.Phone{
			Type:   flag,
			Number: vstr,
		})
	}
}
