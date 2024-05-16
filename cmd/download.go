package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/hswolff/rom-art-scraper/lib"
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
		fmt.Printf("\n~~download images~~\n")
		romFiles, _, _ := lib.CalculateLocalDeltas(console, romDir)

		fmt.Printf("Local ROM files found:\t\t\t %v\n", len(romFiles))
		fmt.Printf("Number of matching ROMs:\t\t %v\n", len(romFiles.Matches()))
		fmt.Println("Downloading matched ROM images. Skipping images already downloaded.")

		if dryRun {
			fmt.Printf("\n~~DRY RUN - NOTHING BEING DOWNLOADED~~\n")
		}

		imageFolderPath := filepath.Join(romDir, "images", console)

		// Create image folder if it doesn't exist
		// ensure folder exists
		dirErr := os.MkdirAll(imageFolderPath, os.ModePerm)
		if dirErr != nil {
			log.Fatal(dirErr)
		}

		for _, romFile := range romFiles {
			if !includeMismatches && romFile.PossibleMismatch() {
				continue
			}

			if romFile.HasMatch() {
				remoteUrl := romFile.RemoteRom.RemoteBoxArt
				downloadPath := filepath.Join(imageFolderPath, romFile.RemoteRom.RemoteImageName)

				fileAlreadyDownloaded := false
				if _, err := os.Stat(downloadPath); err == nil {
					fileAlreadyDownloaded = true
				}

				action := "DOWNLOADING"
				if fileAlreadyDownloaded {
					action = "SKIPING, ALREADY DOWNLOADED"
				}

				debugText := ""
				if romFile.PossibleMismatch() {
					debugText = " (POSSIBLE MISMATCH)"
				}

				fmt.Printf("\n%s  (%s)%s\n", romFile.LocalName, action, debugText)
				fmt.Printf("\tRemote URL: \t%s\n", remoteUrl)
				fmt.Printf("\tSaving to: \t%s\n", downloadPath)

				if !dryRun {
					if fileAlreadyDownloaded {
						continue
					}

					resp, err := http.Get(remoteUrl)
					if err != nil {
						fmt.Printf("\tDownload: FAILED (%v)\n", err)
						continue
					}
					defer resp.Body.Close()

					out, err := os.Create(downloadPath)
					if err != nil {
						log.Fatal(err)
					}
					defer out.Close()

					_, err = io.Copy(out, resp.Body)
					if err != nil {
						fmt.Printf("\tDownload: CANT SAVE TO DISK (%v)\n", err)
						continue
					}

					fmt.Printf("\tDownload: SUCCESS\n")
				}
			}
		}
	},
}
