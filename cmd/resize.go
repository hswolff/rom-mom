package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
	"github.com/hswolff/rom-mom/lib"
	"github.com/spf13/cobra"
)

var (
	width  int
	height int
)

func init() {
	rootCmd.AddCommand(resizeCmd)

	resizeCmd.Flags().IntVar(&width, "width", 640, "width to resize images to")
	resizeCmd.Flags().IntVar(&height, "height", 480, "height to resize images to")
}

// 640x480 image size
var resizeCmd = &cobra.Command{
	Use:   "resize",
	Short: "resize images to desired dimension",
	Long:  "resize images to desired dimension",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Printf("\n~~resize images~~\n")
		romFiles, _, _ := lib.CalculateLocalDeltas(console, romDir)

		fmt.Printf("Local ROM files found:\t\t\t %v\n", len(romFiles))
		fmt.Printf("Number of matching ROMs:\t\t %v\n", len(romFiles.Matches()))

		if dryRun {
			fmt.Printf("\n~~DRY RUN - NOTHING BEING RESIZED~~\n")
		}

		imageFolderPath := filepath.Join(romDir, "images", console)
		images, err := os.ReadDir(imageFolderPath)
		if err != nil {
			return err
		}

		for _, image := range images {
			imagePath := filepath.Join(imageFolderPath, image.Name())
			img, err := imgio.Open(imagePath)
			if err != nil {
				log.Fatal(err)
			}

			bounds := img.Bounds()
			currentWidth, currentheight := bounds.Dx(), bounds.Dy()

			alreadyResized := false
			if width == currentWidth && height == currentheight {
				alreadyResized = true
			}

			action := "RESIZING"
			if alreadyResized {
				action = "SKIPING, ALREADY RESIZED"
			}

			fmt.Printf("\n%s  (%s)\n", imagePath, action)

			if !dryRun {
				if alreadyResized {
					continue
				}

				resized := transform.Resize(img, width, height, transform.Linear)
				if err := imgio.Save(imagePath, resized, imgio.PNGEncoder()); err != nil {
					fmt.Printf("\tResized: Error \t%s\n", err)
				} else {
					fmt.Printf("\tResized: Success\n")
				}
			}

		}

		return nil
	},
}
