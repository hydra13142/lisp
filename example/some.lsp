(define
	sum
	(omission
		(lambda
			(x y)
			(each
				(loop
					()
					(not (atom y))
					(each
						(define z (car y))
						(define y (cdr y))
						(define x (+ x z))
					)
				)
				x
			)
		)
	)
)
(define
	max
	(omission
		(lambda
			(x y)
			(each
				(loop
					()
					(not (atom y))
					(each
						(define z (car y))
						(define y (cdr y))
						(if
							(< x z)
							(define x z)
						)
					)
				)
				x
			)
		)
	)
)
(define
	min
	(omission
		(lambda
			(x y)
			(each
				(loop
					()
					(not (atom y))
					(each
						(define z (car y))
						(define y (cdr y))
						(if
							(> x z)
							(define x z)
						)
					)
				)
				x
			)
		)
	)
)
(define
	(init n)
	(lambda
		()
		(update
			n
			(+ n 1)
		)
	)
)
(quote "ok")

