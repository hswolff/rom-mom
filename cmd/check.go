package cmd

import (
	"fmt"
	"strings"

	"github.com/hswolff/rom-mom/lib"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(checkCmd)
}

var (
	list        string
	listOptions []string = []string{"match", "mismatch", "missing"}
	verbose     bool
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check if matches are found and display any missing or mismatches",
	Long:  `Check if matches are found and display any missing or mismatches`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		// Validate list flag
		if list != "" && !lib.StringInSlice(list, listOptions) {
			return fmt.Errorf("invalid list option: %s. It should be one of %v", list, strings.Join(listOptions, ", "))
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("\n~~check~~\n")
		romFiles, possibleMismatches, missingRemotes := lib.CalculateLocalDeltas(console, romDir)

		fmt.Printf("Local ROM files found:\t\t\t %v\n", len(romFiles))
		fmt.Printf("Number of matching ROMs:\t\t %v\n", len(romFiles.Matches()))
		fmt.Printf("Number of ROMs not matched:\t\t %v\n", len(missingRemotes))
		fmt.Printf("Number of possible mismatched ROMs:\t %v\n", len(possibleMismatches))

		if list != "" {
			fmt.Printf("Listing %s\n", list)

			switch list {
			case "match":
				fmt.Println("\t\tLOCAL NAME\t\t\t\t\tREMOTE NAME")
				for _, romFile := range romFiles {
					if romFile.HasMatch() {
						fmt.Printf("MATCHED   %s  ->  %s\n", romFile.LocalName, romFile.RemoteRom.RemoteName)
					}
				}
			case "mismatch":
				if !verbose {
					fmt.Println("\t\tLOCAL NAME\t\t\t\t\tREMOTE NAME")
				}

				for _, romFile := range romFiles {
					if romFile.PossibleMismatch() {
						if verbose {
							fmt.Printf("\nMISMATCH   %s\n", romFile.LocalName)
							fmt.Println("    All Matches")
							for _, match := range romFile.AllMatches {
								fmt.Printf("\t%s\t(distance: %d)\n", match.Target, match.Distance)
							}
						} else {
							fmt.Printf("MISMATCH   %s  ->  %s\n", romFile.LocalName, romFile.RemoteRom.RemoteName)
						}
					}
				}
			case "missing":
				fmt.Println("\t\tLOCAL NAME\t\t\t\t\tREMOTE NAME")
				for _, romFile := range romFiles {
					if !romFile.HasMatch() {
						fmt.Printf("MISSING   %s\n", romFile.LocalName)
					}
				}
			}
		}

	},
}

func init() {
	checkCmd.Flags().StringVarP(&list, "list", "l", "", fmt.Sprintf("list the roms. one of %v", strings.Join(listOptions, ", ")))
	checkCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "show more information")
}
