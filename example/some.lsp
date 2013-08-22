(define
	sum
	(omission
		(lambda
			(x y)
			(cond
				((atom y)
					x
				)
				(1
					(self (+ x (car y)) (cdr y))
				)
			)
		)
	)
)
(define
	max
	(omission
		(lambda
			(x y)
			(cond
				((atom y)
					x
				)
				((>= x (car y))
					(self x (cdr y))
				)
				(1
					(self (car y) (cdr y))
				)
			)
		)
	)
)
(define
	min
	(omission
		(lambda
			(x y)
			(cond
				((atom y)
					x
				)
				((<= x (car y))
					(self x (cdr y))
				)
				(1
					(self (car y) (cdr y))
				)
			)
		)
	)
)
(define
	(init n)
	(lambda
		()
		(define
			n
			(+ n 1)
		)
	)
)
(quote "ok")