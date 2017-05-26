package main

import (
	"flag"

	"log"
	"os"

	"gitlab.com/jannickfahlbusch/hideFile"
)

var (
	filePath string
	out      string
	fileType string
	decode   bool
)

func init() {
	flag.StringVar(&filePath, "filePath", "", "Path to the file")
	flag.StringVar(&out, "out", "", "Output directory")
	flag.StringVar(&fileType, "type", "JPEG", "The type to which the file should be converted")
	flag.BoolVar(&decode, "decode", false, "Decode the file")
}

func main() {
	flag.Parse()

	hider := hideFile.NewHider()

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	toType, err := hider.GetType(fileType)
	if err != nil {
		log.Fatal(err)
	}

	if decode {
		err = hider.Deconvert(file, toType, out)
		if err != nil {
			log.Fatal(err)
		}

	} else {
		err = hider.Convert(file, toType, out)
		if err != nil {
			log.Fatal(err)
		}

	}
}
