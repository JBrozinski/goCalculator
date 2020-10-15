package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("This Calculator was built with Go!")
	cmd := readLine("Enter command: [a]dd, [s]ubtract, [m]ultiply, [d]ivide: ")
	fmt.Print(cmd)
	if cmd == "a" {
		no1, no2 := getUserNumbers()
		result := no1 + no2
		sResult := fmt.Sprintf("%d", result)
		fmt.Println(sResult)
	} else if cmd == "s" {
		no1, no2 := getUserNumbers()
		result := no1 - no2
		sResult := fmt.Sprintf("%d", result)
		fmt.Println(sResult)
	} else if cmd == "m" {
		no1, no2 := getUserNumbers()
		result := no1 * no2
		sResult := fmt.Sprintf("%d", result)
		fmt.Println(sResult)
	} else if cmd == "d" {
		no1, no2 := getUserNumbers()
		result := float32(no1) / float32(no2)
		sResult := fmt.Sprintf("%f", result)
		fmt.Println(sResult)
	} else {
		fmt.Println("Invalid input")
	}
	readLine("press 'enter' to exit")
}

func readLine(message string) string {
	fmt.Print(message)
	var input string
	fmt.Scanln(&input)
	return input
}

func getUserNumbers() (int, int) {
	no1String := readLine("First Number: ")
	no1, _ := strconv.Atoi(no1String)
	no2String := readLine("Second Number: ")
	no2, _ := strconv.Atoi(no2String)
	return no1, no2
}
