package main

import "github.com/hydra13142/lisp"

func main() {
	lisp.IO()
	lisp.EX()
	console := lisp.NewLisp()
	console.Eval(`
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
	`)
}