(define
	(len l)
	(if
		(atom l)
		0
		(+ (self (cdr l)) 1)
	)
)
(define
	(index l n)
	(if
		(atom l)
		(raise "out of range")
		(if
			(== n 0)
			(car l)
			(self (cdr l) (- n 1))
		)
	)
)
(define
	(reverse l)
	(each
		(define
			(rev s c)
			(if
				(atom s)
				c
				(self
					(cdr s)
					(cons (car s) c)
				)
			)
		)
		(rev l '())
	)
)
(define
	(filter l f)
	(each
		(define
			(pick s c)
			(if
				(atom s)
				c
				(if
					(f (car s))
					(self
						(cdr s)
						(cons (car s) c)
					)
					(self (cdr s) c)
				)
			)
		)
		(reverse (pick l '()))
	)
)
(define
	(map l f)
	(each
		(define
			(change s c)
			(if
				(atom s)
				c
				(self
					(cdr s)
					(cons (f (car s)) c)
				)
			)
		)
		(reverse (change l '()))
	)
)
(define
	(range a b)
	(if
		(< a b)
		(cons a (self (+ a 1) b))
		()
	)
)
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