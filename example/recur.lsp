(define
	lambda
	(stable
		(define
			'(recur p c)
			(lambda
				p
				(eval
					(cons
						(define
							this
							(lambda p c)
						)
						(quote p)
					)
				)
			)
		)
	)
)
(quote "ok")