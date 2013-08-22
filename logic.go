package lisp

func init() {
	Add("and", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 2 {
			return None, ErrParaNum
		}
		x, err := p.Exec(t[0])
		if err != nil {
			return None, err
		}
		if !x.Bool() {
			return None, nil
		}
		y, err := p.Exec(t[1])
		if err != nil {
			return None, err
		}
		if y.Bool() {
			return True, nil
		} else {
			return None, nil
		}
	})
	Add("or", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 2 {
			return None, ErrParaNum
		}
		x, err := p.Exec(t[0])
		if err != nil {
			return None, err
		}
		if x.Bool() {
			return True, nil
		}
		y, err := p.Exec(t[1])
		if err != nil {
			return None, err
		}
		if y.Bool() {
			return True, nil
		} else {
			return None, nil
		}
	})
	Add("xor", func(t []Token, p *Lisp) (Token, error) {
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
		if x.Bool() != y.Bool() {
			return True, nil
		} else {
			return None, nil
		}
	})
	Add("not", func(t []Token, p *Lisp) (Token, error) {
		if len(t) != 1 {
			return None, ErrParaNum
		}
		x, err := p.Exec(t[0])
		if err != nil {
			return None, err
		}
		if x.Bool() {
			return None, nil
		} else {
			return True, nil
		}
	})
}
