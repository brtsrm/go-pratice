/*
Parametresi ile aldığı int türden 3 sayıdan ortancasına geri dönen Mid isimli fonksiyonu yazınız ve test ediniz
*/
package main

import "fmt"

func main() {
	runApp()
}
func runApp() {
	var a, b, c int
	fmt.Print("Lütfen 3 değer giriniz")
	fmt.Scan(&a, &b, &c)
	middle := Mid(a, b, c)
	fmt.Printf("%d,%d,%d, %d", a, b, c, middle)
}
func Mid(a, b, c int) int {
	if (a <= b && b <= c) || (c <= b && b <= a) {
		return b
	} else if (b <= a && a <= c) || (c <= a && a <= c) {
		return a
	}
	return c
}
