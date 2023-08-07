// repl reads from the input source until encountering a newline, takes the
// just read line and passes it to the lexer instance, and then prints out
// all the tokens the lexer gives us until it encounters the end of the file.

package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/GStechschulte/go-interpreter/src/monkey/lexer"
	"github.com/GStechschulte/go-interpreter/src/monkey/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	// infinite loop that prints the prompt and scans the input
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
