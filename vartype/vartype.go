package main

import (
	"fmt"
)

func foo(i interface{}){
	
	switch i.(type){
	case int:
		fmt.Println("I'm an integer!")
	case string:
		fmt.Println("I'm a string!")
	default:
		fmt.Println("Unknown type!")
	}

}

func main(){

var s = int(42)
	foo(s)
}