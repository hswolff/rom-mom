package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"sort"
	"strings"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

func main() {
	// show file and line number
	log.SetFlags(log.Llongfile)

	remoteRomFiles, allRemoteRomNames := getRemoteRomFiles("snes")

	// Now go through my local files and look for a match on a remote file
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

	romFiles := make([]RomFile, len(files))
	var missingRemotes []string

	for index, file := range files {
		fileName := file.Name()
		fileNameWithoutExtension := strings.TrimSuffix(fileName, filesExtensions)
		results := fuzzy.RankFindFold(fileNameWithoutExtension, allRemoteRomNames)
		sort.Sort(results)
		// fmt.Printf("Searching for \"%s\"\n", fileNameWithoutExtension)
		// fmt.Println(results)

		if results.Len() > 0 {
			remoteRomFile := remoteRomFiles[results[0].Target]
			romFiles[index] = RomFile{
				localName: fileName,
				remoteRom: remoteRomFile,
			}
		} else {
			missingRemotes = append(missingRemotes, fileName)
		}
	}

	PrettyPrint(missingRemotes)
}

type RemoteRomFile struct {
	remoteName   string
	remoteBoxArt string
}

type RomFile struct {
	localName string
	remoteRom RemoteRomFile
}
