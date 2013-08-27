package lisp

import "fmt"

type Token struct {
	Kind
	Text interface{}
}

func (t *Token) Bool() bool {
	switch t.Kind {
	case Null:
		return false
	case Int:
		return t.Text.(int64) != 0
	case Float:
		return t.Text.(float64) != 0
	case String:
		return t.Text.(string) != ""
	case Chan:
		return len(t.Text.(chan Token)) != 0
	case List:
		return len(t.Text.([]Token)) != 0
	}
	return true
}

func (t *Token) Eq(p *Token) bool {
	var (
		a, b []Token
		c, d []Name
	)
	if t.Kind != p.Kind {
		return false
	}
	switch t.Kind {
	case Null:
		return true
	case Int:
		return t.Text.(int64) == p.Text.(int64)
	case Float:
		return t.Text.(float64) == p.Text.(float64)
	case String:
		return t.Text.(string) == p.Text.(string)
	case Back, Chan:
		return false
	case Label:
		return t.Text.(Name) == p.Text.(Name)
	case Front:
		if t.Text.(*Lfac).Make != p.Text.(*Lfac).Make {
			return false
		}
		a, b = t.Text.(*Lfac).Text, p.Text.(*Lfac).Text
		c, d = t.Text.(*Lfac).Para, p.Text.(*Lfac).Para
	case Macro:
		a, b = t.Text.(*Hong).Text, p.Text.(*Hong).Text
		c, d = t.Text.(*Hong).Para, p.Text.(*Hong).Para
	case Fold, List:
		a, b = t.Text.([]Token), p.Text.([]Token)
	}
	m, n := len(a), len(b)
	if m != n {
		return false
	}
	for i := 0; i < m; i++ {
		if !a[i].Eq(&b[i]) {
			return false
		}
	}
	if c != nil {
		m, n := len(c), len(d)
		if m != n {
			return false
		}
		for i := 0; i < m; i++ {
			if c[i] != d[i] {
				return false
			}
		}
	}
	return false
}

func (t *Token) Cmp(p *Token) int {
	var a, b bool
	switch t.Kind {
	case Int:
		switch p.Kind {
		case Int:
			a = t.Text.(int64) > p.Text.(int64)
			b = t.Text.(int64) < p.Text.(int64)
		case Float:
			a = float64(t.Text.(int64)) > p.Text.(float64)
			b = float64(t.Text.(int64)) < p.Text.(float64)
		default:
			return 0
		}
	case Float:
		switch p.Kind {
		case Int:
			a = t.Text.(float64) > float64(p.Text.(int64))
			b = t.Text.(float64) < float64(p.Text.(int64))
		case Float:
			a = t.Text.(float64) > p.Text.(float64)
			b = t.Text.(float64) < p.Text.(float64)
		default:
			return 0
		}
	case String:
		switch p.Kind {
		case Int, Float:
			return 1
		case String:
			a = t.Text.(string) > p.Text.(string)
			b = t.Text.(string) < p.Text.(string)
		default:
			return 0
		}
	case List:
		switch p.Kind {
		case Int, Float, String:
			return 1
		case List:
			x, y := t.Text.([]Token), p.Text.([]Token)
			m, n := len(x), len(y)
			for i := 0; i < m && i < n; i++ {
				j := x[i].Cmp(&y[i])
				if j != 0 {
					return j
				}
			}
			a = m > n
			b = m < n
		default:
			return 0
		}
	default:
		return 0
	}
	if a {
		return +1
	}
	if b {
		return -1
	}
	return 0
}

func (t Token) String() string {
	switch t.Kind {
	case Null:
		return ""
	case Chan:
		return "channel"
	default:
		return fmt.Sprint(t.Text)
	}
}
