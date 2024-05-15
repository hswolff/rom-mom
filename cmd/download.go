package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(downloadCmd)
}

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download the box art images",
	Long:  "Download the box art images",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("DOWNLOAD")
	},
}
