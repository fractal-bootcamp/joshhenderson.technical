package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(hello)
}

var hello = &cobra.Command{
	Use:   "hello",
	Short: "prints hello",
	Long:  `prints hello`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("hello")

	},
}
