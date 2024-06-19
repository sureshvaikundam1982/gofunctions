package main

import (
	"fmt"
)

func themeSong() {
	fmt.Println("Go Gadget Go")
}

func goGoGadget(x string) {
	fmt.Println("Go go Gadget", x)
}

func printGo(x int) {
	for i := 0; i <= x; i++ {
		fmt.Println("Go")
	}
}

func printGo2(x int, str string) {
	for i := 0; i < x; i++ {
		fmt.Println(str)
	}
}

func arrayeg() {
	var arr [5]int

	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5

	fmt.Println(arr)
	fmt.Println(len(arr))
	fmt.Println(cap(arr))

}

func main() {

	themeSong()
	goGoGadget("Go")
	printGo(5)
	printGo2(6, "suresh")
	arrayeg()

}
