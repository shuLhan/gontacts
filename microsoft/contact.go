package microsoft

import (
	"github.com/shuLhan/gontacts"
)

//
// Contact define Microsoft Live's contact format.
//
// Some of the fields are disabled for speed up.
//
type Contact struct {
	//ETag string `json:"@odata.etag,omitempty"`
	//Id   string `json:"id,omitempty"`
	//Created string `json:"createdDateTime,omitempty"`
	//LastModified string `json:"lastModifiedDateTime,omitempty"`
	//ChangeKey string `json:"changeKey,omitempty"`
	//Categories []string `json:"categories,omitempty"`
	//ParentFolderID string `json:"parentFolderId,omitempty"`
	//FileAs string `json:"fileAs,omitempty"`

	Birthday string `json:"birthday,omitempty"`

	DisplayName string `json:"displayName,omitempty"`
	GivenName   string `json:"givenName,omitempty"`
	Initials    string `json:"initials,omitempty"`
	MiddleName  string `json:"middleName,omitempty"`
	NickName    string `json:"nickName,omitempty"`
	SurName     string `json:"surname,omitempty"`
	Title       string `json:"title,omitempty"`
	Generation  string `json:"generation,omitempty"`

	//YomiGivenName string `json:"yomiGivenName,omitempty"`
	//YomiSurname string `json:"yomiSurname,omitempty"`
	//YomiCompanyName string `json:"yomiCompanyName,omitempty"`

	IMAddresses []string `json:"imAddresses,omitempty"`

	JobTitle         string `json:"jobTitle,omitempty"`
	Company          string `json:"companyName,omitempty"`
	Department       string `json:"department,omitempty"`
	OfficeLocation   string `json:"officeLocation,omitempty"`
	Profession       string `json:"profession,omitempty"`
	BusinessHomePage string `json:"businessHomePage,omitempty"`
	AssistantName    string `json:"assistantName,omitempty"`
	Manager          string `json:"manager,omitempty"`

	HomePhones     []string `json:"homePhones,omitempty"`
	MobilePhone    string   `json:"mobilePhone,omitempty"`
	BusinessPhones []string `json:"businessPhones,omitempty"`

	SpouseName    string   `json:"spouseName,omitempty"`
	PersonalNotes string   `json:"personalNotes,omitempty"`
	Children      []string `json:"children,omitempty"`

	Emails []Email `json:"emailAddresses,omitempty"`

	HomeAddress     Address `json:"homeAddress,omitempty"`
	BusinessAddress Address `json:"businessAddress,omitempty"`
	OtherAddress    Address `json:"otherAddress,omitempty"`
}

func (contact *Contact) decodeEmails(to *gontacts.Contact) {
	var flag string

	for x, email := range contact.Emails {
		switch x {
		case 0:
			flag = gontacts.TypeMain
		case 1:
			flag = gontacts.TypeHome
		case 2:
			flag = gontacts.TypeWork
		default:
			flag = gontacts.TypeOther
		}

		to.Emails = append(to.Emails, gontacts.Email{
			Type:    flag,
			Address: email.Address,
		})
	}
}

func (contact *Contact) decodePhones(to *gontacts.Contact) {
	if len(contact.HomePhones) > 0 {
		to.Phones = append(to.Phones, gontacts.Phone{
			Type:   gontacts.TypeHome,
			Number: contact.HomePhones[0],
		})
	}

	if contact.MobilePhone != "" {
		to.Phones = append(to.Phones, gontacts.Phone{
			Type:   gontacts.TypeMobile,
			Number: contact.MobilePhone,
		})
	}

	if len(contact.BusinessPhones) > 0 {
		to.Phones = append(to.Phones, gontacts.Phone{
			Type:   gontacts.TypeWork,
			Number: contact.BusinessPhones[0],
		})
	}
}

func (contact *Contact) decodeLinks(to *gontacts.Contact) {
	if len(contact.IMAddresses) > 0 {
		to.Links = append(to.Links, contact.IMAddresses...)
	}

	if contact.BusinessHomePage != "" {
		to.Links = append(to.Links, contact.BusinessHomePage)
	}
}

func (contact *Contact) decodeNotes(to *gontacts.Contact) {
	if contact.PersonalNotes != "" {
		to.Notes = append(to.Notes, contact.PersonalNotes)
	}
}

//
// Decode will convert Microsoft's Contact to our Contact format.
//
func (contact *Contact) Decode() (to *gontacts.Contact) {
	to = &gontacts.Contact{
		Name: gontacts.Name{
			Given:  contact.GivenName,
			Middle: contact.MiddleName,
			Family: contact.SurName,
			Prefix: contact.Title,
			Suffix: contact.Generation,
		},
		Addresses: []gontacts.Address{{
			Type:        "home",
			Street:      contact.HomeAddress.Street,
			City:        contact.HomeAddress.City,
			StateOrProv: contact.HomeAddress.State,
			PostalCode:  contact.HomeAddress.PostalCode,
			Country:     contact.HomeAddress.Country,
		}, {
			Type:        "work",
			Street:      contact.BusinessAddress.Street,
			City:        contact.BusinessAddress.City,
			StateOrProv: contact.BusinessAddress.State,
			PostalCode:  contact.BusinessAddress.PostalCode,
			Country:     contact.BusinessAddress.Country,
		}},
		Company:  contact.Company,
		JobTitle: contact.JobTitle,
	}

	to.SetBirthday(contact.Birthday)

	contact.decodeEmails(to)
	contact.decodePhones(to)
	contact.decodeLinks(to)
	contact.decodeNotes(to)

	return
}
