package mlParser

import (
	"bufio"
	"fmt"
	"os"
)

// PROMPT console prompt to indicate accepting code
const PROMPT = "go>> "

func main() {
	var ret float64
	var err error
	var line string
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		if line = scanner.Text(); line == "exit" {
			os.Exit(0)
		}
		if ret, err = ParseAndEval(line); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(ret)
		}
	}
}
