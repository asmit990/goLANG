package main 

import "fmt"

func main() {
	for i:=0 ; i < 10 ; i++ {
		fmt.Println( i)
	}

	// counter := 0 
	// for{
	// 	fmt.Println("Infinite Loop")
	// 	if count == 3 {
	// 		break
	// 	}
	// }
	numbers := []int{1, 2, 3, 4, 5}
for index, value := range numbers {
    fmt.Printf("Index of number is %d, and value is %d\n", index, value)
}

}
	 

