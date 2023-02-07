package cmd

import (
	"testing"

	"github.com/Deepanshu276/arithmatic_cli/cmd"
)

// for integer value
func TestAdder(t *testing.T) {
	input := cmd.AddInt([]string{"1", "2"})
	expect := 4
	if input != expect {
		t.Errorf("got %q, wanted %q", input, expect)
	}

}
