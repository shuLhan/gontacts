package yahoo

//
// Root define the root of JSON response.
//
type Root struct {
	Contacts Contacts `json:"contacts"`
}
