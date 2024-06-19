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

func main() {

	themeSong()
	goGoGadget("Go")
	printGo(5)
	printGo2(6, "suresh")

}
