package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var currentVersion = "0.0.1"

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Program version",
	Long:  `Shows the current program version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("version %v\n", currentVersion)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

}
