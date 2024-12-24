package main 
import "fmt"

func main() {
	fmt.Println("we are learning array in goLang ")

	var name[5]string

	name[0] = "prince"
	name[2] = "akash"

	fmt.Println("Name of Person is : ", name)

	var numbers = [5]int{1,2,3,4,5}
	fmt.Println("Number is : ", numbers)
     
	fmt.Println("length of Number array is : ", len(numbers))
	fmt.Println("value of name at 2 index is ; ", name(2))
   fmt.Println("name array is %q\n", name)
	// go make the array default initialization with
	// wapas kuch khaarab hai  
}   