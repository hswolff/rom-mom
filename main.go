package main

import (
	"log"

	"github.com/hswolff/rom-art-scraper/cmd"
)

func main() {
	log.SetFlags(log.Llongfile)
	cmd.Execute()
}
