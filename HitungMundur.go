package main

import "fmt"

func main() {
	var n int

	fmt.Println("Selamat datang!")
	fmt.Println("Ini adalah program recursive")

	n = 5
	i := 1

	//iterasi sama dengan hitung maju
	for i <= n {
		fmt.Println(i)
		i++
	}

	htgMundur(n)
	htgMaju(1, n)

	fmt.Println("exit")
}

//procedure
func htgMundur(i int) {

	if i > 0 {
		fmt.Println(i)
		htgMundur(i - 1)
	}

}

//procedure
func htgMaju(i int, n int) {

	if i <= n {
		fmt.Println(i)
		htgMaju(i+1, n)
	}

}