package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// evenCmd represents the even command
var multiplyCmd = &cobra.Command{
	Use:   "multiply",
	Short: "Multiply two values",
	Long:  `This is used to multiply two values `,
	Run: func(cmd *cobra.Command, args []string) {

		fstatus, _ := cmd.Flags().GetBool("mulFloat")

		if fstatus {
			multiplyFloat(args)
		} else {
			MultiplyInt(args)
		}
	},
}

func init() {
	rootCmd.AddCommand(multiplyCmd)
	//BoolP is like Bool, but accepts a shorthand letter that can be used after a single dash.
	multiplyCmd.Flags().BoolP("mulFloat", "m", false, "multiply Floating Numbers")
}

// Logic for multiplying integer value
func MultiplyInt(args []string) int {
	mul := 1

	for _, ival := range args {
		itemp, err := strconv.Atoi(ival)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		mul = mul * itemp
	}

	fmt.Printf("multiplication of integer numbers %s is %d\n", args, mul)
	return mul
}

// logic for multiplying float value
func multiplyFloat(args []string) {
	mul := 1.0
	for _, fval := range args {
		//ParseFloat converts the string s to a floating-point number with the precision specified by bitSize
		ftemp, err := strconv.ParseFloat(fval, 64)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		mul = mul * ftemp
	}

	fmt.Printf("Multiplication of floating numbers %s is %f\n", args, mul)
}
