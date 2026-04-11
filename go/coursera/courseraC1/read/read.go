package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct{
	fname string
	lname string
}


func main(){
	fmt.Println("Please enter the filename to read from:")
	var filename string
	_, err := fmt.Scan(&filename)
	if err != nil{
		fmt.Println("Error scaning filename", err)
	}
	file, err := os.Open(filename);
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	slice := make([]Name, 1)

	for scanner.Scan() {
		line := scanner.Text()
		var personName Name
		name := strings.Split(line, " ");
		personName.fname = name[0]
		personName.lname = name[1]
		slice = append(slice, personName)
	}

	if err := scanner.Err(); err != nil{
		fmt.Println("Error reading file:", err)
	}

	for _, v := range slice {
		fmt.Println(v.fname, v.lname)
	}
}