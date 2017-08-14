package microsoft

//
// Email format on response.
//
type Email struct {
	Name    string `json:"name,omitempty"`
	Address string `json:"address,omitempty"`
}
