//
// Package microsoft implement Microsoft Live's contact API.
//
// Reference,
// (1) https://developer.microsoft.com/en-us/graph/docs/api-reference/v1.0/api/user_list_contacts
//
package microsoft

import (
	"io/ioutil"
	"net/http"

	"github.com/json-iterator/go"

	"github.com/shuLhan/gontacts"
)

const (
	// List of provider APIs.
	apiContactsURL = "https://graph.microsoft.com/v1.0/me/contacts"
)

//
// ImportFromJSON will parse Microsoft Live's JSON contact response and return
// list of gontacts Contact on success.
//
func ImportFromJSON(jsonb []byte) (
	contacts []*gontacts.Contact,
	err error,
) {
	root := &Root{}

	err = jsoniter.Unmarshal(jsonb, root)
	if err != nil {
		return
	}

	for _, mscontact := range root.Contacts {
		contact := mscontact.Decode()
		contacts = append(contacts, contact)
	}

	return
}

//
// ImportWithOAuth will send a request to user's contact API using OAuth
// authentication code, and return pointer to Contacts object.
//
// On fail, it will return nil Contacts with error.
//
func ImportWithOAuth(
	tokenType string,
	accessToken string,
) (
	contacts []*gontacts.Contact,
	err error,
) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", apiContactsURL, nil)

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
