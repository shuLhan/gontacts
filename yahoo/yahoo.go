//
// Package yahoo implement user's contacts import using Yahoo API.
//
package yahoo

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
	oauthKey        = "xoauth_yahoo_guid"
	oauthRequestURL = "https://api.login.yahoo.com/oauth2/request_auth"
	oauthConfirmURL = "https://api.login.yahoo.com/oauth2/get_token"

	// List of APIs
	apiUserURL       = "https://social.yahooapis.com/v1/user/"
	apiContactSuffix = "/contacts?format=json&count=max"

	// List of default values
	defOAuthRedirectURL = "oob"
)

//
// Client define a client for Yahoo API.
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
func NewClient(cid, csecret, redirectURL string) (yc *Client) {
	if redirectURL == "" {
		redirectURL = defOAuthRedirectURL
	}

	yc = &Client{
		oauth: &oauth2.Config{
			ClientID:     cid,
			ClientSecret: csecret,
			RedirectURL:  redirectURL,
			Endpoint: oauth2.Endpoint{
				AuthURL:  oauthRequestURL,
				TokenURL: oauthConfirmURL,
			},
		},
	}

	return
}

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
// ImportContactsWithGUID will send a request to user's contact API based on
// GUID, parse it, and convert and return it as pointer to Contacts object.
//
// The response will be in JSON format with entire list of contacts (see
// apiContactSuffix).
//
// On fail it will return nil and error.
//
// References,
//
// [1] https://developer.yahoo.com/social/rest_api_guide/contacts-resource.html
//
func (yc *Client) ImportContactsWithGUID(guid string) (
	contacts []*gontacts.Contact,
	err error,
) {
	api := apiUserURL + guid + apiContactSuffix

	res, err := yc.http.Get(api)
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

//
// ImportContactsWithOAuth will send a request to user's contact API using OAuth
// authentication code, and return list of Contact.
//
// On fail, it will return nil Contacts with error.
//
func (yc *Client) ImportContactsWithOAuth(
	code string,
) (
	contacts []*gontacts.Contact,
	err error,
) {
	token, err := yc.oauth.Exchange(context.Background(), code)
	if err != nil {
		return
	}

	yc.http = yc.oauth.Client(context.Background(), token)

	return yc.ImportContactsWithGUID(token.Extra(oauthKey).(string))
}
