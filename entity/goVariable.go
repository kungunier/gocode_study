package main

import "fmt"

/*
	值类型和引用类型
	所有像 int、float、bool 和 string 这些基本类型都属于值类型，使用这些类型的变量直接指向存在内存中的值：(int) i = 7
	当使用等号 = 将一个变量的值赋值给另一个变量时，如：j = i，实际上是在内存中将 i 的值进行了拷贝：(int) i = 7 (int) j = 7

	可以通过 &i 来获取变量 i 的内存地址，例如：0xf840000040（每次的地址都可能不一样）。值类型的变量的值存储在栈中。
	内存地址会根据机器的不同而有所不同，甚至相同的程序在不同的机器上执行后也会有不同的内存地址。
	因为每台机器可能有不同的存储器布局，并且位置分配也可能不同。
	更复杂的数据通常会需要使用多个字，这些数据一般使用引用类型保存。
	一个引用类型的变量 r1 存储的是 r1 的值所在的内存地址（数字），或内存地址中第一个字所在的位置。

	这个内存地址为称之为指针，这个指针实际上也被存在另外的某一个字中。
	同一个引用类型的指针指向的多个字可以是在连续的内存地址中（内存布局是连续的），
	这也是计算效率最高的一种存储形式；也可以将这些字分散存放在内存中，每个字都指示了下一个字所在的内存地址。
	当使用赋值语句 r2 = r1 时，只有引用（地址）被复制。
	如果 r1 的值被改变了，那么这个值的所有引用都会指向被修改后的内容，在这个例子中，r2 也会受到影响。
*/
/*
	简短形式，使用 := 赋值操作符
	我们知道可以在变量的初始化时省略变量的类型而由系统自动推断，声明语句写上 var 关键字其实是显得有些多余了，因此我们可以将它们简写为 a := 50 或 b := false。
	a 和 b 的类型（int 和 bool）将由编译器自动推断。
	这是使用变量的首选形式，但是它只能被用在函数体内，而不可以用于全局变量的声明与赋值。使用操作符 := 可以高效地创建一个新的变量，称之为初始化声明。
	如果在相同的代码块中，我们不可以再次对于相同名称的变量使用初始化声明，例如：a := 20 就是不被允许的，
	编译器会提示错误 no new variables on left side of :=，但是 a = 20 是可以的，因为这是给相同的变量赋予一个新的值。
	如果你在定义变量 a 之前使用它，则会得到编译错误 undefined: a。
	如果你声明了一个局部变量却没有在相同的代码块中使用它，同样会得到编译错误a declared and not used。

	但是全局变量是允许声明但不使用的。
	a, b, c := 5, 7, "abc"
	如果想要交换两个变量的值，则可以简单地使用 a, b = b, a，两个变量的类型必须是相同。
*/

// 这种因式分解关键字的写法一般用于声明全局变量
var (
	qjBl1 = "测试声明全局变量1"
	qjBl2 = 23.432423
	qjBl3 = "测试声明全局变量3"
	qjBl4 = 3456.756756
)

func main() {

	/*
		变量声明的三种方法：
		1、指定变量类型，如果没有初始化，则变量默认为零值：var 变量名 type
		2、根据值自行判定变量类型：var 变量名 = value
		3、省略 var, 注意 := 左侧如果没有声明新的变量，就产生编译错误，格式：变量名 := value
	*/
	var bl string
	var bl1, bl2 int
	var bl3 bool
	var bl4 float64
	fmt.Println(bl, bl1, bl2, bl3, bl4)

	var bl5 = true
	fmt.Println(bl5)

	bl6 := "测试:=声明变量"
	fmt.Println(bl6)

	/*
		多变量声明
		//类型相同多个变量, 非全局变量
		var vname1, vname2, vname3 type
		vname1, vname2, vname3 = v1, v2, v3

		//不需要显示声明类型，自动推断
		var vname1, vname2, vname3 = v1, v2, v3

		//出现在 := 左侧的变量不应该是已经被声明过的，否则会导致编译错误
		vname1, vname2, vname3 := v1, v2, v3
	*/

	//全局变量输出
	fmt.Println(qjBl1, qjBl2)

	//如果想要交换两个变量的值，则可以简单地使用 a, b = b, a，两个变量的类型必须是相同。
	qjBl1, qjBl3 = qjBl3, qjBl1
	fmt.Println(qjBl1, qjBl3)
	qjBl2, qjBl4 = qjBl4, qjBl2
	fmt.Println(qjBl2, qjBl4)

	//空白标识符只取需要的值，并抛弃不需要的值，做到不使用声明变量而不报错
	_, b1, c1 := kongbai()
	fmt.Println(b1, c1)

}

/*
	空白标识符 _ 也被用于抛弃值，如值 5 在：_, b = 5, 7 中被抛弃。
	_ 实际上是一个只写变量，你不能得到它的值。这样做是因为 Go 语言中你必须使用所有被声明的变量，但有时你并不需要使用从一个函数得到的所有返回值。
*/
func kongbai() (int, int, string) {
	a, b, c := 123, 456, "测试空白标识符"
	return a, b, c
}