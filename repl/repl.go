package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/lexer"
	"monkey/token"
)

func Start(in io.Reader, out io.Writer) {
	prompt_in := "monke<< "
	prompt_out := "monke > "

	for line := ""; ; {
		// fmt.Println(line)
		fmt.Printf("%v", prompt_in)
		scanner := bufio.NewScanner(in)
		if !scanner.Scan() {
			return
		}
		line = scanner.Text()
		if line == "bye" {
			fmt.Fprintf(out, "%v%v\n", prompt_out, "Goodbye !")
			return
		}
		lex := lexer.NewLexer(line)

		for tok := lex.NextToken(); tok.Type != token.EOF; tok = lex.NextToken() {
			fmt.Fprintf(out, "%v %+v\n", prompt_out, tok)
		}
	}
}
