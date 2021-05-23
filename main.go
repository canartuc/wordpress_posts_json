package main

import (
	"flag"
	"log"
)

func main() {
	// This will read command line parameter "out" as output file and writes all data in it
	var outFile string
	flag.StringVar(&outFile, "out", "", "Enter relative path of output JSON file")
	flag.Parse()

	if outFile == "" {
		log.Fatal("(ERR) No config no output. Please define output file's relative path including filename!")
	}

	Write2Json(outFile)
}
