package main

import (
	"fmt"
	"workbench/demo/sortme"
)

//var myArray = []int{5}

var sortMe = []int{99, 23, 22, 25, 29, 150, 1, 7, 6}

func main() {
	sortme.Bubble_sort(sortMe)
	fmt.Println(sortMe)
}
