package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"

	"github.com/PuerkitoBio/goquery"
)

const rootPath = "https://thumbnails.libretro.com"
const boxArtPath = "Named_Boxarts"

var consoleToPath = map[string]string{
	"snes": "Nintendo%20-%20Super%20Nintendo%20Entertainment%20System",
}

func createBoxArtUrl(console string) string {
	res, err := url.JoinPath(rootPath, consoleToPath[console], boxArtPath)
	if err != nil {
		log.Fatal(err)
	}
	return res
}

var boxArtUrl = createBoxArtUrl("snes")

func getBoxArtHtml(console string) *os.File {
	folderName := "boxart"
	htmlFile := path.Join(folderName, fmt.Sprintf("%s.html", console))

	f, err := os.Open(htmlFile)
	if err == nil {
		fmt.Println("Using cached box art html file", htmlFile)
		return f
	}

	// ensure folder exists
	dirErr := os.MkdirAll(folderName, os.ModePerm)
	if dirErr != nil {
		log.Fatal(dirErr)
	}

	out, err := os.Create(htmlFile)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Downloading box art from: ", boxArtUrl)
	res, err := http.Get(boxArtUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	_, err = io.Copy(out, res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return out
}

func getRemoteRomFiles(console string) (map[string]RemoteRomFile, []string) {
	f := getBoxArtHtml(console)
	defer f.Close()

	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		log.Fatal(err)
	}

	remoteRomFiles := make(map[string]RemoteRomFile)

	allRemoteRomNames := doc.Find("a").Map(func(_ int, s *goquery.Selection) string {
		// For each item found, get the title
		anchor, _ := s.Attr("href")
		remoteName, _ := url.QueryUnescape(anchor)
		// fmt.Printf("ROM: %s | %s\n", title, remoteName)

		remoteBoxArt, _ := url.JoinPath(boxArtUrl, anchor)

		remoteRomFiles[remoteName] = RemoteRomFile{
			remoteName:   remoteName,
			remoteBoxArt: remoteBoxArt,
		}

		return remoteName
	})

	return remoteRomFiles, allRemoteRomNames
}
