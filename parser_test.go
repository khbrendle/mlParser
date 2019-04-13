package mlParser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

// ParserTestArgs ...
type ParserTestArgs struct {
	Input  string  `json:"input"`
	Output float64 `json:"output"`
}

// InitParserTestArgs ...
func InitParserTestArgs(filePath string) ([]ParserTestArgs, error) {
	var jsonFile *os.File
	var err error
	// Open our jsonFile
	if jsonFile, err = os.Open(filePath); err != nil {
		return []ParserTestArgs{}, err
	}
	// read our opened xmlFile as a byte array.
	var byteValue []byte
	if byteValue, err = ioutil.ReadAll(jsonFile); err != nil {
		return []ParserTestArgs{}, err
	}

	var testArgs []ParserTestArgs
	if err = json.Unmarshal(byteValue, &testArgs); err != nil {
		// handleError("Error unmarshalling test data", err)
		return []ParserTestArgs{}, err
	}

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	return testArgs, err
}

func TestEval(t *testing.T) {
	var testArgs []ParserTestArgs
	var err error
	if testArgs, err = InitParserTestArgs("parser_tests.json"); err != nil {
		t.Errorf("error reading test args: %s", err)
	}

	var r float64
	for _, a := range testArgs {
		if debug {
			fmt.Println(a.Input)
		}
		if r, err = ParseAndEval(a.Input); err != nil {
			t.Errorf("error evaluating expression: %s", err)
		} else if r != a.Output {
			t.Errorf("Error in formula %s, expected %f, got %f", a.Input, a.Output, r)
		}
	}
}

// func BenchmarkExpression(b *testing.B) {
// 	var err error
// 	var ret float64
// 	formula := "((1)) * 3.48315304084922 + (3.5) * 0.42707287702065 + (0.2) * 1.261487384549 + (math.Exp(1.4)) * 0.00132690681446317 + (3.5 * 0.2) * -0.202511439852321 + (3.5 * math.Exp(1.4)) * 0.000231099657391578"
// 	b.Run("EvalBench", func(b *testing.B) {
// 		for i := 0; i < b.N; i++ {
// 			if ret, err = ParseAndEval(formula); err != nil {
// 				fmt.Println(err)
// 			}
// 		}
// 	})
// 	fmt.Println(ret)
// }
