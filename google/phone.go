package google

// Phone format.
type Phone struct {
	Rel    string `json:"rel,omitempty"`
	Number string `json:"$t,omitempty"`
}
