package lisp

func init() {
	Add("builtin", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 1 {
			return None, ErrParaNum
		}
		if t[0].Kind != Label {
			return None, ErrFitType
		}
		ans, ok := Global.env[t[0].Text.(Name)]
		if !ok {
			return None, ErrNotFind
		}
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
		case List:
			if b.Kind != List {
				return None, ErrFitType
			}
			t = a.Text.([]Token)
			if len(t) <= 0 {
				return None, ErrParaNum
			}
			x := make([]Name, len(t))
			for i, c := range t {
				if c.Kind != Label {
					return None, ErrNotName
				}
				x[i] = c.Text.(Name)
			}
			ans = Token{Front, &Lfac{x[1:], b.Text.([]Token), p}}
			p.env[x[0]] = ans
			return ans, nil
		}
		return None, ErrFitType
	})
	Add("update", func(t []Token, p *Lisp) (ans Token, err error) {
		if len(t) != 2 {
			return None, ErrParaNum
		}
		a, b := t[0], t[1]
		var n Name
		switch a.Kind {
		case Label:
			n = a.Text.(Name)
		case List:
			if b.Kind != List {
				return None, ErrFitType
			}
			t = a.Text.([]Token)
			if len(t) <= 0 {
				return None, ErrParaNum
			}
			n = t[0].Text.(Name)
		default:
			return None, ErrFitType
		}
		for v := p; p != Global; p = p.dad {
			_, ok := p.env[n]
			if ok {
				if a.Kind == Label {
					ans, err = p.Exec(b)
					if err == nil {
						p.env[n] = ans
					}
					return ans, err
				} else {
					x := make([]Name, len(t)-1)
					for i, c := range t[1:] {
						if c.Kind != Label {
							return None, ErrNotName
						}
						x[i] = c.Text.(Name)
					}
					ans = Token{Front, &Lfac{x, b.Text.([]Token), p}}
					v.env[n] = ans
					return ans, nil
				}
			}
		}
		_, ok := p.env[n]
		if !ok {
			return None, ErrNotFind
		}
		return None, ErrRefused
	})
	Add("remove", func(t []Token, p *Lisp) (ans Token, err error) {
		if len(t) != 1 {
			return None, ErrParaNum
		}
		if t[0].Kind != Label {
			return None, ErrFitType
		}
		n := t[0].Text.(Name)
		for ; p != Global; p = p.dad {
			_, ok := p.env[n]
			if ok {
				delete(p.env, n)
				return None, nil
			}
		}
		_, ok := p.env[n]
		if !ok {
			return None, ErrNotFind
		}
		return None, ErrRefused
	})
	Add("present", func(t []Token, p *Lisp) (ans Token, err error) {
		if len(t) != 0 {
			return None, ErrParaNum
		}
		x := make([]Token, 0, len(p.env))
		for i, _ := range p.env {
			x = append(x, Token{Label, i})
		}
		return Token{List, x}, nil
	})
	Add("context", func(t []Token, p *Lisp) (ans Token, err error) {
		if len(t) != 0 {
			return None, ErrParaNum
		}
		x := make([]Token, 0, 128)
		for v := p; v != nil; v = v.dad {
			for i, _ := range v.env {
				x = append(x, Token{Label, i})
			}
		}
		return Token{List, x}, nil
	})
	Add("clear", func(t []Token, p *Lisp) (ans Token, err error) {
		if len(t) != 0 {
			return None, ErrParaNum
		}
		p.env = map[Name]Token{}
		return None, nil
	})
}
