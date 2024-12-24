package main 

import "fmt"

func simpleFunction(){
  fmt.Println("simple")
}

func add(a int ,b int) int{	
	return  a + b	
}

func main() { 
	fmt.Println("yayay")
	simpleFunction()
	ans := add(3, 4)
	fmt.Println("add of two number is ", ans)
}