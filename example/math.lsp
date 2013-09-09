(define
	(P m n)
	(if
		(or (> m n) (< m 0))
		(raise "make sure 0 <= m <= n")
		(each
			(define s 1)
			(loop
				(define i m)
				(< i n)
				(each
					(define i (+ i 1))
					(define s (* s i))
				)
			)
			s
		)
	)
)
(define
	(C m n)
	(if
		(or (> m n) (< m 0))
		(raise "make sure 0 <= m <= n")
		(each
			(define s (P m n))
			(loop
				(define i (+ (- n m) 1))
				(> i 1)
				(each
					(define i (- i 1))
					(define s (/ s i))
				)
			)
			s
		)
	)
)
(define
	(abs x)
	(if
		(< x 0)
		(- 0 x)
		x
	)
)
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
(quote "ok")