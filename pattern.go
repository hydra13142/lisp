package lisp

import . "github.com/hydra13142/lisp/parser"

func init() {
	bnd := func(c byte) bool {
		return c == '(' || c == ')' || IsSpace(c)
	}
	pattern.Add(func(s []byte) (interface{}, int) {
		if len(s) > 0 {
			switch s[0] {
			case '(', ')':
				return s[0], 1
			case '\'':
				if len(s) > 2 && s[1] == '(' && s[2] != '\'' {
					return byte('['), 2
				}
			}
		}
		return nil, 0
	})
	pattern.Add(func(s []byte) (interface{}, int) {
		a, i := ParseInt(s)
		if i > 0 {
			if _, j := ParseFloat(s); i == j && (i >= len(s) || bnd(s[i])) {
				return a, i
			}
		}
		return nil, 0
	})
	pattern.Add(func(s []byte) (interface{}, int) {
		a, i := ParseFloat(s)
		if i > 0 && (i >= len(s) || bnd(s[i])) {
			return a, i
		}
		return nil, 0
	})
	pattern.Add(func(s []byte) (interface{}, int) {
		a, i := ParseChar(s)
		if i > 0 && (i >= len(s) || bnd(s[i])) {
			return int64(a), i
		}
		return nil, 0
	})
	pattern.Add(func(s []byte) (interface{}, int) {
		a, i := ParseString(s)
		if i > 0 && (i >= len(s) || bnd(s[i])) {
			return a, i
		}
		return nil, 0
	})
	pattern.Add(func(s []byte) (interface{}, int) {
		i := 0
		for i < len(s) && !bnd(s[i]) {
			i++
		}
		a := Name(string(s[:i]))
		return a, i
	})
}
