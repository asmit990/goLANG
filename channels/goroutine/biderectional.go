package main 

import "fmt"


func main() { 
chanl1 := make(chan string)
chanl2 := make(chan string) 

//if i use the make only send (chan ->string) unidirec
// if i just c want to (-> chan string) 

go sending(chanl1)
valueFromChanel := <- chanl1
fmt.Println("valueFromChanel", valueFromChanel)
go recieving(chanl2)

chanl2 <- valueFromChanel


chanl1 := make(chan string)
go convert(chanl1)
fmt.Println(<-chanl1)
}
func convert(s chan<- string){
	s <- "some value"
}
func sending(s chan string){
	s <- "go asmit"
}
func recieving(s chan string){
	fmt.Println(<-s)

}
func multiplyWithChanel(ch chan int){

}