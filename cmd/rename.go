package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(renameCmd)
}

var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename all ROM files to matched file names",
	Long:  `Rename all ROM files to matched file names. This doesn't include missing matches or low quality matches`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("RENAME")
	},
}
