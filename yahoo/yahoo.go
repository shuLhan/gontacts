//
// Package yahoo implement user's contacts import using Yahoo API.
//
// References,
//
// [1] https://developer.yahoo.com/social/rest_api_guide/contacts-resource.html
//
package yahoo

import (
	"io/ioutil"
	"net/http"

	"github.com/json-iterator/go"

	"github.com/shuLhan/gontacts"
)

const (
	// List of APIs
	apiContactsURL    = "https://social.yahooapis.com/v1/user/"
	apiContactsSuffix = "/contacts?format=json&count=max"
)

//
// ImportFromJSON will parse JSON input and return list of Contact on success.
//
// On fail it will return nil and error.
//
func ImportFromJSON(jsonb []byte) (contacts []*gontacts.Contact, err error) {
	root := &Root{}

	err = jsoniter.Unmarshal(jsonb, root)
	if err != nil {
		return
	}

	for _, ycontact := range root.Contacts.Contact {
		contact := ycontact.Decode()
		contacts = append(contacts, contact)
	}

	return
}

//
// ImportWithOAuth will send a request to user's contact API using OAuth
// authentication code, and return list of Contact.
//
// On fail, it will return nil Contacts with error.
//
func ImportWithOAuth(
	tokenType string,
	accessToken string,
	guid string,
) (
	contacts []*gontacts.Contact,
	err error,
) {
	client := &http.Client{}
	api := apiContactsURL + guid + apiContactsSuffix

	req, err := http.NewRequest("GET", api, nil)

	req.Header.Add("Authorization", tokenType+" "+accessToken)

	res, err := client.Do(req)
	if err != nil {
		return
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	err = res.Body.Close()
	if err != nil {
		return
	}

	contacts, err = ImportFromJSON(resBody)

	return
}
