package lisp

func init() {
	Add("+", func(t []Token, p *Lisp) (Token, error) {
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
		case Int:
			switch y.Kind {
			case Int:
				return Token{Int, x.Text.(int64) + y.Text.(int64)}, nil
			case Float:
				return Token{Float, float64(x.Text.(int64)) + y.Text.(float64)}, nil
			}
		case Float:
			switch y.Kind {
			case Int:
				return Token{Float, x.Text.(float64) + float64(y.Text.(int64))}, nil
			case Float:
				return Token{Float, x.Text.(float64) + y.Text.(float64)}, nil
			}
		case String:
			switch y.Kind {
			case String:
				return Token{String, x.Text.(string) + y.Text.(string)}, nil
			}
		case List:
			switch y.Kind {
			case List:
				a, b := x.Text.([]Token), y.Text.([]Token)
				c := make([]Token, len(a)+len(b))
				copy(c, a)
				copy(c[len(a):], b)
				return Token{List, c}, nil
			}

		}
		return None, ErrFitType
	})
	Add("-", func(t []Token, p *Lisp) (Token, error) {
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
		case Int:
			switch y.Kind {
			case Int:
				return Token{Int, x.Text.(int64) - y.Text.(int64)}, nil
			case Float:
				return Token{Float, float64(x.Text.(int64)) - y.Text.(float64)}, nil
			}
		case Float:
			switch y.Kind {
			case Int:
				return Token{Float, x.Text.(float64) - float64(y.Text.(int64))}, nil
			case Float:
				return Token{Float, x.Text.(float64) - y.Text.(float64)}, nil
			}
		}
		return None, ErrFitType
	})
	Add("*", func(t []Token, p *Lisp) (Token, error) {
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
		case Int:
			switch y.Kind {
			case Int:
				return Token{Int, x.Text.(int64) * y.Text.(int64)}, nil
			case Float:
				return Token{Float, float64(x.Text.(int64)) * y.Text.(float64)}, nil
			}
		case Float:
			switch y.Kind {
			case Int:
				return Token{Float, x.Text.(float64) * float64(y.Text.(int64))}, nil
			case Float:
				return Token{Float, x.Text.(float64) * y.Text.(float64)}, nil
			}
		}
		return None, ErrFitType
	})
	Add("/", func(t []Token, p *Lisp) (Token, error) {
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
		case Int:
			switch y.Kind {
			case Int:
				return Token{Int, x.Text.(int64) / y.Text.(int64)}, nil
			case Float:
				return Token{Float, float64(x.Text.(int64)) / y.Text.(float64)}, nil
			}
		case Float:
			switch y.Kind {
			case Int:
				return Token{Float, x.Text.(float64) / float64(y.Text.(int64))}, nil
			case Float:
				return Token{Float, x.Text.(float64) / y.Text.(float64)}, nil
			}
		}
		return None, ErrFitType
	})
	Add("%", func(t []Token, p *Lisp) (Token, error) {
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
		if x.Kind == Int && y.Kind == Int {
			return Token{Int, x.Text.(int64) % y.Text.(int64)}, nil
		}
		return None, ErrFitType
	})
}
