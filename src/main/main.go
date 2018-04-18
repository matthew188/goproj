package main

import (
	"fmt"
	"fundamental"
)

func main() {
	fmt.Println("hello")
	//println(1 <<10)

	//1 常量测试
	//fundamental.Test_const();

        //modify 2
	//2 for质数测试
	//fundamental.UseWhileFindPrim();

        //modify 3
	//3 函数测试 找最大值
	//fmt.Println("最大值为:",fundamental.FindMax(12123.432,5432))

	//4 数组测试
	//fundamental.ArrayTest();

	//5 指针测试
	//var a ,b int=11,22
	//fundamental.Swap(&a,&b)
	//fmt.Println(a ,b)
	//var x ,y int=55,77
	//fundamental.CannotSwapAddr(x,y)
	//fmt.Println(x,y)

	//6
        //6 结构体
	//var book fundamental.Books  //调用其他包的结构体 要加包名
	////fundamental.PrintBooks(book)
	//fundamental.PprintBooks(&book)

	//7 切片
	//fundamental.SliceTest()

	// 8  range函数
	//fundamental.RangeTest()

	// 9 map数据类型
	//fundamental.MapTest()

	//10 递归
	//var i uint32
	//for i = 0; i < 10; i++ {
	//	println(i,"的费布拉奇数是 ", fundamental.Fibonacci(i))
	//}

	//11 强制类型转换
	//fundamental.Forcetrans()

	//12 接口啊 interface
	// fundamental.InterfaceTest()
        //8
	//13 error错误处理
	fundamental.ErrorTest()
}
