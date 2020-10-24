package main

import (
	"fmt"

	//"github.com/baseline/ast"
	"github.com/baseline/parser"
)

func main() {

	//call := ast.Call{"Fun", []ast.AST{&ast.Number{14}}}
	s := parser.Source{"yello World", 0}
	p := parser.NewParser()
	fmt.Println(p.Regex("hello").Or(p.Regex("Hello")).Parse(s))

}
