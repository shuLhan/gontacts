package vcard

//
// Resource define common resource located in URI or embeded in Data.
//
type Resource struct {
	Type string
	URI  string
	Data []byte
}
