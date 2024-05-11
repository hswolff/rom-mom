package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

func getBoxArtHtml() *os.File {
	const htmlFile = "boxart.html"

	f, err := os.Open(htmlFile)
	if err == nil {
		fmt.Println("Using cached box art html file")
		return f
	}

	const boxArtUrl = "https://thumbnails.libretro.com/Nintendo%20-%20Super%20Nintendo%20Entertainment%20System/Named_Boxarts/"
	fmt.Println("Downloading box art from: ", boxArtUrl)
	res, err := http.Get(boxArtUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	out, err := os.Create(htmlFile)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	_, err = io.Copy(out, res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return out
}

func main() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	romDir := path.Join(homeDir, "Downloads", "games for anbernic", "SNES")

	files, err := os.ReadDir(romDir)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("In directory: ", romDir)

	filesExtensions := path.Ext(files[0].Name())
	fmt.Println("File extension: ", filesExtensions)

	if false {
		for _, file := range files {
			fileName := file.Name()
			fileNameWithoutExtension := strings.TrimSuffix(fileName, filesExtensions)
			fmt.Println(fileNameWithoutExtension)
		}
	}

	f := getBoxArtHtml()
	defer f.Close()
	fmt.Println("filename: ", f.Name())
}
