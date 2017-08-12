//
// Package google implement Google's contact API v3 for import.
//
package google

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/json-iterator/go"
	"golang.org/x/oauth2"

	"github.com/shuLhan/gontacts"
)

const (
	// List of OAuth2 APIs and keys
	oauthRequestURL = "https://accounts.google.com/o/oauth2/v2/auth"
	oauthConfirmURL = "https://www.googleapis.com/oauth2/v4/token"

	// List of APIs
	apiUserURL = "https://www.google.com/m8/feeds/contacts/default/full?alt=json&max-results=50000&v=3.0"
)

//
// Client define a client for Google API.
//
type Client struct {
	oauth *oauth2.Config
	http  *http.Client
}

//
// NewClient will initialize OAuth2 with client ID `cid`, client secret
// `csecret`, and redirect URL (optional).
//
// If redirectURL is empty it will be set to `oob`.
//
func NewClient(cid, csecret, redirectURL string) (client *Client) {
	client = &Client{
		oauth: &oauth2.Config{
			ClientID:     cid,
			ClientSecret: csecret,
			RedirectURL:  redirectURL,
			Scopes: []string{
				"https://www.googleapis.com/auth/contacts.readonly",
			},
			Endpoint: oauth2.Endpoint{
				AuthURL:  oauthRequestURL,
				TokenURL: oauthConfirmURL,
			},
		},
	}

	return
}

//
// NewContacts will parse JSON input and return Contacts object on success.
//
// On fail it will return nil and error.
//
func NewContacts(jsonb []byte) (contacts []*gontacts.Contact, err error) {
	root := &Root{}

	err = jsoniter.Unmarshal(jsonb, root)
	if err != nil {
		return
	}

	for _, gcontact := range root.Feed.Contacts {
		contact := gcontact.Decode()
		contacts = append(contacts, contact)
	}

	return
}

func (client *Client) fetchContacts() (
	contacts []*gontacts.Contact,
	err error,
) {
	res, err := client.http.Get(apiUserURL)
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

	contacts, err = NewContacts(resBody)

	return
}

//
// ImportContactsWithOAuth will send a request to user's contact API using OAuth
// authentication code, and return pointer to Contacts object.
//
// On fail, it will return nil Contacts with error.
//
func (client *Client) ImportContactsWithOAuth(
	code string,
) (
	contacts []*gontacts.Contact,
	err error,
) {
	token, err := client.oauth.Exchange(context.Background(), code)
	if err != nil {
		return
	}

	client.http = client.oauth.Client(context.Background(), token)

	return client.fetchContacts()
}
