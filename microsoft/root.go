package microsoft

// Root of response.
type Root struct {
	Context  string    `json:"@odata.context"`
	Contacts []Contact `json:"value"`
}
