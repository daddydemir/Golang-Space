package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	// string to int
	sayi, _ := strconv.Atoi("101")
	println(sayi + 5)

	// int to string
	word := strconv.Itoa(sayi)
	println(word)

	// int to float
	ondalikli := float64(10.056)
	fmt.Println(ondalikli)

	// float to int
	number := int(ondalikli)
	fmt.Println(number)

	// time comvert
	simdi := time.Now().Unix() // 2022-08-03 09:56:18.6959043 +0300 +03 m=+0.003095301 | Unix 1659509836
	t := time.Unix(simdi, 0)
	fmt.Println(t)

}
