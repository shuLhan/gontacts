package yahoo

import (
	"time"
)

//
// Meta define a common metadata inside a struct.
//
type Meta struct {
	Created time.Time `json:"created"`
	Updated time.Time `json:"updated"`
	URI     string    `json:"uri"`
}
