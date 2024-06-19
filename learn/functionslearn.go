package learn

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

func main() {

	themeSong()
	goGoGadget("Go")
	printGo(5)

}
