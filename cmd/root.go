package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/hswolff/rom-art-scraper/lib"
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	console   string
	romFolder string

	rootCmd = &cobra.Command{
		Use:   "rom-art-scraper",
		Short: "Fix your ROM collection.",
		Args: func(cmd *cobra.Command, args []string) error {
			n := 1
			if len(args) != n {
				return fmt.Errorf("accepts %d arg(s), received %d", n, len(args))
			}

			console = args[0]
			if _, exist := lib.ConsolesAvailable[console]; !exist {
				return fmt.Errorf("selected console \"%s\" not supported", console)
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			romDir, _ := cmd.Flags().GetString("dir")
			console = args[0]
			// fmt.Println("hello run", romDir, args)

			if strings.HasPrefix(romDir, "~") {
				homeDir, err := os.UserHomeDir()
				if err != nil {
					log.Fatal(err)
				}
				romDir = filepath.Join(homeDir, romDir[1:])
			}

			if _, err := os.Stat(romDir); os.IsNotExist(err) {
				log.Fatalf("Directory does not exist: %s", romDir)
			}

			fmt.Println("Using local ROM directory:", romDir)

			// lib.CalculateLocalDeltas(console, romDir)
		},
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&console, "console", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&romFolder, "dir", "d", "", "ROM folder location")
	// rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	// viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	// viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	// viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	// viper.SetDefault("license", "apache")

	// rootCmd.AddCommand(addCmd)
	// rootCmd.AddCommand(initCmd)
}

/*
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
*/
