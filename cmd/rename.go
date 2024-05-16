package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hswolff/rom-art-scraper/lib"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(renameCmd)
}

var (
	includeMismatches bool
)

var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename all ROM files to matched file names",
	Long:  `Rename all ROM files to matched file names. This doesn't include missing matches or low quality matches`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("\n~~rename~~\n")
		romFiles, _, _ := lib.CalculateLocalDeltas(console, romDir)

		fmt.Printf("Local ROM files found:\t\t\t %v\n", len(romFiles))
		fmt.Printf("Number of matching ROMs:\t\t %v\n", len(romFiles.Matches()))
		fmt.Println("Starting renaming matching files. Skipping files that don't need to be renamed.")

		if dryRun {
			fmt.Printf("\n~~DRY RUN - NOTHING BEING RENAMED~~\n")
		}

		for _, romFile := range romFiles {
			if !includeMismatches && romFile.PossibleMismatch() {
				continue
			}

			if romFile.HasMatch() {
				oldPath := filepath.Join(romDir, romFile.LocalName)
				newPath := filepath.Join(romDir, romFile.RemoteRom.RemoteName)

				action := "RENAMING"
				if oldPath == newPath {
					action = "SKIPING"
				}

				debugText := ""
				if romFile.PossibleMismatch() {
					debugText = " (POSSIBLE MISMATCH)"
				}

				fmt.Printf("\n%s  (%s)%s\n", romFile.LocalName, action, debugText)
				fmt.Printf("\tCurrent: \t%s\n", oldPath)
				fmt.Printf("\tNew: \t\t%s\n", newPath)

				if !dryRun {
					if err := os.Rename(oldPath, newPath); err != nil {
						fmt.Printf("\tRENAME: FAILED %v\n", err)
					} else if action != "SKIPING" {
						fmt.Printf("\tRENAME: SUCCESS\n")
					}
				}
			}
		}
	},
}

func init() {
	renameCmd.Flags().BoolVar(&includeMismatches, "include-mismatch", false, "also rename possible mismatch file names")
}
