package main

import (
	"fmt"
	"log"
	"os"
	"path"
)

func main() {
	// show file and line number
	log.SetFlags(log.Llongfile)

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	romDir := path.Join(homeDir, "Downloads", "games for anbernic", "SNES")
	fmt.Println("In directory: ", romDir)

	calculateLocalDeltas("snes", romDir)
}
