(define
	'(loop a b c)
	(each
		a
		(cond
			(b (none))
			(1
				(each
					c
					(loop () b c)
				)
			)
		)
	)
)
(define
	if
	(default
		(define
			'(_ a b c)
			(cond (a b) (1 c))
		)
		(none)
	)
)