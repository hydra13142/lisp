package lisp

import (
	"bytes"
	"io"
	"io/ioutil"
	"os"
)

type Lisp struct {
	dad *Lisp
	env map[Name]Token
}

func NewLisp() *Lisp {
	x := new(Lisp)
	x.env = map[Name]Token{}
	x.dad = Global
	return x
}

func Add(s string, f func([]Token, *Lisp) (Token, error)) {
	Global.env[Name(s)] = Token{Back, Gfac(f)}
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
	case Label:
		for v := l; ; {
			ct, ok = v.env[f.Text.(Name)]
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
		case Label:
			for v := l; ; {
				ct, ok = v.env[ls[0].Text.(Name)]
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
			return ct.Text.(Gfac)(ls[1:], l)
		case Macro:
			mp := ct.Text.(Macr)
			if len(ls) != len(mp.Para)+1 {
				return None, ErrParaNum
			}
			xp := map[Name]Token{}
			for i, t := range ls[1:] {
				xp[mp.Para[i]] = t
			}
			return l.Exec(Repl(Token{List, mp.Text}, xp))
		case Front:
			lp := ct.Text.(Lfac)
			q := &Lisp{dad: l, env: map[Name]Token{}}
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
			ans, err = q.Exec(Token{Text: lp.Text, Kind: List})
			if err != nil {
				return None, err
			}
			for i, _ := range lp.Make {
				lp.Make[i] = q.env[i]
			}
			return ans, nil
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

func (l *Lisp) Load(s string) (Token, error) {
	var file *os.File
	var data []byte
	var err error
	file, err = os.Open(s)
	if err != nil {
		file, err = os.Open(s + ".lsp")
		if err != nil {
			return None, err
		}
	}
	defer file.Close()
	data, err = ioutil.ReadAll(file)
	if err != nil {
		return None, err
	}
	buf := bytes.NewBuffer(data)
	one := block{}
	for {
		data, err := buf.ReadBytes('\n')
		if err != nil {
			if err != io.EOF {
				return None, err
			}
			err = one.feed(data)
			break
		}
		err = one.feed(data)
		if err != nil {
			return None, err
		}
	}
	if !one.over() {
		return None, ErrUnquote
	}
	return l.Eval(one.total)
}
