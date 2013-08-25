# 井号表示该行剩余部分为注释
# 我的lisp方言内置实现了if和loop
# 当然lisp优点之一就是可以自编程，以下是if和loop的宏版本
#
# if的宏版本很容易理解，只不过我内置实现的可以省略第三个参数
#  |(define
#  |	'(if a b c) # 用'保护列表，表示声明一个宏
#  |	(cond
#  |		(a b)
#  |		(1 c)
#  |	)
#  |)
#
# loop的宏版本实现，使用了宏的递归，each表示依次执行，最后一个作为返回
#  |(define
#  |	'(while b c)
#  |	(each
#  |		(if
#  |			b
#  |			(none)
#  |			(each
#  |				c
#  |				(while b c) # 宏用法和函数一样
#  |			)
#  |		)
#  |	)
#  |)
#  |(define
#  |	'(loop a b c)
#  |	(each
#  |		a
#  |		(while b c)
#  |	)
#  |)
#
# 一般lisp实现都默认有支持顺序执行的，但是本lisp方言要求用each标注
# 如此一来，我们就在lisp里实现了对顺序、选择、循环三种基本结构的支持
# 没有break，但本lisp方言内置有错误系统，可以用产生、捕捉错误来实现

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