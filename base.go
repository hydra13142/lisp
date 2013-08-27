package lisp

import (
	"errors"
	"fmt"
	"github.com/hydra13142/lisp/parser"
)

type Kind int

type Name string

type Gfac func([]Token, *Lisp) (Token, error)

type Hong struct {
	Para []Name
	Text []Token
	Real []Name
}

type Lfac struct {
	Para []Name
	Text []Token
	Make *Lisp
}

const (
	Null Kind = iota
	Int
	Float
	String
	Chan
	Fold
	List
	Back
	Macro
	Front
	Label
	Operator
)

var (
	pattern = &parser.Pattern{}

	True  = Token{Int, int64(1)}
	False = Token{List, []Token(nil)}
	None  = Token{}

	Global = &Lisp{env: map[Name]Token{}, dad: nil}

	ErrNotOver = errors.New("Cannot scan to the end")
	ErrUnquote = errors.New("Quote is unfold")
	ErrNotFind = errors.New("Not find this Name")
	ErrNotFunc = errors.New("Not a function")
	ErrParaNum = errors.New("Wrong parament number")
	ErrFitType = errors.New("Lisp type is wrong")
	ErrNotName = errors.New("This's not a Name")
	ErrIsEmpty = errors.New("Fold is empty")
	ErrNotConv = errors.New("Cannot translate")
	ErrRefused = errors.New("Can't remove a back function")
	ErrIsClose = errors.New("Channel has been closed")
)

func (t Kind) String() string {
	switch t {
	case Int:
		return "int"
	case Float:
		return "float"
	case String:
		return "string"
	case Chan:
		return "channel"
	case Fold:
		return "fold list"
	case List:
		return "list"
	case Back:
		return "go"
	case Macro:
		return "macro"
	case Front:
		return "lisp"
	case Label:
		return "Name"
	case Operator:
		return "operator"
	}
	return "unknown"
}

func (m Hong) String() string {
	return fmt.Sprintf("{macro : %v | %v => %v}", m.Para, m.Real, m.Text)
}

func (l Lfac) String() string {
	return fmt.Sprintf("{front : %v => %v}", l.Para, l.Text)
}
