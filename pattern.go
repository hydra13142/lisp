package lisp

import "github.com/hydra13142/parser"

func init() {
	bnd := func(c byte) bool {
		return c == '(' || c == ')' || parser.IsSpace(c)
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
		a, i := parser.ParseInt(s)
		if i > 0 {
			if _, j := parser.ParseFloat(s); i == j && (i >= len(s) || bnd(s[i])) {
				return a, i
			}
		}
		return nil, 0
	})
	pattern.Add(func(s []byte) (interface{}, int) {
		a, i := parser.ParseFloat(s)
		if i > 0 && (i >= len(s) || bnd(s[i])) {
			return a, i
		}
		return nil, 0
	})
	pattern.Add(func(s []byte) (interface{}, int) {
		if len(s) == 0 {
			return nil, 0
		}
		if s[0] == '\'' {
			a, i := parser.ParseChar(s[1:])
			if i > 0 {
				i += 1
				if i < len(s) && s[i] == '\'' {
					i += 1
					if i >= len(s) || bnd(s[i]) {
						return int64(a), i
					}
				}
			}
		}
		return nil, 0
	})
	pattern.Add(func(s []byte) (interface{}, int) {
		i := 1
		if len(s) == 0 || s[0] != '"' {
			return nil, 0
		}
		for ; i < len(s) && s[i] != '"'; i++ {
			if s[i] == '\\' {
				i++
			}
		}
		if i < len(s) {
			a := string(s[1:i])
			i += 1
			if i >= len(s) || bnd(s[i]) {
				return a, i
			}
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
