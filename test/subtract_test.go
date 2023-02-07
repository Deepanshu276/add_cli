package cmd

import (
	"testing"

	"github.com/Deepanshu276/arithmatic_cli/cmd"
)

// for integer value
func TestSubtract(t *testing.T) {
	input := cmd.SubInt([]string{"1", "2"})
	expect := -1
	if input != expect {
		t.Errorf("got %q, wanted %q", input, expect)
	}

}
