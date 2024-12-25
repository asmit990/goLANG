package main 

import "fmt"

func main() {
	// number := []int{1,2,3,4,5,6}

	// fmt.Println("slice", number)
	// fmt.Println("length", len(number))
	// fmt.Println("Capacity",cap(number))


	numbers:= make([]int, 3,5)
     numbers = append(numbers, 4)
	 numbers = append(numbers, 98)
     numbers = append(numbers, 6)
	fmt.Println("Slice", numbers)
	fmt.Println("Length", len(numbers))
	fmt.Println("Capacity", cap(numbers))
	// capacity seeda 2 hota hai 
	
}