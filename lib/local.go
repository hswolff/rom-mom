package lib

import (
	"fmt"
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

func CalculateLocalDeltas(console string, romDir string) (romFiles, possibleMismatches, missingRemotes []RomFile) {
	dirFiles, err := os.ReadDir(romDir)
	if err != nil {
		log.Fatal(err)
	}

	filesExtensions := path.Ext(dirFiles[0].Name())

	// romFileCache := []RomFile{}
	// possibleMismatches := []RomFile{}
	// missingRemotes := []RomFile{}

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

		if results.Len() > 0 {
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

	mytest := func(r RomFile) bool { return !r.HasMatch() }
	s2 := filter(romFiles, mytest)

	fmt.Println(len(romFiles), len(missingRemotes), len(s2))
	// PrettyPrint(s2)
	return
}
