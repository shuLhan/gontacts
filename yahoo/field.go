package yahoo

import (
	"strings"
)

//
// Field define a composite attribute in Contact.
// Known value for Type: "phone", "name", "address".
//
type Field struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
	Flags []string    `json:"flags"`

	// Ignored fields for speedup.
	//ID         int         `json:"id"`
	//EditedBy   string      `json:"editedBy"`
	//Categories []string    `json:"categories"`
	//Meta
}

//
// getValueType will return the Go type of field's Value.
//
func (field *Field) getValueType() (
	vmap map[string]interface{},
	vstr string,
	ok bool,
) {
	ok = true

	switch v := field.Value.(type) {
	case map[string]interface{}:
		vmap = v
	case string:
		vstr = v
	default:
		ok = false
	}

	return
}

//
// getFlag will return the first flag or empty string if flags fields is empty.
//
func (field *Field) getFlag() string {
	if len(field.Flags) > 0 {
		return strings.ToLower(field.Flags[0])
	}
	return ""
}
