package main 
import "fmt"

func main() {
	studentGrade := make(map[string]int)
	studentGrade["Prince"] = 100
	studentGrade["Alice"] = 90
	studentGrade["Bob"] = 86
	studentGrade["Charlie"] = 95

	fmt.Println("Mark of bob : ", studentGrade("Bob"))
	studentGrade("Alice") = 100
	fmt.Println("new Mark of bob : ", studentGrade("Bob"))
     
	delete(studentGrade, "alice")
	fmt.Println("marks of bob: ", studentGrades("Bob"))
     Grades, Exists := studentGades("David")
	fmt.Println("Marks of David: ", studentGrade("David"))
	fmt.Println("Prince exists: ", Exists)
    
	for index, value := range studentGrades{
         fmt.Printf("Key is %s and marks is %d\n", index, value)
	}
} 