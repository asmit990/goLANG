package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hey, what's your name?");

	reader := bufio.NewReader(os.Stdin)
	name , _ := reader.ReadString('\n')
	fmt.Println("Hello, mr.",name)
}