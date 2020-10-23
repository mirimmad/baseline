package main

import (
	"fmt"

	"github.com/baseline/ast"
)

func main() {

	call := ast.Call{"Fun", []ast.AST{&ast.Number{14}}}
	fmt.Println(call)

}
