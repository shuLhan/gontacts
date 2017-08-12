package google

// Email format.
type Email struct {
	Rel     string `json:"rel,omitempty"`
	Address string `json:"address,omitempty"`
	Primary string `json:"primary,omitempty"`
}
