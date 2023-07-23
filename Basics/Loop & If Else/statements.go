package main

import "fmt"

func main() {
	loop_fun()

	loop_fun2(20)

	switch_fun("mango")
}

func loop_fun() {
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, " is Even")
		} else {
			fmt.Println(i, " is Odd")
		}
	}
}

func loop_fun2(loop_range int) {
	start := 0
	for start < loop_range {
		fmt.Println(start)
		start++
	}
}

func switch_fun(fruit string) {
	switch fruit {
	case "apple":
		fmt.Println("Its yummy apple.")
	case "banana":
		fmt.Println("Its yellow banana.")
	case "mango":
		fmt.Println("Its mangoooo.")
	}
}
