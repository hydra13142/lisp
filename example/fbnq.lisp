(define
	(fbnq n)
	(cond
		((<= n 1)
			1
		)
		(1
			(loop
				(each
					(define i 0)
					(define a 1)
					(define b 1)
				)
				(< i n)
				(each
					(define i (+ i 1))
					(define c b)
					(define b (+ a b))
					(define a c)
				)
			)
		)
	)
)
(println
	(fbnq 0)
	(fbnq 1)
	(fbnq 2)
	(fbnq 3)
	(fbnq 4)
	(fbnq 5)
	(fbnq 6)
	(fbnq 7)
	(fbnq 8)
	(fbnq 9)
)