package lisp

func init() {
	Add(">", func(t []Token, p *Lisp) (Token, error) {
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
		switch x.Kind {
		case Fold, Back, Front:
			return None, ErrFitType
		default:
			switch y.Kind {
			case List, Back, Front:
				return None, ErrFitType
			default:
				if z := x.Cmp(&y); z > 0 {
					return True, nil
				} else {
					return False, nil
				}
			}
		}
	})
	Add(">=", func(t []Token, p *Lisp) (Token, error) {
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
		switch x.Kind {
		case Fold, Back, Front:
			return None, ErrFitType
		default:
			switch y.Kind {
			case Fold, Back, Front:
				return None, ErrFitType
			default:
				if z := x.Cmp(&y); z >= 0 {
					return True, nil
				} else {
					return False, nil
				}
			}
		}
	})
	Add("<", func(t []Token, p *Lisp) (Token, error) {
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
		switch x.Kind {
		case Fold, Back, Front:
			return None, ErrFitType
		default:
			switch y.Kind {
			case Fold, Back, Front:
				return None, ErrFitType
			default:
				if z := x.Cmp(&y); z < 0 {
					return True, nil
				} else {
					return False, nil
				}
			}
		}
	})
	Add("<=", func(t []Token, p *Lisp) (Token, error) {
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
		switch x.Kind {
		case Fold, Back, Front:
			return None, ErrFitType
		default:
			switch y.Kind {
			case Fold, Back, Front:
				return None, ErrFitType
			default:
				if z := x.Cmp(&y); z <= 0 {
					return True, nil
				} else {
					return False, nil
				}
			}
		}
	})
	Add("==", func(t []Token, p *Lisp) (Token, error) {
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
		switch x.Kind {
		case Fold, Back, Front:
			return None, ErrFitType
		default:
			switch y.Kind {
			case Fold, Back, Front:
				return None, ErrFitType
			default:
				if z := x.Cmp(&y); z == 0 {
					return True, nil
				} else {
					return False, nil
				}
			}
		}
	})
	Add("!=", func(t []Token, p *Lisp) (Token, error) {
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
		switch x.Kind {
		case Fold, Back, Front:
			return None, ErrFitType
		default:
			switch y.Kind {
			case Fold, Back, Front:
				return None, ErrFitType
			default:
				if z := x.Cmp(&y); z != 0 {
					return True, nil
				} else {
					return False, nil
				}
			}
		}
	})
}
