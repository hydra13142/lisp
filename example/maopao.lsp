(define
	(maopao l)
	(each
		(define
			(min n m)
			(if
				(atom n)
				m
				(if
					(atom (cdr n))
					(+ n m)
					(each
						(define x (car n))
						(define y (self (cdr n) m))
						(define z (car y))
						(cond
							((<= x z)
								(cons x y)
							)
							(1
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
				(define s (min l ()))
				(cons (car s) (self (cdr s)))
			)
		)
	)
)
(quote "ok")