(define
	if
	(default
		(macro
			(a b c)
			(cond (a b) (1 c))
		)
		(none)
	)
)
(define
	'(while b c)
	(each
		(if
			b
			(each
				c
				(while b c)
			)
			(none)
		)
	)
)
(define
	'(until b c)
	(each
		(if
			b
			(none)
			(each
				c
				(until b c)
			)
		)
	)
)
(define
	'(loop a b c)
	(each
		a
		(while b c)
	)
)
(define
	'(for a b c)
	(lambda
		()
		(each
			(define _ b)
			(until
				(catch
					(define a (_))
				)
				c
			)
		)
	)
)
(define
	lambda
	(macro
		(p c)
		((builtin lambda)
			p
			(eval
				(cons
					(define this ((builtin lambda) p c))
					(quote p)
				)
			)
		)
	)
)
(quote "ok")