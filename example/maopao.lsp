(define
	(maopao l)
	(each
		(define
			(min n)
			(if
				(atom n)
				()
				(each
					(define x (car n))
					(define y (self (cdr n)))
					(if
						(atom y)
						(cons x ())
						(each
							(define z (car y))
							(if
								(<= x z)
								(cons x y)
								(cons z (cons x (cdr y)))
							)
						)
					)
				)
			)
		)
		(if
			(atom l)
			()
			(each
				(define s (min l))
				(cons (car s) (self (cdr s)))
			)
		)
	)
)
(quote "ok")