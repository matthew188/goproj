package fundamental

import "fmt"

type Books struct {
	title  string
	author string
	name string
	price int
}
func PrintBooks(book1 Books){
	book1.author = "matt yuan"
	book1.title = "good book you need read"
	book1.name = "galang execise"
	book1.price = 88
	fmt.Println(book1.title)
	fmt.Println(book1.name)
	fmt.Println(book1.author)
	fmt.Println(book1.price)

}
func PprintBooks(book1 *Books){
	book1.author = "matt yuan"
	book1.title = "good book you need read"
	book1.name = "galang execise"
	book1.price = 88
	fmt.Println(book1.title)
	fmt.Println(book1.name)
	fmt.Println(book1.author)
	fmt.Println(book1.price)
}
