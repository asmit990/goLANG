package main 

import "fmt"


func main() { 
chanl1 := make(chan string)
chanl2 := make(chan string)   

//if i use the make only send (chan ->string) unidirec
// if i just c want to (-> chan string) 

go sending(chanl1)
valueFromchanel := <- chanl1
fmt.Println("valueFromChanel", valueFromChanel)
}

func sending ( s chan string){
	s <- "go asmit"
}