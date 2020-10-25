package ast

import (
	"bytes"
	"fmt"
)

type AST interface {
	String() string
}

type Number struct {
	Value int64
}

func (n *Number) String() string {
	return fmt.Sprintf("%d", n.Value)
}

type Id struct {
	Value string
}

func (i *Id) String() string {
	return fmt.Sprintf("%s", i.Value)
}

type Not struct {
	Term AST
}

func (n *Not) String() string {
	return fmt.Sprintf("Not [%s]", n.Term.String())
}

type Equal struct {
	Left  AST
	Right AST
}

func (e *Equal) String() string {
	return fmt.Sprintf("Eq [%s %s]", e.Left.String(), e.Right.String())
}

type Add struct {
	Left  AST
	Right AST
}

func (a *Add) String() string {
	return fmt.Sprintf("Add [%s %s]", a.Left.String(), a.Right.String())
}

type Subtract struct {
	Left  AST
	Right AST
}

func (s *Subtract) String() string {
	return fmt.Sprintf("Subt [%s %s]", s.Left.String(), s.Right.String())
}

type Multiply struct {
	Left  AST
	Right AST
}

func (m *Multiply) String() string {
	return fmt.Sprintf("Mul [%s %s]", m.Left.String(), m.Right.String())
}

type Divide struct {
	Left  AST
	Right AST
}

func (d *Divide) String() string {
	return fmt.Sprintf("Div [%s %s]", d.Left.String(), d.Right.String())
}

type Call struct {
	Callee string
	Args   []AST
}

func (c *Call) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("Call ([%s] [", c.Callee))
	for _, v := range c.Args {
		buffer.WriteString(v.String())
	}
	buffer.WriteString("])")
	return buffer.String()

}

type Return struct {
	Term AST
}

func (r *Return) String() string {
	return fmt.Sprintf("Return [%s]", r.Term.String())
}

type Block struct {
	Statements []AST
}

func (b *Block) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("Block [\n")
	for _, v := range b.Statements {
		buffer.WriteString(v.String())
		buffer.WriteString("\n")
	}
	buffer.WriteString("]")
	return buffer.String()
}

type IF struct {
	Condition   AST
	Consequence AST
	Alternative AST
}

func (i *IF) String() string {
	return fmt.Sprintf("IF [ %s %s %s ]", i.Condition.String(), i.Consequence.String(), i.Alternative.String())
}

type Defintion struct {
	Name       string
	Parameters []string
	Body       AST
}

func (d *Defintion) String() string {
	var buffer bytes.Buffer
	buffer.WriteString("Def [Name: ")
	buffer.WriteString(d.Name)
	buffer.WriteString("Params: ")
	for _, v := range d.Parameters {
		buffer.WriteString(v)
		buffer.WriteString(" ")
	}
	buffer.WriteString("] Body: ")
	buffer.WriteString(d.Body.String())
	return buffer.String()

}

type Var struct {
	Name  string
	Value AST
}

func (v *Var) String() string {
	return fmt.Sprintf("VAR [%s %s]", v.Name, v.Value.String())
}

type Assignment struct {
	Name  string
	Value AST
}

func (a *Assignment) String() string {
	return fmt.Sprintf("Asssing [%s %s]", a.Name, a.Value.String())
}

type While struct {
	Condition AST
	Body      AST
}

func (w *While) String() string {
	return fmt.Sprintf("While [%s %s]", w.Condition.String(), w.Body.String())
}
