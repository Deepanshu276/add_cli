package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// oddCmd represents the odd command
var subCmd = &cobra.Command{
	Use:   "subtract",
	Short: "Subtraction of numbers",
	Long:  `Used in only subtraction of numbers`,
	Run: func(cmd *cobra.Command, args []string) {
		//logic to add two even numbers
		// returning boolean value if flag float is used
		fstatus, _ := cmd.Flags().GetBool("subfloat")

		if fstatus {
			subFloat(args)
		} else {
			SubInt(args)
		}
	},
}

func init() {
	rootCmd.AddCommand(subCmd)

	//BoolP is like Bool, but accepts a shorthand letter that can be used after a single dash.
	subCmd.Flags().BoolP("subfloat", "s", false, "Add Floating Numbers")
}

// Logic for adding integer value
func SubInt(args []string) int {
	var diff int
	for i, ival := range args {
		itemp, err := strconv.Atoi(ival)
		if i == 0 {
			diff = itemp
		} else {
			diff = diff - itemp
		}
		if err != nil {
			fmt.Println(err)
		}

	}

	fmt.Printf("Addition of numbers %s is %d\n", args, diff)
	return diff

}

// logic for adding float value
func subFloat(args []string) {
	var diff float64
	for i, fval := range args {
		//ParseFloat converts the string s to a floating-point number with the precision specified by bitSize
		ftemp, err := strconv.ParseFloat(fval, 64)
		if i == 0 {
			diff = ftemp
		} else {
			diff = diff - ftemp
		}
		if err != nil {
			fmt.Println(err)
		}

	}

	fmt.Printf("Sum of floating numbers %s is %f\n", args, diff)
}
