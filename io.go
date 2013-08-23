package lisp

import (
	"bufio"
	"fmt"
	"os"
)

func init() {
	Add("scan", func(t []Token, p *Lisp) (Token, error) {
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
	Add("load", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 1 {
			return None, ErrParaNum
		}
		if t[0].Kind != String {
			return None, ErrFitType
		}
		return p.Load(t[0].Text.(string))
	})
	Add("print", func(t []Token, p *Lisp) (x Token, y error) {
		for _, i := range t {
			x, y = p.Exec(i)
			if y != nil {
				return None, y
			}
			fmt.Print(x)
		}
		return x, nil
	})
	Add("println", func(t []Token, p *Lisp) (x Token, y error) {
		for _, i := range t {
			x, y = p.Exec(i)
			if y != nil {
				return None, y
			}
			fmt.Println(x)
		}
		return x, nil
	})
}
