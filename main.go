package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/lithium/macro11-assembler/src/macro11"
)

func main() {
	outputPath := flag.String("output", "a.out", "output path")
	listingPath := flag.String("listing", "a.lst", "listing file path")
	useListing := flag.Bool("use-listing", true, "should a listing be generated")

	flag.Parse()

	inputText, err := collectInput(os.Args[1:])
	if err != nil {
		panic(err)
	}

	outputBytes, listingBytes, err := macro11.Assemble(inputText)
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

func collectInput(args []string) ([]byte, error) {
	var allText []byte;

	if len(args) < 1 {
		return ioutil.ReadAll(os.Stdin)
	}

	for _, arg := range args {
		inputText, err := ioutil.ReadFile(arg)
		if err != nil {
			return nil, err
		}
		allText = append(allText, inputText...)
	}

	return allText, nil
}