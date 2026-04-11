package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter a string to test:")
	s, _ := reader.ReadString('\n')
	input := strings.TrimSpace(strings.ToLower(s))
	if input[0] == 'i' && input[len(input)-1] == 'n' && strings.Contains(input, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
