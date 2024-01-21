package main

import (
	"fmt"
)

func main() {
	runApp()
}
func runApp() {
	var a int
	fmt.Print("Lütfen Sayı giriniz")
	fmt.Scan(&a)
	signum := sigNum(a)
	if signum > 0 {
		fmt.Printf("%d Pozitif", signum)
	} else if signum == 0 {
		fmt.Printf("%d Sıfır", signum)
	} else {
		fmt.Printf("%d Negatif", signum)
	}
}

func sigNum(number int) int {
	if number > 0 {
		return 1
	} else if number == 0 {
		return 0
	}
	return -1
}
