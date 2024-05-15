package cmd

import (
	"fmt"

	"github.com/hswolff/rom-art-scraper/lib"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(checkCmd)
}

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check if matches are found and display any missing or mismatches",
	Long:  `Check if matches are found and display any missing or mismatches`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("check", console, romDir)
		lib.CalculateLocalDeltas(console, romDir)

	},
}
