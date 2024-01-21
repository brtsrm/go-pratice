package main

import (
	"fmt"
)

/*
Klavyeden alınan int türden üç sayı arasındaki büyüklük-küçüklük ilişkisini küçükten büyüğe doğru < ve =
sembolleriyle gösteren programı yazınız:
*/
func main() {

	var a1, a2, a3 int
	fmt.Println("3 sayı giriniz")
	fmt.Scan(&a1, &a2, &a3)
	min, mid, max := a1, a2, a3
	if a1 > a2 {
		min, mid = a2, a1
	}
	if mid > a3 {
		mid, max = a3, mid
	}
	if min > mid {
		min, mid = mid, min
	}
	if mid > max {
		mid, max = max, mid
	}
	if min == mid && mid == max {
		fmt.Printf("%d = %d = %d", min, mid, max)
	} else if min < mid && mid == max {
		fmt.Printf("%d < %d = %d", min, mid, max)
	} else if min == mid && mid < max {
		fmt.Printf("%d = %d < %d", min, mid, max)
	} else {
		fmt.Printf("%d < %d < %d", min, mid, max)
	}
}
