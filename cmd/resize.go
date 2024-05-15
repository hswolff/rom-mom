package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(resizeCmd)
}

var resizeCmd = &cobra.Command{
	Use:   "resize",
	Short: "resize images to desired dimension",
	Long:  "resize images to desired dimension",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("DOWNLOAD")
	},
}
