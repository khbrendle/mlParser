package mlParser

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"math"
	"reflect"
	"strconv"
)

var (
	// ErrNotImplemented used not not implemented AST expressions
	ErrNotImplemented error = errors.New("mlParser : AST not currently implemented")
	// ErrUnhandledOperator used for not supporting mathematical operators
	ErrUnhandledOperator error = errors.New("mlParser : operator not currently implemented")
	// ErrParseFormula used for identifying error in parsing
	ErrParseFormula error = errors.New("mlParser : error parsing formula")
	// ErrEvaluation used for identifying error in evaluation
	ErrEvaluation error = errors.New("mlParser : error evaluating expression")
)

// ErrWrap simple error wrapper to add package name and help build stack trace
func ErrWrap(err error) error {
	return fmt.Errorf("mlParser : %s", err.Error())
}

// Eval is used to evaluate mathematical expressions
func Eval(exp ast.Expr) (float64, error) {
	if debug {
		fmt.Printf("%+v\n", exp)
	}
	switch exp := exp.(type) {
	case *ast.BinaryExpr:
		return EvalBinaryExpr(exp)
	case *ast.ParenExpr:
		return EvalParenExpr(exp)
	case *ast.CallExpr:
		return EvalCallExpr(exp)
	case *ast.UnaryExpr:
		return EvalUnaryExpr(exp)
	case *ast.BasicLit:
		switch exp.Kind {
		case token.INT, token.FLOAT:
			var i float64
			var err error
			if i, err = strconv.ParseFloat(exp.Value, 64); err != nil {
				return math.NaN(), ErrWrap(err)
			}
			return i, nil
		}
	default:
		fmt.Printf("case not handled for %v", reflect.TypeOf(exp))
		return math.NaN(), ErrNotImplemented
	}

	return math.NaN(), ErrNotImplemented
}

// EvalBinaryExpr handles presence of binary expression in AST
func EvalBinaryExpr(exp *ast.BinaryExpr) (float64, error) {
	var left, right float64
	var err error
	if left, err = Eval(exp.X); err != nil {
		return math.NaN(), ErrWrap(err)
	}
	if right, err = Eval(exp.Y); err != nil {
		return math.NaN(), ErrWrap(err)
	}

	switch exp.Op {
	case token.ADD:
		return left + right, nil
	case token.SUB:
		return left - right, nil
	case token.MUL:
		return left * right, nil
	case token.QUO:
		return left / right, nil
	case token.REM:
		return math.Remainder(left, right), nil
	case token.XOR:
		return math.Pow(left, right), nil
	}
	fmt.Printf("operator `%v` not currently implemented", exp.Op)
	return math.NaN(), ErrUnhandledOperator
}

// EvalParenExpr handles presence of parentheses expression in AST
func EvalParenExpr(exp *ast.ParenExpr) (float64, error) {
	x := *exp
	return Eval(x.X)
}

// EvalCallExpr handles presence of call expression in AST
func EvalCallExpr(exp *ast.CallExpr) (float64, error) {
	var err error
	e := *exp
	args := e.Args
	argMap := make(map[int]float64)
	for i, arg := range args {
		if argMap[i], err = Eval(arg); err != nil {
			return math.NaN(), ErrWrap(err)
		}
	}
	switch x := e.Fun.(type) {
	case *ast.Ident:
		if debug {
			fmt.Printf("ident: %+v\n", e.Fun.(*ast.Ident))
			fmt.Printf("%v\n", x)
		}
		f := e.Fun.(*ast.Ident)
		// fmt.Printf("string %v\n", tmp.String())
		// fmt.Printf("name %v\n", tmp.Name)
		switch f.String() {
		case "exp":
			return math.Exp(argMap[0]), nil
		default:
			fmt.Printf("function `%s` not currently implemented", f.String())
		}
	case *ast.SelectorExpr:
		if debug {
			fmt.Printf("selector expr: %+v\n", e.Fun.(*ast.SelectorExpr))
			fmt.Printf("%v\n", x)
		}
		L := e.Fun.(*ast.SelectorExpr)
		l := *L
		pkgName := l.X
		funcName := l.Sel
		f := fmt.Sprintf("%s.%s", pkgName, funcName)
		switch f {
		case "math.Log":
			return math.Log(argMap[0]), nil
		case "math.Exp":
			return math.Exp(argMap[0]), nil
		case "math.Abs":
			return math.Abs(argMap[0]), nil
		default:
			fmt.Printf("function `%s` not currently implemented", f)
		}
	}

	return math.NaN(), ErrNotImplemented
}

// EvalUnaryExpr handles presence of unary expression in AST, negative numbers
func EvalUnaryExpr(exp *ast.UnaryExpr) (float64, error) {
	x0 := exp.X.(*ast.BasicLit)
	x := *x0
	var val float64
	var err error
	if val, err = strconv.ParseFloat(x.Value, 64); err != nil {
		fmt.Printf("cannot parse `%s` as float64\n", x.Value)
		return math.NaN(), ErrUnhandledOperator
	}
	switch exp.Op {
	case token.ADD:
		return val, nil
	case token.SUB:
		return val * -1, nil
	}
	fmt.Printf("unary expression `%s` not handled in EvalUnaryExpr\n", exp.Op)
	return math.NaN(), ErrNotImplemented
}

// ParseAndEval is wrapper for multi-step parse and evaluate
func ParseAndEval(formula string) (float64, error) {
	var exp ast.Expr
	var err error
	if exp, err = parser.ParseExpr(formula); err != nil {
		fmt.Printf("error parsing formula %s:\n\t%s", formula, err.Error())
		return math.NaN(), ErrParseFormula
	}
	if debug {
		fmt.Println("ParseAndEval: check - 0")
	}
	var res float64
	if res, err = Eval(exp); err != nil {
		fmt.Printf("error evaluating ast from formula %s:\n\t%s", formula, err.Error())
		return math.NaN(), ErrWrap(err)
	}
	if debug {
		fmt.Println("ParseAndEval: check - 1")
	}
	return res, nil
}
