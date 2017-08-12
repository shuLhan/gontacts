package google

// Event format.
type Event struct {
	Rel  string    `json:"rel,omitempty"`
	When EventTime `json:"gd$when,omitempty"`
}
