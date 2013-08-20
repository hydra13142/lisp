package main

import "github.com/hydra13142/lisp"

func main() {
	lisp := lisp.NewLisp()
	lisp.Extend()
	lisp.Eval(`
		(define
			(quicksort s)
			(each
				(define
					(end l)
					(cond
						((atom l)
							1
						)
						((atom (cdr l))
							1
						)
						(1 0)
					)
				)
				(define
					(above a b c)
					(cond
						((atom b)
							c
						)
						((<= a (car b))
							(above
								a
								(cdr b)
								(cons (car b) c)
							)
						)
						(1
							(above a (cdr b) c)
						)
					)
				)
				(define
					(below a b c)
					(cond
						((atom b)
							c
						)
						((> a (car b))
							(below
								a
								(cdr b)
								(cons (car b) c)
							)
						)
						(1
							(below a (cdr b) c)
						)
					)
				)
				(define
					(sort s)
					(cond
						((end s)
							s
						)
						(1
							(each
								(define n (car s))
								(+
									(sort (below n (cdr s) '()))
									(cons n (sort (above n (cdr s) '())))
								)
							)
						)
					)
				)
				(sort s)
			)
		)`)
	lisp.Eval(`
		(define
			(fbnq n)
			(cond
				((<= n 1)
					1
				)
				(1
					(+
						(fbnq (- n 1))
						(fbnq (- n 2))
					)
				)
			)
		)`)
	lisp.Eval(`
		(define
			(reverse s)
			(each
				(define
					(rev s c)
					(cond
						((atom s)
							c
						)
						(1
							(rev
								(cdr s)
								(cons (car s) c)
							)
						)
					)
				)
				(rev s '())
			)
		)`)
	lisp.Eval(`
		(define
			(map s f)
			(each
				(define
					(rev s c)
					(cond
						((atom s)
							c
						)
						(1
							(rev
								(cdr s)
								(cons (f (car s)) c)
							)
						)
					)
				)
				(reverse (rev s '()))
			)
		)`)
	lisp.Eval(`
		(println
			(map
				(quicksort '(1 4 2 8 5 7 3 9 0 6))
				fbnq
			)
		)`)
}
