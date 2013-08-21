package main

import . "lisp"

func main() {
	lisp := NewLisp()
	lisp.IO()
	lisp.EX()
	lisp.Load("C:\\Users\\liudiwu\\Documents\\GitHub\\lisp\\example\\fbnq.lisp")
	lisp.Eval(`
	(loop
		()
		1
		(each
			(println "?:")
			(try
				(error
					(println (scan))
				)
			)
		)
	)`)
}
