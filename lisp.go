package lisp

import (
	"errors"
	"fmt"
	"math"
)

type Pattern struct {
	rule []func([]byte) (interface{}, int)
	str  []string
}

type Scanner struct {
	ptn *Pattern
	tkn []byte
	skp bool
}

type kind int

type null struct{}

type name string

type pres func([]Token, *Lisp) (Token, error)

type afts struct {
	para []name
	text []Token
}

type Token struct {
	kind
	text interface{}
}

type Lisp struct {
	dad *Lisp
	env map[name]Token
}

const (
	Null kind = iota
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
	ErrUnmatch = errors.New("Left quote != right quote")
	ErrNotFind = errors.New("Not find this name")
	ErrNotFunc = errors.New("Not a function")
	ErrParaNum = errors.New("Wrong parament number")
	ErrFitType = errors.New("Lisp type is wrong")
	ErrNotName = errors.New("This's not a name")
	ErrIsEmpty = errors.New("Fold is empty")
)

func IsSpace(c byte) bool {
	return c == '\r' || c == '\n' || c == '\t' || c == '\v' || c == '\f' || c == ' '
}

func IsDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func IsLower(c byte) bool {
	return c >= 'a' && c <= 'z'
}

func IsUpper(c byte) bool {
	return c >= 'A' && c <= 'Z'
}

func IsAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func IsFirst(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

func IsLetter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '_'
}

func ParseUint(r []byte) (i int64, l int) {
	var c byte
	for l, c = range r {
		if c < '0' || c > '9' {
			break
		}
		i = i*10 + int64(c-'0')
	}
	if c >= '0' && c <= '9' {
		l++
	}
	return
}

func ParseInt(r []byte) (i int64, l int) {
	if len(r) == 0 {
		return
	}
	c, t := false, 0
	switch r[0] {
	case '-':
		c, t = true, 1
	case '+':
		t = 1
	}
	a, j := ParseUint(r[t:])
	if j > 0 {
		if c {
			i = -a
		} else {
			i = a
		}
		l = j + t
	}
	return
}

func ParseDecimal(r []byte) (f float64, l int) {
	var c byte
	if len(r) == 0 {
		return
	}
	if r[0] == '.' {
		i, j := 0, 1
		for l, c = range r[1:] {
			if c < '0' || c > '9' {
				break
			}
			i, j = i*10+int(c-'0'), j*10
		}
		if c >= '0' && c <= '9' {
			l++
		}
		if l > 0 {
			f, l = float64(i)/float64(j), l+1
		}
	}
	return
}

func ParseFinger(r []byte) (i int64, l int) {
	if len(r) == 0 {
		return
	}
	if r[0] == 'e' || r[0] == 'E' {
		a, j := ParseInt(r[1:])
		if j > 0 {
			i, l = a, j+1
		}
	}
	return
}

func ParseFloat(r []byte) (f float64, l int) {
	if len(r) == 0 {
		return
	}
	p, t := false, 0
	switch r[0] {
	case '-':
		p, t = true, 1
	case '+':
		t = 1
	}
	a, i := ParseUint(r[t:])
	b, j := ParseDecimal(r[t+i:])
	c, k := ParseFinger(r[t+i+j:])
	if i > 0 || j > 0 {
		f = float64(a) + b
		if k > 0 {
			f *= math.Pow10(int(c))
		}
		if p {
			f = -f
		}
		l = t + i + j + k
	}
	return
}

func ParseChar(r []byte) (c rune, l int) {
	if len(r) == 0 {
		return
	}
	if r[0] == '\'' {
		_, err := fmt.Sscanf(string(r), "'%c'", &c)
		if err == nil {
			for l = 1; r[l] != '\''; l++ {
				if r[l] == '\\' {
					l++
				}
			}
			l++
		} else {
			c = 0
		}
	}
	return
}

func ParseString(r []byte) (s string, l int) {
	if len(r) == 0 {
		return
	}
	if r[0] == '"' {
		i := 1
		for ; i < len(r) && r[i] != '"'; i++ {
			if r[i] == '\\' {
				i++
			}
		}
		if i < len(r) {
			s = string(r[1:i])
			l = i + 1
		}
	}
	return
}

