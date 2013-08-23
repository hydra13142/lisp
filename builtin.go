package lisp

func init() {
	Add("none", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 0 {
			return None, ErrParaNum
		}
		return Token{}, nil
	})
	Add("eval", func(t []Token, p *Lisp) (ans Token, err error) {
		if len(t) != 1 {
			return None, ErrParaNum
		}
		ans = t[0]
		if ans.Kind == Label {
			ans, err = p.Exec(ans)
			if err != nil {
				return None, err
			}
		}
		return p.Exec(ans)
	})
	Add("quote", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 1 {
			return None, ErrParaNum
		}
		if t[0].Kind == Label {
			return p.Exec(t[0])
		}
		return t[0], nil
	})
	Add("atom", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 1 {
			return None, ErrParaNum
		}
		x, err := p.Exec(t[0])
		if err != nil {
			return None, err
		}
		if x.Kind != List || len(x.Text.([]Token)) == 0 {
			return True, nil
		} else {
			return None, nil
		}
	})
	Add("eq", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 2 {
			return None, ErrParaNum
		}
		x, err := p.Exec(t[0])
		if err != nil {
			return None, err
		}
		y, err := p.Exec(t[1])
		if err != nil {
			return None, err
		}
		if x.Eq(&y) {
			return True, nil
		} else {
			return None, nil
		}
	})
	Add("car", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 1 {
			return None, ErrParaNum
		}
		x, err := p.Exec(t[0])
		if err != nil {
			return None, err
		}
		if x.Kind == List {
			if len(x.Text.([]Token)) != 0 {
				return x.Text.([]Token)[0], nil
			}
			return None, ErrIsEmpty
		}
		return None, ErrFitType
	})
	Add("cdr", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 1 {
			return None, ErrParaNum
		}
		x, err := p.Exec(t[0])
		if err != nil {
			return None, err
		}
		if x.Kind == List {
			if len(x.Text.([]Token)) != 0 {
				return Token{List, x.Text.([]Token)[1:]}, nil
			}
			return None, ErrIsEmpty
		}
		return None, ErrFitType
	})
	Add("cons", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 2 {
			return None, ErrParaNum
		}
		x, err := p.Exec(t[0])
		if err != nil {
			return None, err
		}
		y, err := p.Exec(t[1])
		if err != nil {
			return None, err
		}
		if y.Kind == List {
			a := y.Text.([]Token)
			b := make([]Token, len(a)+1)
			b[0] = x
			copy(b[1:], a)
			return Token{List, b}, nil
		}
		return None, ErrFitType
	})
	Add("cond", func(t []Token, p *Lisp) (ans Token, err error) {
		if len(t) == 0 {
			return None, ErrParaNum
		}
		for _, i := range t {
			if i.Kind == List {
				t := i.Text.([]Token)
				if len(t) == 2 {
					ans, err = p.Exec(t[0])
					if err != nil {
						return None, err
					}
					if ans.Bool() {
						return p.Exec(t[1])
					}
					continue
				}
				return None, ErrParaNum
			}
			return None, ErrFitType
		}
		return None, nil
	})
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
}
