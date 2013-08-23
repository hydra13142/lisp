package lisp

import "fmt"

func init() {
	Add("lambda", func(t []Token, p *Lisp) (ans Token, err error) {
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
		x := make([]Name, 0, len(t))
		for _, i := range t {
			if i.Kind != Label {
				return None, ErrNotName
			}
			x = append(x, i.Text.(Name))
		}
		u := make(map[Name]Token)
		for i, j := range p.env {
			u[i] = j
		}
		ans = Token{Front, Lfac{x, b.Text.([]Token), u}}
		u[Name("self")] = ans
		return ans, nil
	})
	Add("define", func(t []Token, p *Lisp) (ans Token, err error) {
		if len(t) != 2 {
			return None, ErrParaNum
		}
		a, b := t[0], t[1]
		switch a.Kind {
		case Label:
			ans, err = p.Exec(b)
			if err == nil {
				p.env[a.Text.(Name)] = ans
			}
			return ans, err
		case Fold:
			if b.Kind != List {
				return None, ErrFitType
			}
			t = a.Text.([]Token)
			x := make([]Name, len(t))
			for i, c := range t {
				if c.Kind != Label {
					return None, ErrNotName
				}
				x[i] = c.Text.(Name)
			}
			ans = Token{Macro, Macr{x[1:], b.Text.([]Token)}}
			p.env[x[0]] = ans
			return ans, nil
		case List:
			if b.Kind != List {
				return None, ErrFitType
			}
			t = a.Text.([]Token)
			x := make([]Name, len(t))
			for i, c := range t {
				if c.Kind != Label {
					return None, ErrNotName
				}
				x[i] = c.Text.(Name)
			}
			u := make(map[Name]Token)
			for i, j := range p.env {
				u[i] = j
			}
			ans = Token{Front, Lfac{x[1:], b.Text.([]Token), u}}
			p.env[x[0]] = ans
			u[x[0]] = ans
			return ans, nil
		}
		return None, ErrFitType
	})
	Add("remove", func(t []Token, p *Lisp) (ans Token, err error) {
		if len(t) != 1 {
			return None, ErrParaNum
		}
		if t[0].Kind != Label {
			return None, ErrFitType
		}
		n := t[0].Text.(Name)
		var v *Lisp
		var ok bool
		for v = p; ; {
			_, ok = v.env[n]
			if ok {
				break
			}
			v = v.dad
			if v == Global {
				break
			}
		}
		if !ok {
			_, ok = v.env[n]
			if !ok {
				return None, ErrNotFind
			}
			return None, ErrRefused
		}
		delete(v.env, n)
		return None, nil
	})
	Add("clear", func(t []Token, p *Lisp) (ans Token, err error) {
		if len(t) != 0 {
			return None, ErrParaNum
		}
		p.env = map[Name]Token{}
		return None, nil
	})
	Add("present", func(t []Token, p *Lisp) (ans Token, err error) {
		if len(t) != 0 {
			return None, ErrParaNum
		}
		for i, _ := range p.env {
			fmt.Println(string(i))
		}
		return None, nil
	})
	Add("context", func(t []Token, p *Lisp) (ans Token, err error) {
		if len(t) != 0 {
			return None, ErrParaNum
		}
		for v := p; v != nil; v = v.dad {
			for i, _ := range v.env {
				fmt.Println(string(i))
			}
		}
		return None, nil
	})
	Add("solid", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 1 {
			return None, ErrParaNum
		}
		ans, err := p.Exec(t[0])
		if err != nil {
			return None, err
		}
		return Hard(ans), nil
	})
}
