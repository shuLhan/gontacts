//
// Package microsoft implement Microsoft Live's contact API.
//
// Reference,
// (1) https://developer.microsoft.com/en-us/graph/docs/api-reference/v1.0/api/user_list_contacts
//
package microsoft

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
	oauthRequestURL = "https://login.microsoftonline.com/common/oauth2/v2.0/authorize"
	oauthConfirmURL = "https://login.microsoftonline.com/common/oauth2/v2.0/token"

	// List of provider APIs.
	apiContactsURL = "https://graph.microsoft.com/v1.0/me/contacts"
)

//
// Client define a client for Microsoft OAuth API.
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
func NewClient(cid, csecret, redirectURL string) (c *Client) {
	c = &Client{
		oauth: &oauth2.Config{
			ClientID:     cid,
			ClientSecret: csecret,
			RedirectURL:  redirectURL,
			Scopes: []string{
				"Contacts.Read",
			},
			Endpoint: oauth2.Endpoint{
				AuthURL:  oauthRequestURL,
				TokenURL: oauthConfirmURL,
			},
		},
		http: &http.Client{},
	}

	return
}

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
// Fetch will get user contact from provider API with JSON format,
// parse it and return list of contacts on success.
//
func (client *Client) Fetch(url string) (
	contacts []*gontacts.Contact,
	err error,
) {
	res, err := client.http.Get(url)
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

	err = ioutil.WriteFile("/tmp/live_contacts.json", resBody, 0600)
	if err != nil {
		return
	}

	contacts, err = ImportFromJSON(resBody)

	return
}

//
// ImportWithOAuth will send a request to user's contact API using OAuth
// authentication code, and return pointer to Contacts object.
//
// On fail, it will return nil Contacts with error.
//
func (client *Client) ImportWithOAuth(
	code string,
) (
	token *oauth2.Token,
	contacts []*gontacts.Contact,
	err error,
) {
	token, err = client.oauth.Exchange(context.Background(), code)
	if err != nil {
		return
	}

	client.http = client.oauth.Client(context.Background(), token)

	contacts, err = client.Fetch(apiContactsURL)

	return
}
