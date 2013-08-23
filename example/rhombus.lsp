(define
	(show n c)
	(each
		(loop
			()
			(> n 0)
			(each
				(print c)
				(define n (- n 1))
			)
		)
		(none)
	)
)
(define
	(rhombus n)
	(each
		(define i 0)
		(loop
			()
			(<= i n)
			(each
				(show (- n i) " ")
				(show (+ (* i 2) 1) "*")
				(define i (+ i 1))
				(println (none))
			)
		)
		(loop
			(define i (- i 2))
			(>= i 0)
			(each
				(show (- n i) " ")
				(show (+ (* i 2) 1) "*")
				(define i (- i 1))
				(println (none))
			)
		)
		(none)
	)
)
(quote "ok")