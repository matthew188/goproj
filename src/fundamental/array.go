package fundamental

import "fmt"

type printType struct {
	myint   []int
	myfloat []float32
}

func ArrayTest() {
	var arr = [8]float32{1, 2, 3, 4, 5.2};
	var salary = arr[4]
	fmt.Println(arr, salary);
}

func SliceTest() {
	var myarr printType
	var slice []int
	slice = [] int{12, 3, 3, 4, 3, 4}
	slice = append(slice, 6)
	var slice1 []float32 = make([]float32, 5)
	fmt.Println(slice, slice1)
	myarr.myint = slice
	myarr.myfloat = slice1
	printSlice(myarr)
	slice2 :=make([]int , len(slice) ,cap(slice)*3)
	copy(slice2,slice)
	myarr.myint=slice2
	printSlice(myarr)
}
func printSlice(slice printType ) {
	if slice.myint != nil {
		fmt.Println("int 数组",slice.myint)
	}
	if slice.myfloat != nil {
		fmt.Println("float 数组 ",slice.myfloat)
	}
}

func Swap(x *int, y *int) {
	//fmt.Println("交换前",x,y)
	var temp int
	temp = *x /* 保存 x 地址的值 */
	*x = *y   /* 将 y 赋值给 x */
	*y = temp /* 将 temp 赋值给 y */
	//fmt.Println("交换后",x,y)
}
func CannotSwapAddr(x int, y int) { //这么交换没有意义
	var tmp *int
	tmp = &x
	x = y
	y = *tmp
}
