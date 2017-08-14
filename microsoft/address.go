package microsoft

//
// Address format on response.
//
type Address struct {
	Street     string `json:"street,omitempty"`
	City       string `json:"city,omitempty"`
	State      string `json:"state,omitempty"`
	Country    string `json:"countryOrRegion,omitempty"`
	PostalCode string `json:"postalCode,omitempty"`
}
