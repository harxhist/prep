package main

import (
    "fmt"
)

func main(){
	var userInput float64
	fmt.Printf("Enter a floating point number\n")
	_, err := fmt.Scan(&userInput)
	if err != nil {
		fmt.Println("Error in reading input")
	}
	ans := int64(userInput)
	fmt.Println(ans)
}
