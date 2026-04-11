package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func GenDisplaceFn(acceleration, velocity, initialDisplacement float64) func(int) float64 {
	return func (t int) float64 {
		displacement := (0.5) * acceleration * float64((t * t)) + (velocity * float64(t)) + initialDisplacement;
		return displacement;
	}
}

func handleInput() (float64, float64, float64, int){
	fmt.Println("Please enter acceleration, initial velocity, and initial displacement")
	reader := bufio.NewReader(os.Stdin);
	stringInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Error in reading input %s", err);
	}
	inputSlice := strings.Split(stringInput, " ")
	var acceleration, velocity, initialDisplacement float64
	acceleration, _ = strconv.ParseFloat(inputSlice[0], 64)
	velocity, _ = strconv.ParseFloat(inputSlice[1], 64)
	initialDisplacement, _ = strconv.ParseFloat(inputSlice[2], 64)
	fmt.Println("Please enter time")
	var time int
	_, err = fmt.Scan(&time)
	if err != nil {
		fmt.Println("Error in scanning time", err)
	}
	return acceleration, velocity, initialDisplacement, time
}

func main(){
	acceleration, velocity, initialDisplacement, time := handleInput()
	fn := GenDisplaceFn(acceleration, velocity, initialDisplacement)
	fmt.Println(fn(time))
}