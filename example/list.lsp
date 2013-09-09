(define
	list
	(omission
		(lambda
			(l)
			(quote l)
		)
	)
)
(define
	(range i j)
	(each
		(loop
			(define l ())
			(< i j)
			(each
				(define j (- j 1))
				(define l (cons j l))
			)
		)
		l
	)
)
(define
	(len l)
	(if
		(atom l)
		0
		(each
			(loop
				(define i 0)
				(not (atom l))
				(each
					(define l (cdr l))
					(define i (+ i 1))
				)
			)
			i
		)
	)
)
(define
	(index l i)
	(if
		(catch
			(each
				(loop
					()
					(!= i 0)
					(each
						(define i (- i 1))
						(define l (cdr l))
					)
				)
				(define c (car l))
			)
		)
		(raise "out of range")
		c
	)
)
(define
	(reverse l)
	(each
		(loop
			(define s ())
			(not (atom l))
			(each
				(define i (car l))
				(define l (cdr l))
				(define s (cons i s))
			)
		)
		s
	)
)
(define
	(slice l a b)
	(if
		(>= a b)
		()
		(if
			(> a 0)
			(self
				(cdr l)
				(- a 1)
				(- b 1)
			)
			(cons
				(car l)
				(self
					(cdr l)
					0
					(- b 1)
				)
			)
		)
	)
)
(define
	(merge a b)
	(each
		(define
			(cmb a b c)
			(cond
				((atom a)
					(+ (reverse c) b)
				)
				((atom b)
					(+ (reverse c) a)
				)
				(1
					(each
						(define x (car a))
						(define y (car b))
						(if
							(< x y)
							(self
								(cdr a)
								b
								(cons x c)
							)
							(self
								a
								(cdr b)
								(cons y c)
							)
						)
					)
				)
			)
		)
		(cmb a b ())
	)
)
(define
	(map l f)
	(each
		(define s ())
		(loop
			()
			(not (atom l))
			(each
				(define i (car l))
				(define l (cdr l))
				(define s (cons (f i) s))
			)
		)
		(reverse s)
	)
)
(define
	(filter l f)
	(each
		(define s ())
		(loop
			()
			(not (atom l))
			(each
				(define i (car l))
				(define l (cdr l))
				(if
					(f i)
					(define s (cons i s))
					()
				)
			)
		)
		(reverse s)
	)
)
(quote "ok")