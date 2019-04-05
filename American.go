package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
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

var britishAmerican = "british-american.txt"

func americanise(inFile io.Reader, outFile io.Writer) (err error) {

	reader := bufio.NewReader(inFile)
	writer := bufio.NewWriter(outFile)

	defer func() {
		if err == nil {
			err = writer.Flush()
		}
	}()
	var replacer func(string) string
	if replacer, err = makeReplaceFunction(britishAmerican); err != nil {
		return err
	}
	wordRx := regexp.MustCompile("[A-Za-z]+")
	eof := false
	for !eof {
		var line string
		line, err = reader.ReadString('\n')
		if err == io.EOF {
			err = nil
			eof = true
		} else {
			return err
		}
		line = wordRx.ReplaceAllStringFunc(line, replacer)
		if _, err = writer.WriteString(line); err != nil {
			return err
		}
	}

	return nil
}

func makeReplaceFunction(s string) (func(string) string, error) {

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
