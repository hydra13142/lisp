(define
	+=
	(macro
		(a b)
		(define a (+ a b))
	)
)
(define
	-=
	(macro
		(a b)
		(define a (- a b))
	)
)
(define
	*=
	(macro
		(a b)
		(define a (* a b))
	)
)
(define
	/=
	(macro
		(a b)
		(define a (/ a b))
	)
)
(define
	%=
	(macro
		(a b)
		(define a (% a b))
	)
)
(define
	generator
	(macro
		(p q)
		(i o)
		(define
			p
			(each
				(define i (chan))
				(define o (chan))
				(define
					(yield n)
					(each
						(o n)
						(i)
					)
				)
				(define p q)
				(go
					(each
						(i)
						p
						(close i)
						(close o)
					)
				)
				(lambda
					()
					(if
						(catch
							(each
								(i 1)
								(define x (o))
							)
						)
						(raise "nothing to yield")
						x
					)
				)
			)
		)
	)
)
(define
	trace
	(macro
		(p q)
		(define
			p 
			(each
				(print "call (" (car (quote p)))
				(for
					i
					(cdr (quote p))
					(print " " (eval i))
				)
				(println ")")
				q
			)
		)
	)
)
(define
	rewrite
	(macro
		(m f)
		(l)
		(eval
			(cons
				(quote define)
				(cons 
					(car (quote f))
					(cons
						(
							(lambda
								()
								(each
									(define
										l
										(eval (car (quote f)))
									)
									(m
										f
										(eval
											(cons
												l
												(cdr (quote f))
											)
										)
									)
								)
							)
						)
						()
					)
				)
			)
		)
	)
)
(quote "ok")