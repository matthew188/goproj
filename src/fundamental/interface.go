package fundamental

type Phone interface {
	call()
	eat()
}
type Animal interface {
	eat()
}
type dog struct {

}
type Iphone struct {

}
type Oneplus struct {

}
func (iphone Iphone) call() {
	println("iphone de call 方法")
}
func (iphone Iphone) eat() {
	println("iphone de eat 方法")
}
func (a dog) eat(){
	println("dog eat")
}
func (one Oneplus) call(){
	println("oneplus的call方法")
}


func InterfaceTest(){
	var phone Phone
	phone =new(Iphone)
	phone.call()

	newdog := new(dog)
	newdog.eat()

}