(define
	(quicksort s)
	(if
		(atom s)
		s
		(each
			(define n (car s))
			(define
				a
				(filter
					s
					(lambda (x) (< x n))
				)
			)
			(define
				b
				(filter
					s
					(lambda (x) (== x n))
				)
			)
			(define
				c
				(filter
					s
					(lambda (x) (> x n))
				)
			)
			(+ (+ (self a) b) (self c))
		)
	)
)
(quote "ok")