package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"bytes"

	"github.com/lithium/macro11-assembler/src/macro11"
)

func main() {
	outputPath := flag.String("output", "a.out", "output path")
	listingPath := flag.String("listing", "a.lst", "listing file path")
	useListing := flag.Bool("use-listing", true, "should a listing be generated")

	flag.Parse()

	input, err := collectInput(os.Args[1:])
	if err != nil {
		panic(err)
	}

	outputBytes, listingBytes, err := macro11.Assemble(input)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(*outputPath, outputBytes, 0644) 
	if err != nil {
		panic(err)
	}

	if (*useListing) {
		err = ioutil.WriteFile(*listingPath, listingBytes, 0644)
		if err != nil {
			panic(err)
		}
	}

	fmt.Printf("done!\n")
}

func collectInput(args []string) (io.Reader, error) {
	var allInput bytes.Buffer;

	if len(args) < 1 {
		return os.Stdin, nil
	}

	for _, arg := range args {
		argFile, err := os.Open(arg)
		if err != nil {
			return nil, err
		}
		allInput.ReadFrom(argFile)
	}

	return bytes.NewReader(allInput.Bytes()), nil
}