package gontacts

//
// Name define contact's name.
//
type Name struct {
	Given       string `json:"givenName"`
	Middle      string `json:"middleName"`
	Family      string `json:"familyName"`
	Prefix      string `json:"prefix"`
	Suffix      string `json:"suffix"`
	GivenSound  string `json:"givenNameSound"`
	FamilySound string `json:"familyNameSound"`
}
