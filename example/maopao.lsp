(define
	(maopao l)
	(each
		(define
			(min n m)
			(cond
				((atom n)
					m
				)
				((atom (cdr n))
					(+ n m)
				)
				(1
					(each
						(define x (car n))
						(define y (min (cdr n) m))
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
		(cond
			((atom l)
				()
			)
			(1
				(each
					(define s (min l ()))
					(cons (car s) (maopao (cdr s)))
				)
			)
		)
	)
)
(quote "ok")