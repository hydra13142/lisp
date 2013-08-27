package lisp

func init() {
	Add("macro", func(t []Token, p *Lisp) (ans Token, err error) {
		var (
			a, b, c Token
			x, y    []Name
		)
		switch len(t) {
		case 2:
			a, b = t[0], t[1]
		case 3:
			a, c, b = t[0], t[1], t[2]
			if c.Kind != List {
				return None, ErrFitType
			}
			t = c.Text.([]Token)
			y = make([]Name, len(t))
			for i, u := range t {
				if u.Kind != Label {
					return None, ErrNotName
				}
				y[i] = u.Text.(Name)
			}
		default:
			return None, ErrParaNum
		}
		if a.Kind != List || b.Kind != List {
			return None, ErrFitType
		}
		t = a.Text.([]Token)
		x = make([]Name, len(t))
		for i, u := range t {
			if u.Kind != Label {
				return None, ErrNotName
			}
			x[i] = u.Text.(Name)
		}
		ans = Token{Macro, &Hong{x, b.Text.([]Token), y}}
		return ans, nil
	})
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
		ans = Token{Front, &Lfac{x, b.Text.([]Token), p}}
		return ans, nil
	})
}
