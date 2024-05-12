package main

import (
	"log"
)

func main() {
	// show file and line number
	log.SetFlags(log.Llongfile)

	remoteRomFiles, allRemoteRomNames := getRemoteRomFiles("snes")
	calculateLocalDeltas(remoteRomFiles, allRemoteRomNames)
}
