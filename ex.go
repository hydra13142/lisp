package lisp

func (l *Lisp) EX() {
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
			if err != nil {
				return None, err
			}
		}
		return b, err
	})
	l.Add("default", func(t []Token, p *Lisp) (Token, error) {
		var x, y, z Token
		var err error
		if t[0].Kind != Label {
			return None, ErrFitType
		}
		x, err = p.Exec(t[0])
		if err != nil {
			return None, err
		}
		if x.Kind != Front {
			return None, ErrFitType
		}
		f := x.Text.(Lfac)
		n := f.Para
		if len(n) < len(t)-1 {
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
		return Token{Back, Gfac(func(t2 []Token, p2 *Lisp) (Token, error) {
			q := &Lisp{dad: p2, env: map[Name]Token{}}
			if len(t2) > len(n) || len(t2)+len(hold) < len(n) {
				return None, ErrParaNum
			}
			for m, n := range f.Make {
				q.env[m] = n
			}
			var i, j int
			for i, z = range t2 {
				y, err = p.Exec(z)
				if err != nil {
					return None, err
				}
				q.env[n[i]] = y
			}
			for i, j = len(n)-1, 1; i >= len(t2); i, j = i-1, j+1 {
				q.env[n[i]] = hold[len(hold)-j]
			}
			return q.Exec(Token{List, x.Text.(Lfac).Text})
		})}, nil
	})
	l.Add("omission", func(t []Token, p *Lisp) (Token, error) {
		var x, y Token
		var err error
		if len(t) != 1 {
			return None, ErrParaNum
		}
		if t[0].Kind != Label {
			return None, ErrFitType
		}
		x, err = p.Exec(t[0])
		if err != nil {
			return None, err
		}
		if x.Kind != Front {
			return None, ErrFitType
		}
		f := x.Text.(Lfac)
		n := f.Para
		return Token{Back, Gfac(func(t2 []Token, p2 *Lisp) (Token, error) {
			q := &Lisp{dad: p2, env: map[Name]Token{}}
			if len(t2) < len(n)-1 {
				return None, ErrParaNum
			}
			for m, n := range f.Make {
				q.env[m] = n
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
			return q.Exec(Token{List, x.Text.(Lfac).Text})
		})}, nil
	})
}
