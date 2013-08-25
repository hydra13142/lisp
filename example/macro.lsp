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
			(none)
			(each
				c
				(while b c)
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