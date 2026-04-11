package main

import (
	"bufio"
	"os"
	"fmt"
	"strings"
	"strconv"
)

func Swap(slice []int, i int){
	slice[i], slice[i+1] = slice[i+1], slice[i]
}

func BubbleSort(slice []int){
	n := len(slice);
	for i := 0; i < n - 1 ; i++ {
		swapped := false

		for j := 0; j < n - i - 1 ; j++ {
			if slice[j] > slice[j+1] {
				Swap(slice, j);
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func ConvertStringToIntSlice(input []string) []int{
	slice := make([]int, 0)
	for _, v := range input {
		i, err := strconv.Atoi(v)
		if err != nil{
			fmt.Println("Error in converting string to interger", err)
		}
		slice = append(slice, i)
	}
	return slice
} 

func ReadStandardInput() []int {
	fmt.Println("Enter space-seprated sequence of up to 10 integers to sort")
	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')
	if err != nil{
		fmt.Printf("Error in reading input", err)
	}
	inputString := strings.TrimSpace(s);
	input := strings.Split(inputString, " ")
	slice := ConvertStringToIntSlice(input)
	return slice
}

func main(){
	input := ReadStandardInput();
	BubbleSort(input)
	fmt.Println(input)
}