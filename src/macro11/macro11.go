package macro11

import (
	"io"
	"bufio"
	"fmt"
	"regexp"
	"strings"
)


type instruction struct {
	bits int
}

type symboltable struct {

}


func Assemble(input io.Reader) ([]byte, []byte, error) {
	lines, err := preprocess(input)
	if err != nil {
		return nil, nil, err
	}

	instructions, symbols, err := firstPass(lines)
	if err != nil {
		return nil, nil, err
	}

	output, listing, err := secondPass(instructions, symbols)
	if err != nil {
		return nil, nil, err
	} 

	return output, listing, nil
}

func preprocess(input io.Reader) ([]string, error) {
	lines := make([]string, 5)

	r, _ := regexp.Compile("\\s{2,}")

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		stripped_line := r.ReplaceAllString(line, " ")
		comment_idx := strings.Index(stripped_line, ";")
		if comment_idx != -1 {
			stripped_line = stripped_line[:comment_idx]
		}

		lines = append(lines, stripped_line)
	}

	fmt.Println("lines: ", lines)

	return lines, nil
}


func firstPass(lines []string) ([]instruction, symboltable, error) {
	var instructions []instruction;
	var symbols symboltable;
	

	return instructions, symbols, nil
}

func secondPass(instructions []instruction, symbols symboltable) ([]byte, []byte, error) {
	var output []byte;
	var listing []byte;

	return output, listing, nil
}