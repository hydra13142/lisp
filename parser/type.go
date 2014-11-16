package parser

import "fmt"

type Pattern struct {
	rule []func([]byte) (interface{}, int)
	str  []string
}

type Scanner struct {
	ptn *Pattern
	tkn []byte
	skp bool
}

func (p *Pattern) Add(f func([]byte) (interface{}, int)) {
	p.rule = append(p.rule, f)
}

func (p *Pattern) AddString(s string) {
	p.str = append(p.str, s)
}

func (p *Pattern) NewScanner(s []byte, t bool) *Scanner {
	return &Scanner{ptn: p, tkn: s, skp: t}
}

func (s *Scanner) Skip() {
	i, l := 0, len(s.tkn)
	if l == 0 {
		return
	}
	for i < l && IsSpace(s.tkn[i]) {
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
			return t, -(i + 1), nil
		}
	}
	for i, f := range s.ptn.rule {
		a, l := f(s.tkn)
		if l > 0 {
			s.tkn = s.tkn[l:]
			return a, +(i + 1), nil
		}
	}
	return nil, 0, fmt.Errorf("unrecognised")
}

func (s *Scanner) Over() bool {
	return len(s.tkn) == 0
}
