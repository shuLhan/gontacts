package google

import (
	"github.com/shuLhan/gontacts"
)

//
// Contact define a single Google contact data.
//
// Some of the fields are disabled for speed.
//
type Contact struct {
	//ID         GD
	//ETag       string     `json:"gd$etag,omitempty"`
	//Updated    GD         `json:"updated,omitempty"`
	//Edited     GD         `json:"app$edited,omitempty"`
	//Categories []Category `json:"category,omitempty"`
	//Title      GD         `json:"title,omitempty"`
	//Links      []Link     `json:"link,omitempty"`

	Name      Name      `json:"gd$name,omitempty"`
	Birthday  Birthday  `json:"gContact$birthday,omitempty"`
	Orgs      []Org     `json:"gd$organization,omitempty"`
	Emails    []Email   `json:"gd$email,omitempty"`
	Phones    []Phone   `json:"gd$phoneNumber,omitempty"`
	Addresses []Address `json:"gd$structuredPostalAddress,omitempty"`
	Events    []Event   `json:"gContact$event,omitempty"`
	Websites  []Link    `json:"gContact$website,omitempty"`
}

func (gc *Contact) decodeOrg(contact *gontacts.Contact) {
	if len(gc.Orgs) == 0 {
		return
	}

	contact.Company = gc.Orgs[0].Name.Value
	contact.JobTitle = gc.Orgs[0].JobTitle.Value
}

func (gc *Contact) decodeEmails(contact *gontacts.Contact) {
	for _, email := range gc.Emails {
		decodedEmail := gontacts.Email{
			Type:    ParseRel(email.Rel),
			Address: email.Address,
		}
		contact.Emails = append(contact.Emails, decodedEmail)
	}
}

func (gc *Contact) decodePhones(contact *gontacts.Contact) {
	for _, phone := range gc.Phones {
		decodedPhone := gontacts.Phone{
			Type:   ParseRel(phone.Rel),
			Number: phone.Number,
		}
		contact.Phones = append(contact.Phones, decodedPhone)
	}
}

func (gc *Contact) decodeAddresses(contact *gontacts.Contact) {
	for _, adr := range gc.Addresses {
		decAdr := gontacts.Address{
			Type:        ParseRel(adr.Rel),
			POBox:       adr.POBox.Value,
			Street:      adr.Street.Value,
			City:        adr.City.Value,
			StateOrProv: adr.StateOrProv.Value,
			PostalCode:  adr.PostalCode.Value,
			Country:     adr.Country.Value,
		}

		contact.Addresses = append(contact.Addresses, decAdr)
	}
}

func (gc *Contact) decodeEvents(contact *gontacts.Contact) {
	for _, event := range gc.Events {
		if event.Rel == gontacts.TypeAnniversary {
			contact.SetAnniversary(event.When.Start)
		}
	}
}

func (gc *Contact) decodeWebsites(contact *gontacts.Contact) {
	for _, site := range gc.Websites {
		contact.Links = append(contact.Links, site.HRef)
	}
}

//
// Decode will convert Google's Contact to our Contact format.
//
func (gc *Contact) Decode() (contact *gontacts.Contact) {
	contact = &gontacts.Contact{
		Name: gontacts.Name{
			Given:  gc.Name.First.Value,
			Middle: gc.Name.Middle.Value,
			Family: gc.Name.Last.Value,
			Prefix: gc.Name.Prefix.Value,
			Suffix: gc.Name.Suffix.Value,
		},
	}

	contact.SetBirthday(gc.Birthday.When)

	gc.decodeOrg(contact)
	gc.decodeEmails(contact)
	gc.decodePhones(contact)
	gc.decodeAddresses(contact)
	gc.decodeEvents(contact)
	gc.decodeWebsites(contact)

	return
}
