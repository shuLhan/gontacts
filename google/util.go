package google

import (
	"strings"
)

//
// ParseRel will parse Google "rel" value and return the type.
//
func ParseRel(in string) string {
	kv := strings.Split(in, "#")
	if len(kv) != 2 {
		return ""
	}

	return kv[1]
}
