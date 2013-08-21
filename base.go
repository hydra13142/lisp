package lisp

import (
	"errors"
	"fmt"
	"lisp/parser"
)

type Kind int

type Name string

type Gfac func([]Token, *Lisp) (Token, error)

type Lfac struct {
	Para []Name
	Text []Token
	Make map[Name]Token
}

const (
	Null Kind = iota
	Int
	Float
	String
	Fold
	List
	Back
	Front
	Label
	Operator
)

var (
	pattern = &parser.Pattern{}

	True = Token{Int, int64(1)}
	None = Token{List, []Token(nil)}

	Global = &Lisp{env: map[Name]Token{}}

	ErrNotOver = errors.New("Cannot scan to the end")
	ErrUnquote = errors.New("Quote is unfold")
	ErrNotFind = errors.New("Not find this Name")
	ErrNotFunc = errors.New("Not a function")
	ErrParaNum = errors.New("Wrong parament number")
	ErrFitType = errors.New("Lisp type is wrong")
	ErrNotName = errors.New("This's not a Name")
	ErrIsEmpty = errors.New("Fold is empty")
	ErrNotConv = errors.New("Cannot translate")
)

func (t Kind) String() string {
	switch t {
	case Int:
		return "int"
	case Float:
		return "float"
	case String:
		return "string"
	case Fold:
		return "fold list"
	case List:
		return "list"
	case Back:
		return "go"
	case Front:
		return "lisp"
	case Label:
		return "Name"
	case Operator:
		return "operator"
	}
	return "unknown"
}

func (l Lfac) String() string {
	return fmt.Sprintf("{front : (%v,%v)}", l.Para, l.Text)
}
