package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

// evenCmd represents the even command
var evenCmd = &cobra.Command{
	Use:   "even",
	Short: "Adding only even value",
	Long:  `This is used in only addition of even numbers `,
	Run: func(cmd *cobra.Command, args []string) {
		//logic to find the sum of only even numbers
		var evenSum int
		for _, ival := range args {
			itemp, _ := strconv.Atoi(ival)
			if itemp%2 == 0 {
				evenSum = evenSum + itemp
			} else {
				fmt.Printf("number %d is not even number, please enter even number\n", itemp)
				os.Exit(1)
			}
		}

		fmt.Printf("The even addition of %s is %d\n", args, evenSum)
	},
}

func init() {
	addCmd.AddCommand(evenCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// evenCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// evenCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
