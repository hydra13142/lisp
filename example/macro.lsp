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
	while
	(macro
		(b c)
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
)
(define
	until
	(macro
		(b c)
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
)
(define
	loop
	(macro
		(a b c)
		(each
			a
			(while b c)
		)
	)
)
(define
	for
	(pretreat
		(macro
			(a b c)
			(until
				(catch
					(define a (b))
				)
				c
			)
		)
		0
		1
		0
	)
)
(define
	lambda
	(macro
		(p c)
		((builtin lambda)
			p
			(each
				(define
					this
					((builtin lambda)
						p
						c
					)
				)
				(eval
					(cons
						(quote this)
						(quote p)
					)
				)
			)
		)
	)
)
(quote "ok")