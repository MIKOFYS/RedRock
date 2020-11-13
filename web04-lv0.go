package main

import (
	"fmt"
)

func f1(p *int) int {
	*p++
	return *p
}

func f2(p *int) int {
	*p *= 2
	return *p
}

func makeFunc() myFunc{
	return func (p *int) int {
		*p++
		return *p
	}
}
//makeFunc用来返回一个myFunc类型的匿名函数 (该返回值可以直接作为do函数的参数)

type myFunc func(p *int)int

func do(a myFunc, b myFunc) {
	x := 1
	// 调用变量并获取返回值
	// 声明的变量本身就是一种函数变量
	fmt.Println("f1 result: ", a(&x))
	fmt.Println("f2 result: ", b(&x))
}
//x = 0
//f1 result:  2
//f2 result:  4
//此时f2中p的值用的是经过f1之后的值


func work() {
	fmt.Println("加油！打工人！")
}

func squares() int {
	x := 0
		x++
		return x * x
}


func main(){
	/*
	map
	声明
	var m map[string]string
	声明+初始化（为map分配内存并返回一个初始值）
	var m map[string]string = make(map[string]string)
	m := make(map[string]string)
	声明+初始化+赋值
	var m map[string]string = map[string]string{
	"a":"aaa",
	"b":"bbb",
	}
	m := map[string]string{
		"a":"aaa",
		"b":"bbb",
		}
	*/


	//map 查询
	//m := map[string]string{
	//	"a":"aaa",
	//	"b":"bbb",
	//}
	//value, ok := m["a"]
	//fmt.Println(value,ok)     aaa true
	//v, o := m["c"]
	//fmt.Println(v,o)     false

	//map 删除
	//m := map[string]string{
	//	"a":"aaa",
	//	"b":"bbb",
	//}
	//delete(m, "b")
	//fmt.Println(m)    map[a:aaa]

	//map 遍历
	//m := map[string]string{
	//	"a":"aaa",
	//	"b":"bbb",
	//}
	//for key, value := range m{
	//	fmt.Println(key, value)
	//}
	//a aaa
	//b bbb

	//map是引用类型，在函数引用它时，实际上传的是它的指针

	//签名函数相当于函数中的常数
	//可以将函数赋给一个变量
	//c := work
	//c()

	//do(makeFunc(), func(p *int) int {
	//	*p *= 2
	//	return *p
	//})
	//前一个makeFunc返回一个与后面相同的匿名函数

	f3 := squares()
	f2 := squares()
	fmt.Println("first call f1:", f3)
	fmt.Println("second call f1:", f3)

	fmt.Println("first call f2:", f2)
	fmt.Println("second call f2:", f2)
	fmt.Println("third call f2:", f2)




	//在defer中的函数的参数表里有切片类型的变量，那么defer会把拷贝切片的指针（因为是引用类型）。
	//但是如果参数表里有数组变量，defer会拷贝整个数组（因为数组是值类型），进而造成不必要的时间和空间开销。

	a := 10
	defer func(i int) {
		fmt.Println("defer func i is ", i)	// 后：defer func i is  10
	}(a)

	a += 10
	fmt.Println("after defer a is ", a)	// 先：after defer a is  20

	//recover()只有在defer后面的函数体内被直接调用才能捕获panic并打断当前的 panic，进行处理修复，
	//保证不会让单个 handler 影响到整个程序。否则会返回nil，异常继续向外传递
	//连续多个panic的场景只能出现在延迟调用里面,但延迟调用里面只有最后一次panic能被捕获
	

}

