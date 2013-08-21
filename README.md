lisp

====

a simple lisp made by go

一个简单的go-lisp实现，基本遵循lisp标准语法，包含一些（可能）独有的函数或规则

支持输入四种基本类型：整数、浮点数、字符、字符串

内部支持三种基本类型：整数、浮点数、字符（字符作为整数保存）

quote的简写形式，只能用于列表；作为对应，所有的元素都可以“运算”，除了列表之外的元素运算只是简单返回自身

最基本的7种操作quote、atom、eq、car、cdr、cons、cond都是支持的；

支持四则运算、比较运算、逻辑运算（逻辑运算和cons是懒惰执行的）

另支持如下操作：

lambda 产生一个匿名函数

define 声明

	(define (f x) (+ x 2)) 为声明一个函数f ：f(x)=x+2
	
	(define f (+ 2 1))     为声明一个变量f : f=3
	
eval 为将一个列表在当前作用域下执行

each 为顺序执行多个语句，最后一个语句的返回值作为返回值

可以通过EX方法注入4个函数：

	if 三个参数，根据第一个的结果决定执行第二个还是第三个

	loop 三个参数，第一个初始化，第二个判断，循环执行第三个，直到第二个判断为假

	default 用于给函数绑定默认值，返回一个绑定了默认值（因而可以省略后面参数）的函数

	omission 用于产生一个可变参数函数，提供的参数必须是一个函数，该函数的最后一个参数应该为列表

可以通过IO方法注入3个函数

	scan 从控制台获取字符串数据，并将该数据作为lisp语句执行后返回

	print 输出数据

	println 输出数据并回车

可以通过Add方法添加自定义函数：

	func (l *Lisp)Add(name string, func([]Token,*Lisp)(Token,error))

Token为内部用来表示元素的类型

	type Token struct{
		Kind
		Text interface{}
	}

Text只可能装入如下类型：[]Token、int64、float64、string

对应的Kind值分别为	    List     Int    Float    String

注意的是为了实现惰性求值，你添加的函数接收到的切片，每个元素都是未运算的，需要你根据需要进行运算（或直接解包）

具体使用参见example
