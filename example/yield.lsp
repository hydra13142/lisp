
(define
	(generator l)
	(lambda
		()
		(if
			(atom l)
			(raise "nothing to yield")
			(each
				(define y (car l))
				(update l (cdr l))
				y
			)
		)
	)
)
(define
	combine
	(omission
		(lambda
			(l)
			(lambda
				()
				(each
					(if
						(atom l)
						(raise "nothing to yield")
					)
					(define
						t
						(catch
							(define
								x
								((car l))
							)
						)
					)
					(if
						t
						(each
							(update l (cdr l))
							(self)
						)
						x
					)
				)
			)
		)
	)
)