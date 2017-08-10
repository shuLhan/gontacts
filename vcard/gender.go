package vcard

//
// Gender contains contact's sex and description.
//
// Sex may contain one of this value: M (male), F (female), O (other), N (none),
// or U (unknown).
//
type Gender struct {
	Sex  rune
	Desc string
}
