package main

import "lisp"

func main() {
	lisp := lisp.NewLisp()
	lisp.IO()
	lisp.EX()
	lisp.Eval(`
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
