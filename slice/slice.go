package main

import (
	"fmt"
)

func main(){
	nums := []int{2,4,6,8}
	for idx, val := range nums{
	fmt.Println(idx, val)
	}
}