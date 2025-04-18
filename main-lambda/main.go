package main

import (
	"fmt"
	"common-libs/utility"
)

func main(){
	fmt.Println(Handler(3));
}

func Handler(x int) (y int) {
	utility.CPrint()
	return x;
}