(define
	(unlimit n)
	(each
		(define n (- n 1))
		(lambda
			()
			(update
				n
				(+ n 1)
			)
		)
	)
)
(define
	(xrange a b)
	(each
		(define a (- a 1))
		(lambda
			()
			(each
				(update a (+ a 1))
				(if
					(< a b)
					a
					(raise "nothing to yield")
				)
			)
		)
	)
)
(define
	(evary l)
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
	generator
	(macro
		(p q) # 输入的参数
		(i o) # 表示额外的替换对象，将替换为随机不重名label
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
(quote "ok")