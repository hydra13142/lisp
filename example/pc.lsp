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
(quote "ok")