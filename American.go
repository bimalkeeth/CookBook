package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	inFilename, outFilename, err := filesNamesFromCommandline()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	inFile, outFile := os.Stdin, os.Stdout
	if inFilename != "" {
		if inFile, err = os.Open(inFilename); err != nil {
			log.Fatal(err)
		}
		defer inFile.Close()
	}
	if inFilename != "" {
		if outFile, err = os.Open(outFilename); err != nil {
			log.Fatal(err)
		}
		defer outFile.Close()
	}
	if err = americanise(inFile, outFile); err != nil {
		log.Fatal(err)
	}

}

func americanise(file *os.File, file2 *os.File) error {

	return nil
}

func filesNamesFromCommandline() (inFileName string, outFileName string, err error) {
	if len(os.Args) > 1 && os.Args[1] == "-h" || os.Args[1] == "--help" {
		err = fmt.Errorf("usage %s [<] infile.txt[>] outfile.txt", filepath.Base(os.Args[0]))
		return "", "", err
	}
	if len(os.Args) > 1 {
		inFileName = os.Args[1]
		if len(os.Args) > 2 {
			outFileName = os.Args[2]
		}
	}
	if inFileName != "" && inFileName == outFileName {
		log.Fatal("won't overwrite infile")
	}
	return inFileName, outFileName, nil
}
