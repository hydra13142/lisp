package lisp

func init() {
	Add("chan", func(t []Token, p *Lisp) (ans Token, err error) {
		switch len(t) {
		case 0:
			return Token{Chan, make(chan Token)}, nil
		case 1:
			u, err := p.Exec(t[0])
			if err != nil {
				return None, err
			}
			if u.Kind != Int {
				return None, ErrFitType
			}
			return Token{Chan, make(chan Token, int(u.Text.(int64)))}, nil
		}
		return None, ErrParaNum
	})
	Add("close", func(t []Token, p *Lisp) (ans Token, err error) {
		if len(t) != 1 {
			return None, ErrParaNum
		}
		u, err := p.Exec(t[0])
		if err != nil {
			return None, err
		}
		if u.Kind != Chan {
			return None, ErrFitType
		}
		close(u.Text.(chan Token))
		return None, nil
	})
	Add("go", func(t []Token, p *Lisp) (ans Token, err error) {
		if len(t) != 1 {
			return None, ErrParaNum
		}
		go p.Exec(t[0])
		return None, nil
	})
}
