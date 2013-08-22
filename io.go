package lisp

import (
	"bufio"
	"fmt"
	"os"
)

type block struct {
	quote bool
	count int
	total string
}

func (b *block) feed(s []byte) error {
	single := false
	for i, l := 0, len(s); i < l; i++ {
		if b.quote {
			switch s[i] {
			case '"':
				b.quote = false
			case '\\':
				i++
			}
		} else if single {
			switch s[i] {
			case '\'':
				single = false
			case '\\':
				i++
			}
		} else {
			switch s[i] {
			case '(':
				b.count++
			case ')':
				b.count--
			case '\'':
				if i+1 < len(s) {
					if s[i+1] == '(' && (i+2 >= len(s) || s[i+2] != '\'') {
						b.count++
						i++
					} else {
						single = true
					}
				}
			case '"':
				b.quote = true
			}
		}
	}
	if single || b.count < 0 {
		return ErrUnquote
	}
	b.total += string(s)
	return nil
}

func (b *block) over() bool {
	return b.count == 0 && b.quote == false
}

func (l *Lisp) IO() {
	l.Add("scan", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 0 {
			return None, ErrParaNum
		}
		buf := bufio.NewReader(os.Stdin)
		one := block{}
		for {
			data, err := buf.ReadBytes('\n')
			if err != nil {
				return None, err
			}
			err = one.feed(data)
			if err != nil {
				return None, err
			}
			if one.over() {
				break
			}
		}
		return p.Eval(one.total)
	})
	l.Add("load", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 1 {
			return None, ErrParaNum
		}
		if t[0].Kind != String {
			return None, ErrFitType
		}
		return p.Load(t[0].Text.(string))
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
