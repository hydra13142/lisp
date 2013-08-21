package lisp

import "fmt"

type Token struct {
	Kind
	Text interface{}
}

func (t *Token) Bool() bool {
	switch t.Kind {
	case Int:
		return t.Text.(int64) != 0
	case Float:
		return t.Text.(float64) != 0
	case String:
		return t.Text.(string) != ""
	case List:
		return len(t.Text.([]Token)) != 0
	case Null:
		return false
	}
	return true
}

func (t *Token) Eq(p *Token) bool {
	if t.Kind != p.Kind {
		return false
	}
	switch t.Kind {
	case Int:
		return t.Text.(int64) == p.Text.(int64)
	case Float:
		return t.Text.(float64) == p.Text.(float64)
	case String:
		return t.Text.(string) == p.Text.(string)
	case Fold, List:
		a, b := t.Text.([]Token), p.Text.([]Token)
		m, n := len(a), len(b)
		for i := 0; i < m && i < n; i++ {
			j := a[i].Eq(&b[i])
			if !j {
				return false
			}
		}
		return true
	default:
		return false
	}
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
		return 1
	}
	if b {
		return -1
	}
	return 0
}

func (t Token) String() string {
	return fmt.Sprint(t.Text)
}
