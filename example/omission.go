package main

import "github.com/hydra13142/lisp"

func main() {
	console := lisp.NewLisp()
	console.Eval([]byte(`
		(each
			(define
				(g x y)
				(cons x y)
			)
			(define
				G
				(omission g)
			)
			(println
				(G 3)
			)
			(println
				(G 3 2)
			)
			(println
				(G 3 2 1)
			)
		)
	`))
}