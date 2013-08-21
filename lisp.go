package lisp

import (
	"errors"
	"fmt"
	. "lisp/parser"
)

type Kind int

type null struct{}

type name string

type pres func([]Token, *Lisp) (Token, error)

type afts struct {
	Para []name
	Text []Token
	Make map[name]Token
}

type Token struct {
	Kind
	Text interface{}
}

type Lisp struct {
	dad *Lisp
	env map[name]Token
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
	Name
	Operator
)

var pattern = &Pattern{}

var (
	True = Token{Int, int64(1)}
	None = Token{List, []Token(nil)}
)

var Global = &Lisp{env: map[name]Token{}}

var (
	ErrNotOver = errors.New("Cannot scan to the end")
	ErrUnquote = errors.New("Quote is unfold")
	ErrNotFind = errors.New("Not find this name")
	ErrNotFunc = errors.New("Not a function")
	ErrParaNum = errors.New("Wrong parament number")
	ErrFitType = errors.New("Lisp type is wrong")
	ErrNotName = errors.New("This's not a name")
	ErrIsEmpty = errors.New("Fold is empty")
)

func Scan(s string) (list []Token, err error) {
	scanner := pattern.NewScanner(s, true)
	list = make([]Token, 0, 100)
	for {
		a, b, c := scanner.Scan()
		if c != nil {
			break
		}
		switch b {
		case 1:
			list = append(list, Token{Operator, a})
		case 2:
			list = append(list, Token{Int, a})
		case 3:
			list = append(list, Token{Float, a})
		case 4:
			list = append(list, Token{Int, a})
		case 5:
			list = append(list, Token{String, a})
		case 6:
			list = append(list, Token{Name, a})
		}
	}
	if !scanner.Over() {
		err = ErrNotOver
	}
	return
}

func Tree(tkn []Token) ([]Token, error) {
	var f Token
	var s int
	if len(tkn) == 0 {
		return nil, nil
	}
	if tkn[0].Kind == Operator {
		var t bool
		switch tkn[0].Text.(byte) {
		case '(':
			t = true
		case '[':
			t = false
		default:
			return nil, ErrUnquote
		}
		i, j, l := 1, 1, len(tkn)
		for i < l && j > 0 {
			if tkn[i].Kind == Operator {
				switch tkn[i].Text.(byte) {
				case '(', '[':
					j++
				case ')':
					j--
				}
			}
			i++
		}
		if j <= 0 {
			fold, err := Tree(tkn[1 : i-1])
			if err != nil {
				return nil, err
			}
			if t {
				f = Token{Text: fold, Kind: List}
			} else {
				f = Token{Text: fold, Kind: Fold}
			}
			s = i
		} else {
			return nil, ErrUnquote
		}
	} else {
		f = tkn[0]
		s = 1
	}
	rest, err := Tree(tkn[s:])
	if err != nil {
		return nil, err
	}
	ans := make([]Token, 1+len(rest))
	ans[0] = f
	copy(ans[1:], rest)
	return ans, nil
}

func (t Kind) String() string {
	switch t {
	case Null:
		return "null"
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
	case Name:
		return "name"
	case Operator:
		return "operator"
	}
	return "unknown"
}

func (l afts) String() string {
	return fmt.Sprintf("{front : (%v,%v)}", l.Para, l.Text)
}

