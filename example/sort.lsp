(load "list")
(define
	(QuickSort s)
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
(define
	(BubbleSort l)
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
(define
	(MergeSort l)
	(each
		(define
			(make l)
			(if
				(atom l)
				()
				(cons
					(cons (car l) ())
					(self (cdr l))
				)
			)
		)
		(define
			(reduce l)
			(if
				(atom l)
				l
				(each
					(define x (car l))
					(define p (cdr l))
					(if
						(atom p)
						l
						(each
							(define y (car p))
							(define p (cdr p))
							(cons
								(merge x y)
								(self p)
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
				(define L (make l))
				(until
					(atom (cdr L))
					(define L (reduce L))
				)
				(car L)
			)
		)
	)
)
(quote "ok")
