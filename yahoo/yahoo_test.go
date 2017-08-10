package yahoo

import (
	"reflect"
	"runtime/debug"
	"testing"
)

func assert(t *testing.T, exp, got interface{}, equal bool) {
	if reflect.DeepEqual(exp, got) == equal {
		return
	}

	debug.PrintStack()

	t.Fatalf("\n"+
		">>> Expecting '%v'\n"+
		"          got '%v'\n", exp, got)
}
