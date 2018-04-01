package fundamental

func Fibonacci(a uint32) (result uint32) {
	if a < 2 {
		result =  a
	} else {
		result = Fibonacci(a-1) + Fibonacci(a-2)
	}
	return result
}

func Forcetrans(){
	var value int = 123
	println(float32(value))
}