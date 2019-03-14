package repl

import (
	"bufio"
	"fmt"
	"github.com/TaigaMikami/monkey/lexer"
	"github.com/TaigaMikami/monkey/token"
	"io"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			// exit と入力されると終了
			if tok.Literal == "exit" {
				return
			}

			fmt.Printf("%+v\n", tok)
		}
	}
}