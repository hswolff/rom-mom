package main

import (
	"log"
	"os"
	"path"
	"sort"
	"strings"

	"github.com/lithammer/fuzzysearch/fuzzy"
)

type RomFile struct {
	LocalName  string
	RemoteRom  RemoteRomFile
	AllMatches fuzzy.Ranks
}

func calculateLocalDeltas(console string, romDir string) {
	romFiles, err := os.ReadDir(romDir)
	if err != nil {
		log.Fatal(err)
	}

	filesExtensions := path.Ext(romFiles[0].Name())
	romFileCache := []RomFile{}
	var missingRemotes []string
	possibleMismatches := []RomFile{}

	remoteRomFiles, allRemoteRomNames := getRemoteRomFiles(console)

	for _, file := range romFiles {
		fileName := file.Name()
		fileNameWithoutExtension := strings.TrimSuffix(fileName, filesExtensions)
		results := fuzzy.RankFindFold(fileNameWithoutExtension, allRemoteRomNames)
		sort.Sort(results)
		// fmt.Printf("Searching for \"%s\"\n", fileNameWithoutExtension)

		if results.Len() > 0 {
			remoteRomFile := remoteRomFiles[results[0].Target]
			romFile := RomFile{
				LocalName:  fileName,
				RemoteRom:  remoteRomFile,
				AllMatches: results,
			}
			romFileCache = append(romFileCache, romFile)

			if results[0].Distance > 20 {
				possibleMismatches = append(possibleMismatches, romFile)
			}
		} else {
			missingRemotes = append(missingRemotes, fileName)
		}
	}

	PrettyPrint(possibleMismatches)
}
