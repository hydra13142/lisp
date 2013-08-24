(define
	lambda
	(solid
		(define
			'(_ p c)
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
