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

func (r *RomFile) HasMatch() bool {
	return r.AllMatches.Len() > 0
}

func (r *RomFile) PossibleMismatch() bool {
	distance := 20

	return r.HasMatch() && r.AllMatches[0].Distance > distance
}

type RomFiles []RomFile

func (r *RomFiles) Matches() RomFiles {
	result := filter(*r, func(r RomFile) bool {
		return r.HasMatch()
	})
	return result
}

func CalculateLocalDeltas(console string, romDir string) (romFiles, possibleMismatches, missingRemotes RomFiles) {
	dirFiles, err := os.ReadDir(romDir)
	if err != nil {
		log.Fatal(err)
	}

	filesExtensions := path.Ext(dirFiles[0].Name())

	remoteRomFiles, allRemoteRomNames := getRemoteRomFiles(console)

	for _, file := range dirFiles {
		fileName := file.Name()
		fileNameWithoutExtension := strings.TrimSuffix(fileName, filesExtensions)
		results := fuzzy.RankFindFold(fileNameWithoutExtension, allRemoteRomNames)
		sort.Sort(results)

		romFile := RomFile{
			LocalName:  fileName,
			AllMatches: results,
		}

		if romFile.HasMatch() {
			remoteRomFile := remoteRomFiles[results[0].Target]
			romFile.RemoteRom = remoteRomFile

			if romFile.PossibleMismatch() {
				possibleMismatches = append(possibleMismatches, romFile)
			}
		} else {
			missingRemotes = append(missingRemotes, romFile)
		}

		romFiles = append(romFiles, romFile)
	}

	return
}
