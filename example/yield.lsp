(define
	(unlimit n)
	(each
		(define n (- n 1))
		(lambda
			()
			(update
				n
				(+ n 1)
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
	(generator l)
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