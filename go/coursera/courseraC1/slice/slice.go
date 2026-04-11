package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main(){
	s := make([]int, 0, 3)
	var i string
	for i != "X" {
		_ , err := fmt.Scan(&i)
		if(err != nil){
			fmt.Println("issue in scan")
		}
		var n int
		if i == "X" {
			break
		}
		n, err = strconv.Atoi(i);
		if(err != nil){
			fmt.Println("error in atoi")
		}
		s = append(s, n)
		sort.Ints(s)
		fmt.Println(s);
	}
}