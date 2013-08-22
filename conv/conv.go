package conv

import (
	"lisp"
	"lisp/parser"
)

func Int(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 1 {
		return lisp.None, lisp.ErrParaNum
	}
	u, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, err
	}
	switch u.Kind {
	case lisp.Int:
		return u, nil
	case lisp.Float:
		return lisp.Token{lisp.Int, int64(u.Text.(float64))}, nil
	case lisp.String:
		a, b := parser.ParseInt([]byte(u.Text.(string)))
		if b == 0 {
			return lisp.None, lisp.ErrNotConv
		}
		return lisp.Token{lisp.Int, a}, nil
	}
	return lisp.None, lisp.ErrFitType
}

func Float(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 1 {
		return lisp.None, lisp.ErrParaNum
	}
	u, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, err
	}
	switch u.Kind {
	case lisp.Int:
		return lisp.Token{lisp.Float, float64(u.Text.(int64))}, nil
	case lisp.Float:
		return u, nil
	case lisp.String:
		a, b := parser.ParseFloat([]byte(u.Text.(string)))
		if b == 0 {
			return lisp.None, lisp.ErrNotConv
		}
		return lisp.Token{lisp.Float, a}, nil
	}
	return lisp.None, lisp.ErrFitType
}


func List(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 1 {
		return lisp.None, lisp.ErrParaNum
	}
	u, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, err
	}
	if u.Kind != lisp.String {
		return lisp.None, lisp.ErrFitType
	}
	s := u.Text.(string)
	x := make([]lisp.Token, 0, len(s))
	for _, c := range s {
		x = append(x, lisp.Token{lisp.Int, int64(c)})
	}
	return lisp.Token{lisp.List, x}, nil
}

func String(t []lisp.Token, p *lisp.Lisp) (lisp.Token, error) {
	if len(t) != 1 {
		return lisp.None, lisp.ErrParaNum
	}
	u, err := p.Exec(t[0])
	if err != nil {
		return lisp.None, err
	}
	if u.Kind != lisp.List {
		return lisp.None, lisp.ErrFitType
	}
	s := u.Text.([]lisp.Token)
	x := make([]rune, 0, len(s))
	for _, c := range s {
		if c.Kind != lisp.Int {
			return lisp.None, lisp.ErrFitType
		}
		x = append(x, rune(c.Text.(int64)))
	}
	return lisp.Token{lisp.String, string(x)}, nil
}
