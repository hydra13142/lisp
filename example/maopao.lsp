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
						(define y (min (cdr n) m))
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
		(if
			(atom l)
			()
			(each
				(define s (min l ()))
				(cons (car s) (maopao (cdr s)))
			)
		)
	)
)
(quote "ok")