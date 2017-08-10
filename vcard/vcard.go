//
// Package vcard implement RFC6350 for encoding and decoding VCard formatted
// data.
//
package vcard

import (
	"github.com/shuLhan/gontacts"
)

//
// VCard define vcard 4.0 data structure.
//
type VCard struct {
	UID          string
	Source       []string
	Kind         string
	Fn           string
	N            gontacts.Name
	Nickname     []string
	Photo        []Resource
	Bday         gontacts.Date
	Anniversary  gontacts.Date
	Gender       Gender
	Adr          []gontacts.Address
	Tel          []gontacts.Phone
	Email        []gontacts.Email
	Impp         []Messaging
	Lang         []string
	TZ           string
	Geo          []GeoLocation
	Title        []string
	Role         []string
	Logo         []Resource
	Org          []string
	Related      []Relation
	Categories   []string
	Note         []string
	ProdID       string
	Sound        []Resource
	ClientPIDMap string
	Key          []Resource
}
