package main

import (
	"fmt"
)

func main() {

	//Soal nomor 1
	fmt.Println("[=========== soal no 1 ==============]")
	for i := 1; i <= 20; i++ {

		if i%2 == 0 {
			fmt.Println(i, "- Berkualitas")
		} else if i%3 == 0 {
			fmt.Println(i, "- I Love Coding")
		} else {
			fmt.Println(i, "- Santai")
		}
	}
	fmt.Printf("\n")

	//Soal nomor 2
	fmt.Println("[=========== soal no 2 ==============]")

	for i := 1; i <= 7; i++ {
		for j := i; j >= 1; j-- {
			fmt.Printf("#")
		}
		fmt.Printf("\n")
	}

	fmt.Printf("\n")

	//Soal nomor 3
	fmt.Println("[=========== soal no 3 ==============]")
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}

	slice := kalimat[2:]

	fmt.Println(slice)
	fmt.Printf("\n")

	//Soal nomor 4
	fmt.Println("[=========== soal no 4 ==============]")
	var sayuran = []string{}

	slice = append(sayuran, " ", "Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun")

	for i, sayur := range slice {
		fmt.Printf("%d. %s\n", i, sayur)
	}
	fmt.Printf("\n")

	//Soal nomor 5
	fmt.Println("[=========== soal no 5 ==============]")
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}

	satuan["Volume balok"] = 168

	for key, element := range satuan {
		fmt.Println(key, " = ", element)
	}

}
