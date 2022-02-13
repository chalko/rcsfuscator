package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	message, _ := reader.ReadString('\n')
	fmt.Println("Hello, world.")
	fmt.Println(message)

}
