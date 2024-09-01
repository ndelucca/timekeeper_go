/*
Copyright © 2024 ndelucca ndelucca@protonmail.com
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "timekeeper",
	Short: "Timekeeping cli tool.",
	Long:  `Registers time ranges and keeps track of them day-by-day.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
