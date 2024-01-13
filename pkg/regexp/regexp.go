package regexp

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func Reg() {
	in := bufio.NewReader(os.Stdin)
	exprParse := YoYoNewParser()
	for {
		if _, err := os.Stdout.WriteString("> "); err != nil {
			log.Fatalf("WriteString: %s", err)
		}
		line, err := in.ReadBytes('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("ReadBytes: %s", err)
		}

		line = line[:len(line)-1]
		exprLex := NewLexer(line)
		exprParse.Parse(exprLex)
		fmt.Printf("line: %s, ast: \n%2v\n", line, exprLex.ast)
		//fmt.Println(exprLex.ast)
	}
}
