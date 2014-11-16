package parser

import (
	"bytes"
	"math"
	"unicode/utf8"
)

func IsSpace(c byte) bool {
	return c == '\r' || c == '\n' || c == '\t' || c == '\v' || c == '\f' || c == ' '
}

func IsDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func IsLower(c byte) bool {
	return c >= 'a' && c <= 'z'
}

func IsUpper(c byte) bool {
	return c >= 'A' && c <= 'Z'
}

func IsAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func IsFirst(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

func IsLetter(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') || c == '_'
}

func ParseOct(r []byte, L int) (i int64, l int) {
	var c byte
	for l, c = range r {
		if l >= L || c < '0' || c > '7' {
			return
		}
		i = i*8 + int64(c-'0')
	}
	l++
	return
}

func ParseDec(r []byte, L int) (i int64, l int) {
	var c byte
	for l, c = range r {
		if l >= L || c < '0' || c > '9' {
			return
		}
		i = i*10 + int64(c-'0')
	}
	l++
	return
}

func ParseHex(r []byte, L int) (i int64, l int) {
	var c byte
	for l, c = range r {
		if l >= L {
			return
		}
		switch {
		case c >= '0' && c <= '9':
			i = i*16 + int64(c-'0')
		case c >= 'A' && c <= 'Z':
			i = i*16 + int64(c-'A'+10)
		case c >= 'a' && c <= 'z':
			i = i*16 + int64(c-'a'+10)
		default:
			return
		}
	}
	l++
	return
}

func ParseInt(r []byte) (i int64, l int) {
	if len(r) == 0 {
		return
	}
	c, t := false, 0
	switch r[0] {
	case '-':
		c, t = true, 1
	case '+':
		t = 1
	}
	a, j := ParseDec(r[t:], 16)
	if j > 0 {
		if c {
			i = -a
		} else {
			i = a
		}
		l = j + t
	}
	return
}

func xiaoshu(r []byte) (f float64, l int) {
	var c byte
	if len(r) == 0 {
		return
	}
	if r[0] == '.' {
		i, j := 0, 1
		for l, c = range r[1:] {
			if c < '0' || c > '9' {
				break
			}
			i, j = i*10+int(c-'0'), j*10
		}
		if c >= '0' && c <= '9' {
			l++
		}
		if l > 0 {
			f, l = float64(i)/float64(j), l+1
		}
	}
	return
}

func zhishu(r []byte) (i int64, l int) {
	if len(r) == 0 {
		return
	}
	if r[0] == 'e' || r[0] == 'E' {
		a, j := ParseInt(r[1:])
		if j > 0 {
			i, l = a, j+1
		}
	}
	return
}

func ParseFloat(r []byte) (f float64, l int) {
	if len(r) == 0 {
		return
	}
	p, t := false, 0
	switch r[0] {
	case '-':
		p, t = true, 1
	case '+':
		t = 1
	}
	a, i := ParseDec(r[t:], 16)
	b, j := xiaoshu(r[t+i:])
	c, k := zhishu(r[t+i+j:])
	if i > 0 || j > 0 {
		f = float64(a) + b
		if k > 0 {
			f *= math.Pow10(int(c))
		}
		if p {
			f = -f
		}
		l = t + i + j + k
	}
	return
}

func ParseChar(r []byte) (c rune, l int) {
	var i int64
	if len(r) > 0 {
		if r[0] == '\\' {
			if r[1] >= '0' && r[1] <= '7' {
				i, l = ParseOct(r[1:], 3)
				if i < 256 {
					return rune(i), l + 1
				}
			} else {
				l = 2
				switch r[1] {
				case 'x', 'X':
					i, l = ParseHex(r[2:], 2)
					if l != 2 {
						l = 0
					} else {
						l = 4
					}
				case 'u', 'U':
					i, l = ParseHex(r[2:], 4)
					if l != 4 {
						l = 0
					} else {
						l = 6
					}
				case 't':
					i = '\t'
				case 'r':
					i = '\r'
				case 'n':
					i = '\n'
				case 'v':
					i = '\v'
				case 'f':
					i = '\f'
				default:
					i = int64(r[1])
				}
				if l != 0 {
					return rune(i), l
				}
			}
		} else {
			R := bytes.Runes(r)
			c = rune(R[0])
			return c, utf8.RuneLen(c)
		}
	}
	return 0, 0
}
