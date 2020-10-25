package ast

import (
	"fmt"
)

type AST interface {
	String() string
}

type Number struct {
	Value int64
}

func (n *Number) String() string  {
	return fmt.Sprintf("%d", n.Value)
}

type Id struct {
	value string
}

func (i *Id) String() string {
	return ""
}

type Not struct {
	term AST
}

func (n *Not) String() string {
	return ""
}

type String struct {
	left  AST
	right AST
}

func (e *String) String() string {
	return ""
}

type Add struct {
	left  AST
	right AST
}

func (a *Add) String() string {
	return ""
}

type Subtract struct {
	left  AST
	right AST
}

func (s *Subtract) String() string {
	return ""
}

type Multiply struct {
	left  AST
	right AST
}

func (m *Multiply) String() string {
	return ""
}

type Divide struct {
	left  AST
	right AST
}

func (d *Divide) String() string {
	return ""
}

type Call struct {
	Callee string
	Args   []AST
}

func (c *Call) String() string {
	return ""
}

type Return struct {
	term AST
}

func (r *Return) String() string {
	return ""
}

type Block struct {
	statements []AST
}

func (b *Block) String() string {
	return ""
}

type IF struct {
	condition   AST
	consequence AST
	alternative AST
}

func (i *IF) String() string {
	return ""
}

type Defintion struct {
	name       string
	parameters []string
	body       AST
}

func (d *Defintion) String() string {
	return ""
}

type Var struct {
	name  string
	value AST
}

func (v *Var) String() string {
	return ""
}

type Assignment struct {
	name  string
	value AST
}

func (a *Assignment) String() string {
	return ""
}

type While struct {
	condition AST
	body      AST
}

func (w *While) String() string {
	return ""
}
