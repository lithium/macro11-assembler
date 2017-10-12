package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	inputPath := flag.String("input", "", "path to input source file")
	outputPath := flag.String("output", "a.out", "output path")

	flag.Parse()

	var inputText []byte;
	var err error;

	if len(*inputPath) > 0 {
		inputText, err = ioutil.ReadFile(*inputPath)
	} else {
		inputText, err = ioutil.ReadAll(os.Stdin)
	}
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(*outputPath, inputText, 0644)
	if err != nil {
		panic(err)
	}

	fmt.Printf("done!\n")
}