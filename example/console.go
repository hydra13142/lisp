package main

import . "lisp"

func main(){
	lisp:=NewLisp()
	lisp.IO()
	lisp.EX()
	lisp.Eval(`
	(loop
		()
		1
		(each
			(println "?:")
			(println (scan))
		)
	)`)
}