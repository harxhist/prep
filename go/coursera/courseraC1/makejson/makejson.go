package main

import (
	"encoding/json"
	"fmt"
)


func main(){
	var nameInput string
	var addressInput string
	fmt.Println("Please enter the name:")
	_, err := fmt.Scan(&nameInput);
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Please enter the address:")
	_, err = fmt.Scan(&addressInput);
	if err != nil {
		fmt.Println(err)
	}
	mp := make(map[string]string)
	mp["name"] = nameInput;
	mp["address"] = addressInput;
	barr , err := json.Marshal(mp)
	fmt.Println(string(barr))
}
