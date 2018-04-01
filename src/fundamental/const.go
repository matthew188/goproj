package fundamental

import (
	"unsafe"
	"fmt"
)

type ByteSize float64

func Test_const() {
	const LEFT int = 10
	const RIFGT int = 20
	println(LEFT, RIFGT)

	const (
		a = "abcd"
		b = len(a)
		c = unsafe.Sizeof(a)
	)
	println(a, b, c)

	const (
		aa = iota //0
		bb        //1
		cc        //2
		d  = "ha" //独立值，iota += 1
		e         //"ha"   iota += 1
		f  = 100  //iota +=1
		g         //100  iota +=1
		h  = iota //7,恢复计数
		i         //8
	)
	fmt.Println(aa, bb, cc, d, e, f, g, h, i)

	const (
		_           = iota             // ignore first value by assigning to blank identifier
		KB ByteSize = 1 << (10 * iota) // 1 << (10*1)
		MB                             // 1 << (10*2)
		GB                             // 1 << (10*3)
		TB                             // 1 << (10*4)
		PB                             // 1 << (10*5)
	)
	fmt.Println(KB, MB, GB, TB, PB)

}