func ParseName(r []byte) (s string, l int) {
	if len(r) == 0 {
		return
	}
	isf := func(c byte) bool {
		return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
	}
	isw := func(c byte) bool {
		return (c >= '0' && c <= '9') || isf(c)
	}
	if isf(r[0]) {
		for l = 1; l < len(r) && isw(r[l]); l++ {
		}
		s = string(r[:l])
	}
	return
}

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
	if tkn[0].kind == Operator {
		var t bool
		switch tkn[0].text.(byte) {
		case '(':
			t = true
		case '[':
			t = false
		default:
			return nil, ErrUnquote
		}
		i, j, l := 1, 1, len(tkn)
		for i < l && j > 0 {
			if tkn[i].kind == Operator {
				switch tkn[i].text.(byte) {
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
				f = Token{text: fold, kind: List}
			} else {
				f = Token{text: fold, kind: Fold}
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

func (p *Pattern) Add(f func([]byte) (interface{}, int)) {
	p.rule = append(p.rule, f)
}

func (p *Pattern) AddString(s string) {
	p.str = append(p.str, s)
}

func (p *Pattern) NewScanner(s string, t bool) *Scanner {
	return &Scanner{ptn: p, tkn: []byte(s), skp: t}
}

func (s *Scanner) Skip() {
	if len(s.tkn) == 0 {
		return
	}
	i := 0
	for i < len(s.tkn) && IsSpace(s.tkn[i]) {
		i++
	}
	s.tkn = s.tkn[i:]
}

func (s *Scanner) Scan() (interface{}, int, error) {
	if s.skp {
		s.Skip()
	}
	if len(s.tkn) == 0 {
		return nil, 0, fmt.Errorf("empty string")
	}
	for i, t := range s.ptn.str {
		l := len(t)
		if len(s.tkn) >= l && t == string(s.tkn[:l]) {
			s.tkn = s.tkn[l:]
			return t, -i - 1, nil
		}
	}
	for i, f := range s.ptn.rule {
		a, l := f(s.tkn)
		if l > 0 {
			s.tkn = s.tkn[l:]
			return a, +i + 1, nil
		}
	}
	return nil, 0, fmt.Errorf("unrecognised")
}

func (s *Scanner) Over() bool {
	return len(s.tkn) == 0
}

func (t kind) String() string {
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
	return fmt.Sprintf("{front : (%v,%v)}", l.para, l.text)
}

func (t Token) String() string {
	return fmt.Sprint(t.text)
}

func (t *Token) Bool() bool {
	switch t.kind {
	case Int:
		return t.text.(int64) != 0
	case Float:
		return t.text.(float64) != 0
	case String:
		return t.text.(string) != ""
	case List:
		return len(t.text.([]Token)) != 0
	case Null:
		return false
	}
	return true
}

func (t *Token) Eq(p *Token) bool {
	if t.kind != p.kind {
		return false
	}
	switch t.kind {
	case Int:
		return t.text.(int64) == p.text.(int64)
	case Float:
		return t.text.(float64) == p.text.(float64)
	case String:
		return t.text.(string) == p.text.(string)
	case Fold, List:
		a, b := t.text.([]Token), p.text.([]Token)
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
	switch t.kind {
	case Int:
		switch p.kind {
		case Int:
			a = t.text.(int64) > p.text.(int64)
			b = t.text.(int64) < p.text.(int64)
		case Float:
			a = float64(t.text.(int64)) > p.text.(float64)
			b = float64(t.text.(int64)) < p.text.(float64)
		default:
			return 0
		}
	case Float:
		switch p.kind {
		case Int:
			a = t.text.(float64) > float64(p.text.(int64))
			b = t.text.(float64) < float64(p.text.(int64))
		case Float:
			a = t.text.(float64) > p.text.(float64)
			b = t.text.(float64) < p.text.(float64)
		default:
			return 0
		}
	case String:
		switch p.kind {
		case Int, Float:
			return 1
		case String:
			a = t.text.(string) > p.text.(string)
			b = t.text.(string) < p.text.(string)
		default:
			return 0
		}
	case List:
		switch p.kind {
		case Int, Float, String:
			return 1
		case List:
			x, y := t.text.([]Token), p.text.([]Token)
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
	switch f.kind {
	case Fold:
		return Token{List, f.text.([]Token)}, nil
	case Name:
		for v := l; ; {
			ct, ok = v.env[f.text.(name)]
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
		ls = f.text.([]Token)
		if len(ls) == 0 {
			return None, nil
		}
		ct = ls[0]
		switch ct.kind {
		case Name:
			for v := l; ; {
				ct, ok = v.env[ls[0].text.(name)]
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
		switch ct.kind {
		case Back:
			return ct.text.(pres)(ls[1:], l)
		case Front:
			lp := ct.text.(afts)
			q := &Lisp{dad: l, env: map[name]Token{}}
			if len(ls) != len(lp.para)+1 {
				return None, ErrParaNum
			}
			for i, t := range ls[1:] {
				q.env[lp.para[i]], err = l.Exec(t)
				if err != nil {
					return None, err
				}
			}
			return q.Exec(Token{text: lp.text, kind: List})
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

func (l *Lisp) Extend() {
	l.Add("if", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 3 {
			return None, ErrParaNum
		}
		ans, err := p.Exec(t[0])
		if err != nil {
			return None, err
		}
		if ans.Bool() {
			return p.Exec(t[1])
		} else {
			return p.Exec(t[2])
		}
	})
	l.Add("loop", func(t []Token, p *Lisp) (Token, error) {
		var a, b Token
		var err error
		if len(t) != 3 {
			return None, ErrParaNum
		}
		_, err = p.Exec(t[0])
		if err != nil {
			return None, err
		}
		for {
			a, err = p.Exec(t[1])
			if err != nil {
				return None, err
			}
			if !a.Bool() {
				break
			}
			b, err = p.Exec(t[2])
		}
		return b, err
	})
	l.Add("default", func(t []Token, p *Lisp) (Token, error) {
		var x, y, z Token
		var err error
		if t[0].kind != Name {
			return None, ErrFitType
		}
		x, err = p.Exec(t[0])
		if err != nil {
			return None, err
		}
		if x.kind != Front {
			return None, ErrFitType
		}
		n := x.text.(afts).para
		if len(n)+1 != len(t) {
			return None, ErrParaNum
		}
		hold := make([]Token, 0, len(t)-1)
		for _, z = range t[1:] {
			y, err = p.Exec(z)
			if err != nil {
				return None, err
			}
			hold = append(hold, y)
		}
		return Token{Back, pres(func(t2 []Token, p2 *Lisp) (Token, error) {
			q := &Lisp{dad: p2, env: map[name]Token{}}
			if len(t2) > len(n) {
				return None, ErrParaNum
			}
			var i int
			for i, z = range t2 {
				y, err = p.Exec(z)
				if err != nil {
					return None, err
				}
				q.env[n[i]] = y
			}
			for i = len(t2); i < len(n); i++ {
				q.env[n[i]] = hold[i]
			}
			return q.Exec(Token{List, x.text.(afts).text})
		})}, nil
	})
	l.Add("omission", func(t []Token, p *Lisp) (Token, error) {
		var x, y Token
		var err error
		if len(t) != 1 {
			return None, ErrParaNum
		}
		if t[0].kind != Name {
			return None, ErrFitType
		}
		x, err = p.Exec(t[0])
		if err != nil {
			return None, err
		}
		if x.kind != Front {
			return None, ErrFitType
		}
		n := x.text.(afts).para
		return Token{Back, pres(func(t2 []Token, p2 *Lisp) (Token, error) {
			q := &Lisp{dad: p2, env: map[name]Token{}}
			if len(t2) < len(n)-1 {
				return None, ErrParaNum
			}
			var i int
			for i = len(n) - 2; i >= 0; i-- {
				y, err = p.Exec(t2[i])
				if err != nil {
					return None, err
				}
				q.env[n[i]] = y
			}
			z := make([]Token, 0, len(t2)-len(n)+1)
			for i = len(n) - 1; i < len(t2); i++ {
				y, err = p.Exec(t2[i])
				if err != nil {
					return None, err
				}
				z = append(z, y)
			}
			q.env[n[len(n)-1]] = Token{List, z}
			return q.Exec(Token{List, x.text.(afts).text})
		})}, nil
	})
	l.Add("print", func(t []Token, p *Lisp) (Token, error) {
		for _, i := range t {
			x, y := p.Exec(i)
			if y != nil {
				return None, y
			}
			fmt.Print(x)
		}
		return None, nil
	})
	l.Add("println", func(t []Token, p *Lisp) (Token, error) {
		for _, i := range t {
			x, y := p.Exec(i)
			if y != nil {
				return None, y
			}
			fmt.Println(x)
		}
		return None, nil
	})
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
		if t[0].kind == Name {
			return p.Exec(t[0])
		}
		return t[0], nil
	})
	Global.Add("run", func(t []Token, p *Lisp) (ans Token, err error) {
		if len(t) != 1 {
			return None, ErrParaNum
		}
		ans = t[0]
		if ans.kind == Name {
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
		if x.kind != List || len(x.text.([]Token)) == 0 {
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
		if x.kind == List {
			if len(x.text.([]Token)) != 0 {
				return x.text.([]Token)[0], nil
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
		if x.kind == List {
			if len(x.text.([]Token)) != 0 {
				return Token{List, x.text.([]Token)[1:]}, nil
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
		if y.kind == List {
			a := y.text.([]Token)
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
			if i.kind == List {
				t := i.text.([]Token)
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
		if a.kind != List {
			return None, ErrFitType
		}
		if b.kind != List {
			return None, ErrFitType
		}
		t = a.text.([]Token)
		x := make([]name, 0, len(t))
		for _, i := range t {
			if i.kind != Name {
				return None, ErrNotName
			}
			x = append(x, i.text.(name))
		}
		return Token{Front, afts{x, b.text.([]Token)}}, nil
	})
	Global.Add("define", func(t []Token, p *Lisp) (ans Token, err error) {
		if len(t) != 2 {
			return None, ErrParaNum
		}
		a, b := t[0], t[1]
		switch a.kind {
		case Name:
			ans, err = p.Exec(b)
			if err == nil {
				p.env[a.text.(name)] = ans
			}
			return ans, err
		case List:
			if b.kind != List {
				return None, ErrFitType
			}
			t = a.text.([]Token)
			x := make([]name, 0, len(t))
			for _, i := range t {
				if i.kind != Name {
					return None, ErrNotName
				}
				x = append(x, i.text.(name))
			}
			ans = Token{Front, afts{x[1:], b.text.([]Token)}}
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
		switch x.kind {
		case Int:
			switch y.kind {
			case Int:
				return Token{Int, x.text.(int64) + y.text.(int64)}, nil
			case Float:
				return Token{Float, float64(x.text.(int64)) + y.text.(float64)}, nil
			}
		case Float:
			switch y.kind {
			case Int:
				return Token{Float, x.text.(float64) + float64(y.text.(int64))}, nil
			case Float:
				return Token{Float, x.text.(float64) + y.text.(float64)}, nil
			}
		case String:
			switch y.kind {
			case String:
				return Token{String, x.text.(string) + y.text.(string)}, nil
			}
		case List:
			switch y.kind {
			case List:
				a, b := x.text.([]Token), y.text.([]Token)
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
		switch x.kind {
		case Int:
			switch y.kind {
			case Int:
				return Token{Int, x.text.(int64) - y.text.(int64)}, nil
			case Float:
				return Token{Float, float64(x.text.(int64)) - y.text.(float64)}, nil
			}
		case Float:
			switch y.kind {
			case Int:
				return Token{Float, x.text.(float64) - float64(y.text.(int64))}, nil
			case Float:
				return Token{Float, x.text.(float64) - y.text.(float64)}, nil
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
		switch x.kind {
		case Int:
			switch y.kind {
			case Int:
				return Token{Int, x.text.(int64) * y.text.(int64)}, nil
			case Float:
				return Token{Float, float64(x.text.(int64)) * y.text.(float64)}, nil
			}
		case Float:
			switch y.kind {
			case Int:
				return Token{Float, x.text.(float64) * float64(y.text.(int64))}, nil
			case Float:
				return Token{Float, x.text.(float64) * y.text.(float64)}, nil
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
		switch x.kind {
		case Int:
			switch y.kind {
			case Int:
				return Token{Int, x.text.(int64) / y.text.(int64)}, nil
			case Float:
				return Token{Float, float64(x.text.(int64)) / y.text.(float64)}, nil
			}
		case Float:
			switch y.kind {
			case Int:
				return Token{Float, x.text.(float64) / float64(y.text.(int64))}, nil
			case Float:
				return Token{Float, x.text.(float64) / y.text.(float64)}, nil
			}
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
		switch x.kind {
		case Fold, Back, Front:
			return None, ErrFitType
		default:
			switch y.kind {
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
		switch x.kind {
		case Fold, Back, Front:
			return None, ErrFitType
		default:
			switch y.kind {
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
		switch x.kind {
		case Fold, Back, Front:
			return None, ErrFitType
		default:
			switch y.kind {
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
		switch x.kind {
		case Fold, Back, Front:
			return None, ErrFitType
		default:
			switch y.kind {
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
		switch x.kind {
		case Fold, Back, Front:
			return None, ErrFitType
		default:
			switch y.kind {
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
		switch x.kind {
		case Fold, Back, Front:
			return None, ErrFitType
		default:
			switch y.kind {
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
