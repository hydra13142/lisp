package main

import "github.com/hydra13142/lisp"

func main() {
	lisp.IO()
	lisp.EX()
	console := lisp.NewLisp()
	console.Eval(`
		(each
			(define
				(f x y z)
				(+ (+ x y) z)
			)
			(define
				F
				(default f 1 2 3)
			)
			(println
				(F)
			)
			(println
				(F 3)
			)
			(println
				(F 3 2)
			)
			(println
				(F 3 2 1)
			)
		)
	`)
}
