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
	'(loop a b c)
	(each
		a
		(if
			b
			(none)
			(each
				c
				(loop () b c)
			)
		)
	)
)
(define
	lambda
	(solid
		(macro
			(p c)
			(lambda
				p
				(eval
					(cons
						(define this (lambda p c))
						(quote p)
					)
				)
			)
		)
	)
)
(quote "ok")