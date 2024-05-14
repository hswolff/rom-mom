package lib

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

func CalculateLocalDeltas(console string, romDir string) {
	romFiles, err := os.ReadDir(romDir)
	if err != nil {
		log.Fatal(err)
	}

	filesExtensions := path.Ext(romFiles[0].Name())
	romFileCache := []RomFile{}
	possibleMismatches := []RomFile{}
	missingRemotes := []RomFile{}

	remoteRomFiles, allRemoteRomNames := getRemoteRomFiles(console)

	for _, file := range romFiles {
		fileName := file.Name()
		fileNameWithoutExtension := strings.TrimSuffix(fileName, filesExtensions)
		results := fuzzy.RankFindFold(fileNameWithoutExtension, allRemoteRomNames)
		sort.Sort(results)
		// fmt.Printf("Searching for \"%s\"\n", fileNameWithoutExtension)

		romFile := RomFile{
			LocalName: fileName,
		}
		if results.Len() > 0 {
			remoteRomFile := remoteRomFiles[results[0].Target]
			romFile.RemoteRom = remoteRomFile
			romFile.AllMatches = results

			romFileCache = append(romFileCache, romFile)

			if results[0].Distance > 20 {
				possibleMismatches = append(possibleMismatches, romFile)
			}
		} else {
			missingRemotes = append(missingRemotes, romFile)
		}
	}

	PrettyPrint(missingRemotes)
}
