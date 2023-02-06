package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// oddCmd represents the odd command
var oddCmd = &cobra.Command{
	Use:   "odd",
	Short: "Addition of odd numbers",
	Long:  `Used in only addition of odd numbers`,
	Run: func(cmd *cobra.Command, args []string) {
		//logic to add two even numbers
		var oddSum int
		for _, ival := range args {
			itemp, _ := strconv.Atoi(ival)
			//if number is not odd terminate the program with the error message
			if itemp%2 == 0 {
				fmt.Printf("number %d is not odd number, please enter odd number\n", itemp)
				os.Exit(1)

			} else {
				oddSum = oddSum + itemp

			}

		}
		fmt.Printf("The odd addition of %s is %d\n", args, oddSum)

	},
}

func init() {
	addCmd.AddCommand(oddCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// oddCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// oddCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
