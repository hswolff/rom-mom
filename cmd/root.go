package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	console string
	romDir  string

	rootCmd = &cobra.Command{
		Use:   "rom-art-scraper",
		Short: "Fix your ROM collection.",
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			dirArg, err := cmd.Flags().GetString("dir")
			if err != nil {
				log.Fatal(err)
			}

			if c, err := cmd.Flags().GetString("console"); err != nil {
				console = c
			} else if len(args) > 0 {
				console = args[0]
			}

			if console == "" {
				return fmt.Errorf("console value required. Pass in via --console={console}")
			}

			if strings.HasPrefix(dirArg, "~") {
				homeDir, err := os.UserHomeDir()
				if err != nil {
					log.Fatal(err)
				}
				romDir = filepath.Join(homeDir, dirArg[1:])
			}

			if _, err := os.Stat(romDir); os.IsNotExist(err) {
				log.Fatalf("Directory does not exist: %s", romDir)
			}

			fmt.Println("Using local ROM directory:", romDir)
			fmt.Println("With selected console:", console)

			return nil
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&console, "console", "c", "", "console we're renaming")
	rootCmd.PersistentFlags().StringVarP(&romDir, "dir", "d", "", "ROM folder location")
}
