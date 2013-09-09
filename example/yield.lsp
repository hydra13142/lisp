(load "list")
(define
	infinite
	(default
		(lambda
			(n t)
			(each
				(define n (- n t))
				(lambda
					()
					(update n (+ n t))
				)
			)
		)
		0
		1
	)
)
(define
	(iterator n f)
	(lambda
		()
		(each
			(define x n)
			(update n (f n))
			x
		)
	)
)
(define
	(fbnq)
	(each
		(define a 0)
		(define b 1)
		(lambda
			()
			(each
				(define c a)
				(update a b)
				(update b (+ a c))
				a
			)
		)
	)
)
(define
	(xrange a b)
	(each
		(define a (- a 1))
		(lambda
			()
			(each
				(update a (+ a 1))
				(if
					(< a b)
					a
					(raise "nothing to yield")
				)
			)
		)
	)
)
(define
	every
	(omission
		(lambda
			(l)
			(each
				(define l (plain l))
				(lambda
					()
					(if
						(atom l)
						(raise "nothing to yield")
						(each
							(define y (car l))
							(update l (cdr l))
							y
						)
					)
				)
			)
		)
	)
)
(define
	combine
	(omission
		(lambda
			(l)
			(lambda
				()
				(each
					(if
						(atom l)
						(raise "nothing to yield")
					)
					(define
						t
						(catch
							(define
								x
								((car l))
							)
						)
					)
					(if
						t
						(each
							(update l (cdr l))
							(self)
						)
						x
					)
				)
			)
		)
	)
)
(quote "ok")