(define
	(list l)
	(cond
		((atom l)
			(cond
				((eq l '())
					1
				)
				(1 0)
			)
		)
		(1 1)
	)
)
(define
	(len l)
	(cond
		((atom l)
			0
		)
		(1
			(+ (len (cdr l)) 1)
		)
	)
)
(define
	(index l n)
	(cond
		((atom l)
			(raise "out of range")
		)
		((== n 0)
			(car l)
		)
		(1
			(index (cdr l) (- n 1))
		)
	)
)
(define
	(reverse l)
	(each
		(define
			(rev s c)
			(cond
				((atom s)
					c
				)
				(1
					(rev
						(cdr s)
						(cons (car s) c)
					)
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
			(cond
				((atom s)
					c
				)
				((f (car s))
					(pick
						(cdr s)
						(cons (car s) c)
					)
				)
				(1
					(pick (cdr s) c)
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
			(cond
				((atom s)
					c
				)
				(1
					(change
						(cdr s)
						(cons (f (car s)) c)
					)
				)
			)
		)
		(reverse (change l '()))
	)
)
(define
	(range a b)
	(cond
		((< a b)
			(cons
				a
				(range (+ a 1) b)
			)
		)
		(1 '())
	)
)
(define
	(quicksort s)
	(cond
		((atom s)
			s
		)
		(1
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
				(+ (+ (quicksort a) b) (quicksort c))
			)
		)
	)
)