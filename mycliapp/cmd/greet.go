package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var name string

var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "Print a greeting message",
	Long:  `The greet command prints a msg to the console.`,
	Run: func(cmd *cobra.Command, args []string) {
		if verbose {
			fmt.Println("Verbose output enabled.")
		}

		if name != "" {
			fmt.Printf("Hello, %s! welcome to MYCli!\n", name)
		} else {
			fmt.Println("Hello, welcome to MYCli!")
		}
	},
}

func init() {
	greetCmd.Flags().StringVarP(&name, "name", "n", "", "Name of the person to greet")
	rootCmd.AddCommand(greetCmd)
}
