package lisp

func init() {
	Add("each", func(t []Token, p *Lisp) (ans Token, err error) {
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
	Add("block", func(t []Token, p *Lisp) (ans Token, err error) {
		if len(t) == 0 {
			return None, ErrParaNum
		}
		q := &Lisp{dad: p, env: map[Name]Token{}}
		for _, i := range t {
			ans, err = q.Exec(i)
			if err != nil {
				break
			}
		}
		return ans, err
	})
	Add("if", func(t []Token, p *Lisp) (Token, error) {
		if len(t) < 2 || len(t) > 3 {
			return None, ErrParaNum
		}
		ans, err := p.Exec(t[0])
		if err != nil {
			return None, err
		}
		if ans.Bool() {
			return p.Exec(t[1])
		}
		if len(t) == 3 {
			return p.Exec(t[2])
		}
		return None, nil
	})
	Add("cond", func(t []Token, p *Lisp) (Token, error) {
		if len(t) == 0 {
			return None, ErrParaNum
		}
		for _, i := range t {
			if i.Kind != List {
				return None, ErrFitType
			}
			t := i.Text.([]Token)
			if len(t) != 2 {
				return None, ErrParaNum
			}
			ans, err := p.Exec(t[0])
			if err != nil {
				return None, err
			}
			if ans.Bool() {
				return p.Exec(t[1])
			}
		}
		return None, nil
	})
	Add("while", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 2 {
			return None, ErrParaNum
		}
		for {
			a, err := p.Exec(t[0])
			if err != nil {
				return None, err
			}
			if !a.Bool() {
				break
			}
			_, err = p.Exec(t[1])
			if err != nil {
				return None, err
			}
		}
		return None, nil
	})
	Add("until", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 2 {
			return None, ErrParaNum
		}
		for {
			a, err := p.Exec(t[0])
			if err != nil {
				return None, err
			}
			if a.Bool() {
				break
			}
			_, err = p.Exec(t[1])
			if err != nil {
				return None, err
			}
		}
		return None, nil
	})
	Add("loop", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 3 {
			return None, ErrParaNum
		}
		_, err := p.Exec(t[0])
		if err != nil {
			return None, err
		}
		for {
			a, err := p.Exec(t[1])
			if err != nil {
				return None, err
			}
			if !a.Bool() {
				break
			}
			_, err = p.Exec(t[2])
			if err != nil {
				return None, err
			}
		}
		return None, nil
	})
	Add("for", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 3 {
			return None, ErrParaNum
		}
		if t[0].Kind != Label {
			return None, ErrFitType
		}
		iter, err := p.Exec(t[1])
		if err != nil {
			return None, err
		}
		if iter.Kind != Front || len(iter.Text.(*Lfac).Para) != 0 {
			return None, ErrFitType
		}
		g := Token{List, iter.Text.(*Lfac).Text}
		n := t[0].Text.(Name)
		q := &Lisp{dad: iter.Text.(*Lfac).Make, env: map[Name]Token{}}
		q.env[Name("self")] = iter
		for {
			a, err := q.Exec(g)
			if err != nil {
				break
			}
			p.env[n] = a
			_, err = p.Exec(t[2])
			if err != nil {
				return None, err
			}
		}
		return None, nil
	})
}
