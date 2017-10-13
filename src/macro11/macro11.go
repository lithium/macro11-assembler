package macro11

type instruction struct {
	bits int
}

type symboltable struct {

}


func Assemble(inputSource []byte) ([]byte, []byte, error) {
	lines, err := preprocess(inputSource)
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

func preprocess(inputSource []byte) ([]string, error) {
	var lines []string;

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