func (t Token) String() string {
	return fmt.Sprint(t.Text)
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

func NewLisp() *Lisp {
	x := new(Lisp)
	x.env = map[name]Token{}
	x.dad = nil
	for i, f := range Global.env {
		x.env[i] = f
	}
	return x
}

func (l *Lisp) Add(s string, f func([]Token, *Lisp) (Token, error)) {
	l.env[name(s)] = Token{Back, pres(f)}
}

func (l *Lisp) Exec(f Token) (ans Token, err error) {
	var (
		ls []Token
		ct Token
		ok bool
	)
	switch f.Kind {
	case Fold:
		return Token{List, f.Text.([]Token)}, nil
	case Name:
		for v := l; ; {
			ct, ok = v.env[f.Text.(name)]
			if ok {
				break
			}
			v = v.dad
			if v == nil {
				break
			}
		}
		if !ok {
			return None, ErrNotFind
		}
		return ct, nil
	case List:
		ls = f.Text.([]Token)
		if len(ls) == 0 {
			return None, nil
		}
		ct = ls[0]
		switch ct.Kind {
		case Name:
			for v := l; ; {
				ct, ok = v.env[ls[0].Text.(name)]
				if ok {
					break
				}
				v = v.dad
				if v == nil {
					break
				}
			}
			if !ok {
				return None, ErrNotFind
			}
		case List:
			ct, err = l.Exec(ct)
			if err != nil {
				return None, err
			}
		}
		switch ct.Kind {
		case Back:
			return ct.Text.(pres)(ls[1:], l)
		case Front:
			lp := ct.Text.(afts)
			q := &Lisp{dad: l, env: map[name]Token{}}
			if len(ls) != len(lp.Para)+1 {
				return None, ErrParaNum
			}
			for m, n := range lp.Make {
				q.env[m] = n
			}
			for i, t := range ls[1:] {
				q.env[lp.Para[i]], err = l.Exec(t)
				if err != nil {
					return None, err
				}
			}
			return q.Exec(Token{Text: lp.Text, Kind: List})
		default:
			return None, ErrNotFunc
		}
	default:
		return f, nil
	}
	return None, nil
}

func (l *Lisp) Eval(s string) (Token, error) {
	var (
		a, b []Token
		c, d Token
		e    error
	)
	a, e = Scan(s)
	if e != nil {
		return None, e
	}
	b, e = Tree(a)
	if e != nil {
		return None, e
	}
	for _, c = range b {
		d, e = l.Exec(c)
		if e != nil {
			return None, e
		}
	}
	return d, nil
}

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
		a := name(string(s[:i]))
		return a, i
	})
	Global.Add("quote", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 1 {
			return None, ErrParaNum
		}
		if t[0].Kind == Name {
			return p.Exec(t[0])
		}
		return t[0], nil
	})
	Global.Add("eval", func(t []Token, p *Lisp) (ans Token, err error) {
		if len(t) != 1 {
			return None, ErrParaNum
		}
		ans = t[0]
		if ans.Kind == Name {
			ans, err = p.Exec(ans)
			if err != nil {
				return None, err
			}
		}
		return p.Exec(ans)
	})
	Global.Add("atom", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 1 {
			return None, ErrParaNum
		}
		x, err := p.Exec(t[0])
		if err != nil {
			return None, err
		}
		if x.Kind != List || len(x.Text.([]Token)) == 0 {
			return True, nil
		} else {
			return None, nil
		}
	})
	Global.Add("eq", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 2 {
			return None, ErrParaNum
		}
		x, err := p.Exec(t[0])
		if err != nil {
			return None, err
		}
		y, err := p.Exec(t[1])
		if err != nil {
			return None, err
		}
		if x.Cmp(&y) == 0 {
			return True, nil
		} else {
			return None, nil
		}
	})
	Global.Add("car", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 1 {
			return None, ErrParaNum
		}
		x, err := p.Exec(t[0])
		if err != nil {
			return None, err
		}
		if x.Kind == List {
			if len(x.Text.([]Token)) != 0 {
				return x.Text.([]Token)[0], nil
			}
			return None, ErrIsEmpty
		}
		return None, ErrFitType
	})
	Global.Add("cdr", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 1 {
			return None, ErrParaNum
		}
		x, err := p.Exec(t[0])
		if err != nil {
			return None, err
		}
		if x.Kind == List {
			if len(x.Text.([]Token)) != 0 {
				return Token{List, x.Text.([]Token)[1:]}, nil
			}
			return None, ErrIsEmpty
		}
		return None, ErrFitType
	})
	Global.Add("cons", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 2 {
			return None, ErrParaNum
		}
		x, err := p.Exec(t[0])
		if err != nil {
			return None, err
		}
		y, err := p.Exec(t[1])
		if err != nil {
			return None, err
		}
		if y.Kind == List {
			a := y.Text.([]Token)
			b := make([]Token, len(a)+1)
			b[0] = x
			copy(b[1:], a)
			return Token{List, b}, nil
		}
		return None, ErrFitType
	})
	Global.Add("cond", func(t []Token, p *Lisp) (ans Token, err error) {
		if len(t) == 0 {
			return None, ErrParaNum
		}
		for _, i := range t {
			if i.Kind == List {
				t := i.Text.([]Token)
				if len(t) == 2 {
					ans, err = p.Exec(t[0])
					if err != nil {
						return None, err
					}
					if ans.Bool() {
						return p.Exec(t[1])
					}
					continue
				}
			}
			return None, ErrFitType
		}
		return None, nil
	})
	Global.Add("each", func(t []Token, p *Lisp) (ans Token, err error) {
		if len(t) == 0 {
			return None, ErrParaNum
		}
		for _, i := range t {
			ans, err = p.Exec(i)
			if err != nil {
				break
			}
		}
		return ans, err
	})
	Global.Add("lambda", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 2 {
			return None, ErrParaNum
		}
		a, b := t[0], t[1]
		if a.Kind != List {
			return None, ErrFitType
		}
		if b.Kind != List {
			return None, ErrFitType
		}
		t = a.Text.([]Token)
		x := make([]name, 0, len(t))
		for _, i := range t {
			if i.Kind != Name {
				return None, ErrNotName
			}
			x = append(x, i.Text.(name))
		}
		u := make(map[name]Token)
		for i, j := range p.env {
			u[i] = j
		}
		return Token{Front, afts{x, b.Text.([]Token), u}}, nil
	})
	Global.Add("define", func(t []Token, p *Lisp) (ans Token, err error) {
		if len(t) != 2 {
			return None, ErrParaNum
		}
		a, b := t[0], t[1]
		switch a.Kind {
		case Name:
			ans, err = p.Exec(b)
			if err == nil {
				p.env[a.Text.(name)] = ans
			}
			return ans, err
		case List:
			if b.Kind != List {
				return None, ErrFitType
			}
			t = a.Text.([]Token)
			x := make([]name, 0, len(t))
			for _, i := range t {
				if i.Kind != Name {
					return None, ErrNotName
				}
				x = append(x, i.Text.(name))
			}
			u := make(map[name]Token)
			for i, j := range p.env {
				u[i] = j
			}
			ans = Token{Front, afts{x[1:], b.Text.([]Token), u}}
			p.env[x[0]] = ans
			return ans, nil
		}
		return None, ErrFitType
	})
	Global.Add("+", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 2 {
			return None, ErrParaNum
		}
		x, err := p.Exec(t[0])
		if err != nil {
			return None, err
		}
		y, err := p.Exec(t[1])
		if err != nil {
			return None, err
		}
		switch x.Kind {
		case Int:
			switch y.Kind {
			case Int:
				return Token{Int, x.Text.(int64) + y.Text.(int64)}, nil
			case Float:
				return Token{Float, float64(x.Text.(int64)) + y.Text.(float64)}, nil
			}
		case Float:
			switch y.Kind {
			case Int:
				return Token{Float, x.Text.(float64) + float64(y.Text.(int64))}, nil
			case Float:
				return Token{Float, x.Text.(float64) + y.Text.(float64)}, nil
			}
		case String:
			switch y.Kind {
			case String:
				return Token{String, x.Text.(string) + y.Text.(string)}, nil
			}
		case List:
			switch y.Kind {
			case List:
				a, b := x.Text.([]Token), y.Text.([]Token)
				c := make([]Token, len(a)+len(b))
				copy(c, a)
				copy(c[len(a):], b)
				return Token{List, c}, nil
			}

		}
		return None, ErrFitType
	})
	Global.Add("-", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 2 {
			return None, ErrParaNum
		}
		x, err := p.Exec(t[0])
		if err != nil {
			return None, err
		}
		y, err := p.Exec(t[1])
		if err != nil {
			return None, err
		}
		switch x.Kind {
		case Int:
			switch y.Kind {
			case Int:
				return Token{Int, x.Text.(int64) - y.Text.(int64)}, nil
			case Float:
				return Token{Float, float64(x.Text.(int64)) - y.Text.(float64)}, nil
			}
		case Float:
			switch y.Kind {
			case Int:
				return Token{Float, x.Text.(float64) - float64(y.Text.(int64))}, nil
			case Float:
				return Token{Float, x.Text.(float64) - y.Text.(float64)}, nil
			}
		}
		return None, ErrFitType
	})
	Global.Add("*", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 2 {
			return None, ErrParaNum
		}
		x, err := p.Exec(t[0])
		if err != nil {
			return None, err
		}
		y, err := p.Exec(t[1])
		if err != nil {
			return None, err
		}
		switch x.Kind {
		case Int:
			switch y.Kind {
			case Int:
				return Token{Int, x.Text.(int64) * y.Text.(int64)}, nil
			case Float:
				return Token{Float, float64(x.Text.(int64)) * y.Text.(float64)}, nil
			}
		case Float:
			switch y.Kind {
			case Int:
				return Token{Float, x.Text.(float64) * float64(y.Text.(int64))}, nil
			case Float:
				return Token{Float, x.Text.(float64) * y.Text.(float64)}, nil
			}
		}
		return None, ErrFitType
	})
	Global.Add("/", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 2 {
			return None, ErrParaNum
		}
		x, err := p.Exec(t[0])
		if err != nil {
			return None, err
		}
		y, err := p.Exec(t[1])
		if err != nil {
			return None, err
		}
		switch x.Kind {
		case Int:
			switch y.Kind {
			case Int:
				return Token{Int, x.Text.(int64) / y.Text.(int64)}, nil
			case Float:
				return Token{Float, float64(x.Text.(int64)) / y.Text.(float64)}, nil
			}
		case Float:
			switch y.Kind {
			case Int:
				return Token{Float, x.Text.(float64) / float64(y.Text.(int64))}, nil
			case Float:
				return Token{Float, x.Text.(float64) / y.Text.(float64)}, nil
			}
		}
		return None, ErrFitType
	})
	Global.Add("%", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 2 {
			return None, ErrParaNum
		}
		x, err := p.Exec(t[0])
		if err != nil {
			return None, err
		}
		y, err := p.Exec(t[1])
		if err != nil {
			return None, err
		}
		if x.Kind == Int && y.Kind == Int {
			return Token{Int, x.Text.(int64) % y.Text.(int64)}, nil
		}
		return None, ErrFitType
	})
	Global.Add(">", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 2 {
			return None, ErrParaNum
		}
		x, err := p.Exec(t[0])
		if err != nil {
			return None, err
		}
		y, err := p.Exec(t[1])
		if err != nil {
			return None, err
		}
		switch x.Kind {
		case Fold, Back, Front:
			return None, ErrFitType
		default:
			switch y.Kind {
			case List, Back, Front:
				return None, ErrFitType
			default:
				if z := x.Cmp(&y); z > 0 {
					return True, nil
				} else {
					return None, nil
				}
			}
		}
	})
	Global.Add(">=", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 2 {
			return None, ErrParaNum
		}
		x, err := p.Exec(t[0])
		if err != nil {
			return None, err
		}
		y, err := p.Exec(t[1])
		if err != nil {
			return None, err
		}
		switch x.Kind {
		case Fold, Back, Front:
			return None, ErrFitType
		default:
			switch y.Kind {
			case Fold, Back, Front:
				return None, ErrFitType
			default:
				if z := x.Cmp(&y); z >= 0 {
					return True, nil
				} else {
					return None, nil
				}
			}
		}
	})
	Global.Add("<", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 2 {
			return None, ErrParaNum
		}
		x, err := p.Exec(t[0])
		if err != nil {
			return None, err
		}
		y, err := p.Exec(t[1])
		if err != nil {
			return None, err
		}
		switch x.Kind {
		case Fold, Back, Front:
			return None, ErrFitType
		default:
			switch y.Kind {
			case Fold, Back, Front:
				return None, ErrFitType
			default:
				if z := x.Cmp(&y); z < 0 {
					return True, nil
				} else {
					return None, nil
				}
			}
		}
	})
	Global.Add("<=", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 2 {
			return None, ErrParaNum
		}
		x, err := p.Exec(t[0])
		if err != nil {
			return None, err
		}
		y, err := p.Exec(t[1])
		if err != nil {
			return None, err
		}
		switch x.Kind {
		case Fold, Back, Front:
			return None, ErrFitType
		default:
			switch y.Kind {
			case Fold, Back, Front:
				return None, ErrFitType
			default:
				if z := x.Cmp(&y); z <= 0 {
					return True, nil
				} else {
					return None, nil
				}
			}
		}
	})
	Global.Add("==", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 2 {
			return None, ErrParaNum
		}
		x, err := p.Exec(t[0])
		if err != nil {
			return None, err
		}
		y, err := p.Exec(t[1])
		if err != nil {
			return None, err
		}
		switch x.Kind {
		case Fold, Back, Front:
			return None, ErrFitType
		default:
			switch y.Kind {
			case Fold, Back, Front:
				return None, ErrFitType
			default:
				if z := x.Cmp(&y); z == 0 {
					return True, nil
				} else {
					return None, nil
				}
			}
		}
	})
	Global.Add("!=", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 2 {
			return None, ErrParaNum
		}
		x, err := p.Exec(t[0])
		if err != nil {
			return None, err
		}
		y, err := p.Exec(t[1])
		if err != nil {
			return None, err
		}
		switch x.Kind {
		case Fold, Back, Front:
			return None, ErrFitType
		default:
			switch y.Kind {
			case Fold, Back, Front:
				return None, ErrFitType
			default:
				if z := x.Cmp(&y); z != 0 {
					return True, nil
				} else {
					return None, nil
				}
			}
		}
	})
	Global.Add("and", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 2 {
			return None, ErrParaNum
		}
		x, err := p.Exec(t[0])
		if err != nil {
			return None, err
		}
		if !x.Bool() {
			return None, nil
		}
		y, err := p.Exec(t[1])
		if err != nil {
			return None, err
		}
		if y.Bool() {
			return True, nil
		} else {
			return None, nil
		}
	})
	Global.Add("or", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 2 {
			return None, ErrParaNum
		}
		x, err := p.Exec(t[0])
		if err != nil {
			return None, err
		}
		if x.Bool() {
			return True, nil
		}
		y, err := p.Exec(t[1])
		if err != nil {
			return None, err
		}
		if y.Bool() {
			return True, nil
		} else {
			return None, nil
		}
	})
	Global.Add("xor", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 2 {
			return None, ErrParaNum
		}
		x, err := p.Exec(t[0])
		if err != nil {
			return None, err
		}
		y, err := p.Exec(t[1])
		if err != nil {
			return None, err
		}
		if x.Bool() != y.Bool() {
			return True, nil
		} else {
			return None, nil
		}
	})
	Global.Add("not", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 1 {
			return None, ErrParaNum
		}
		x, err := p.Exec(t[0])
		if err != nil {
			return None, err
		}
		if x.Bool() {
			return None, nil
		} else {
			return True, nil
		}
	})
}
