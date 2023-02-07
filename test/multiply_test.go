package cmd

import (
	"testing"

	"github.com/Deepanshu276/arithmatic_cli/cmd"
)

// for integer value
func TestMultiply(t *testing.T) {
	input := cmd.MultiplyInt([]string{"1", "2"})
	expect := 2
	if input != expect {
		t.Errorf("got %q, wanted %q", input, expect)
	}

}
