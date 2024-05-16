package main

import (
	"log"

	"github.com/hswolff/rom-mom/cmd"
)

func main() {
	log.SetFlags(log.Llongfile)
	cmd.Execute()
}
