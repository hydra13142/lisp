package lisp

import (
	"io/ioutil"
	"os"
)

func (l *Lisp) Load(s string) (Token, error) {
	file, err := os.Open(s)
	if err != nil {
		return None, err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return None, err
	}
	return l.Eval(string(data))
}
