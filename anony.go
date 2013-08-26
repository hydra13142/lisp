package lisp

func init() {
	Add("macro", func(t []Token, p *Lisp) (ans Token, err error) {
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
		x := make([]Name, len(t))
		for i, c := range t {
			if c.Kind != Label {
				return None, ErrNotName
			}
			x[i] = c.Text.(Name)
		}
		ans = Token{Macro, &Hong{x, b.Text.([]Token)}}
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
