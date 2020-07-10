package main

import "fmt"

/*
	Go 语言变量作用域
	作用域为已声明标识符所表示的常量、类型、变量、函数或包在源代码中的作用范围。
	Go 语言中变量可以在三个地方声明：
		函数内定义的变量称为局部变量
		函数外定义的变量称为全局变量
		函数定义中的变量称为形式参数
*/

/*
	局部变量
	在函数体内声明的变量称之为局部变量，它们的作用域只在函数体内，参数和返回值变量也是局部变量。
*/

/*
	全局变量
	在函数体外声明的变量称之为全局变量，全局变量可以在整个包甚至外部包（被导出后）使用。
	全局变量可以在任何函数中使用。
	Go 语言程序中全局变量与局部变量名称可以相同，但是函数内的局部变量会被优先考虑。
*/

/* 声明全局变量 */
var x int
var g = 100
var i = 11

/*
	形式参数
	形式参数会作为函数的局部变量来使用。
*/
func sum(i, j int) int {
	fmt.Printf("sum()函数中i值为：%d\n", i)
	fmt.Printf("sum()函数中j值为：%d\n", j)
	return i + j
}

/*
	初始化局部和全局变量
	不同类型的局部和全局变量默认值为：
	数据类型		初始化默认值
	int			0
	float32		0
	pointer		nil
*/

func main() {

	/* 声明局部变量 */
	var a, b, c int
	/* 初始化参数 */
	a = 10
	b = 20
	c = a + b
	fmt.Printf("a = %d, b = %d and c = %d\n", a, b, c)

	/* 使用全局变量 */
	x = a + b + c
	fmt.Printf("x的值为:%d\n", x)

	//Go 语言程序中全局变量与局部变量名称可以相同，但是函数内的局部变量会被优先考虑。
	var g = 10
	fmt.Printf("优先使用的是函数内的局部变量 g = %d\n", g)

	//形式参数会作为函数的局部变量来使用
	var i = 111
	var j = 222
	var k = 0
	fmt.Printf("在main()函数中i值为：%d\n", i)
	fmt.Printf("在main()函数中j值为：%d\n", j)
	k = sum(i, j)
	fmt.Printf("在main()函数中k值为：%d\n", k)

}
