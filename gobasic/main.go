package main

import "fmt" // fmt is a package that implements formatted I/O with functions analogous to C's printf and scanf.

func main() {
	var name string = "John Doe" // string is a sequence of bytes
	var age int = 25             // int is a signed integer type that is at least 32 bits in size
	const isCool = true          // bool is a boolean type
	friend := "Jack"             // := is the short assignment statement
	friend_age := 30             // := is the short assignment statement
	fmt.Println("Hello, World!")
	fmt.Println("I am learning Go!")
	fmt.Println("My name is", name, "and I am", age, "years old.")
	fmt.Println("Is it cool?", isCool)
	fmt.Printf("%T\n", name)   // %T prints the type of the variable
	fmt.Printf("%T\n", age)    // %T prints the type of the variable
	fmt.Printf("%T\n", isCool) // %T prints the type of the variable
	fmt.Printf("My friend %v is %v years old.", friend, friend_age)
}
