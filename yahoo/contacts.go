package yahoo

//
// Contacts define the holder for root of contacts response.
//
type Contacts struct {
	Contact []Contact `json:"contact"`
	Start   int       `json:"start"`
	Count   int       `json:"count"`
	Total   int       `json:"total"`
	URI     string    `json:"uri"`
}
