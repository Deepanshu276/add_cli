package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// divideCmd represents the divide command
var divideCmd = &cobra.Command{
	Use:   "divide",
	Short: "divide the numbers ",
	Long:  `This is used to divide numbers. Numbers can be both float type number and integer type number `,
	Run: func(cmd *cobra.Command, args []string) {

		fstatus, _ := cmd.Flags().GetBool("divFloat")

		if fstatus {
			divideFloat(args)
		} else {
			DivideInt(args)
		}
	},
}

func init() {
	rootCmd.AddCommand(divideCmd)
	//BoolP is like Bool, but accepts a shorthand letter that can be used after a single dash.
	divideCmd.Flags().BoolP("divFloat", "d", false, "multiply Floating Numbers")
}

// Logic for dividing integer value
func DivideInt(args []string) int {
	var div int
	for i, ival := range args {
		itemp, err := strconv.Atoi(ival)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if i == 0 {
			div = int(itemp)
		} else {
			div /= int(itemp)
		}

	}

	fmt.Printf("division of integer numbers %s is  %d\n", args, div)
	return div

}

// logic for dividing float value
func divideFloat(args []string) float64 {
	var div float64
	for i, fval := range args {
		//ParseFloat converts the string s to a floating-point number with the precision specified by bitSize
		ftemp, err := strconv.ParseFloat(fval, 64)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if i == 0 {
			div = float64(ftemp)
		} else {
			div /= float64(ftemp)
		}

		fmt.Printf("division of floating numbers %s is %f\n", args, div)
	}
	return div
}
