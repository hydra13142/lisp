package parser

import (
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
