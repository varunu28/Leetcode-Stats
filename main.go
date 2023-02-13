package main

import (
	"flag"
	"leetcode-stats/cmd"
	"log"
)

var directoryPath = flag.String("directory", "", "path for leetcode directory")

func main() {
	flag.Parse()
	if *directoryPath == "" {
		log.Fatal("You need to provide a directory path for the script")
	}

	cmd.Compute(*directoryPath)
}
