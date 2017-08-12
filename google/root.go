package google

//
// Root define the root of Google's contact in JSON.
//
type Root struct {
	Version  string `json:"version,omitempty"`
	Encoding string `json:"encoding,omitempty"`
	Feed     Feed   `json:"feed,omitempty"`
}
