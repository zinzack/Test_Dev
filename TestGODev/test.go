package main

import (
	"bufio"
	"fmt"
	"os"
)

var chc, chc1 int
var inp []int

func menu1() {
	for chc1 != 2 {
		fmt.Println("\nAdd number ? ")
		fmt.Println("1. Yes")
		fmt.Println("2. No")
		fmt.Print("Choice : ")
		fmt.Scan(&chc1)
		if chc1 == 1 {
			fmt.Println("\nInput number : ")
			fmt.Scanln(&inp)
			bufio.NewReader(os.Stdin).ReadLine()
			fmt.Scanln()
		}

	}

}

func main() {

	for chc == 1 || chc <= 2 {
		fmt.Print("\n\n")
		fmt.Println("Test App")
		fmt.Println("========")
		fmt.Println("1. Input Number")
		fmt.Println("2. View Number(s)")
		fmt.Println("3. Exit")
		fmt.Print("Choice : ")
		fmt.Scan(&chc)

		switch chc {
		case 1:
			menu1()

		case 2:
			if len(inp) != 0 {
				fmt.Print(inp)
			} else {
				fmt.Println("\nNo number")
				bufio.NewReader(os.Stdin).ReadLine()
				fmt.Scanln()

			}

		case 3:
			os.Exit(0)
		}
	} /* else {
		os.Exit(0)
	} */

}
