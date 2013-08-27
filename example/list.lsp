(define
	(len l)
	(if
		(atom l)
		0
		(each
			(loop
				(define i 0) # 使用label表示声明一个变量
				(not (atom l))
				(each
					(define l (cdr l))
					(define i (+ i 1))
				)
			)
			i
		)
	)
)
(define
	(index l i)
	(if
		(catch # 捕捉错误并转化为一个字符串数据
			(each
				(loop
					() # 空表不会被执行，返回自身
					(!= i 0)
					(each
						(define i (- i 1))
						(define l (cdr l))
					)
				)
				(define c (car l))
			)
		)
		(raise "out of range") # 产生错误
		c # 宏和内置函数的执行不会产生内层环境
	)
)
(define
	(reverse l)
	(each
		(loop
			(define s ())
			(not (atom l))
			(each
				(define i (car l))
				(define l (cdr l))
				(define s (cons i s))
			)
		)
		s
	)
)
(define
	(map l f)
	(each
		(define s ())
		(loop
			()
			(not (atom l))
			(each
				(define i (car l))
				(define l (cdr l))
				(define s (cons (f i) s))
			)
		)
		(reverse s)
	)
)
(define
	(filter l f)
	(each
		(define s ())
		(loop
			()
			(not (atom l))
			(each
				(define i (car l))
				(define l (cdr l))
				(if
					(f i)
					(define s (cons i s))
					()
				)
			)
		)
		(reverse s)
	)
)
(define
	(range i j)
	(each
		(loop
			(define l ())
			(< i j)
			(each
				(define j (- j 1))
				(define l (cons j l))
			)
		)
		l
	)
)
(quote "ok")