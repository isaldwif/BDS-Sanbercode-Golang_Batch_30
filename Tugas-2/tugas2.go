package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	//Soal nomor 1
	fmt.Println("=========== soal no 1 ==============")
	nama1 := "Bootcamp"
	nama2 := "Digital"
	nama3 := "Skill"
	nama4 := "Sanbercode"
	nama5 := "Golang"

	fmt.Println(nama1, nama2, nama3, nama4, nama5)

	//soal nomor 2
	fmt.Println("=========== soal no 2 ==============")
	halo := "Halo Dunia"
	var find = "Dunia"
	var replaceWith = "Golang"

	var newHalo = strings.Replace(halo, find, replaceWith, 1)
	fmt.Println(newHalo)

	//Soal nomor 3
	fmt.Println("=========== soal no 3 ==============")
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"

	var find1 = "r"
	var replaceWith1 = "R"

	var newkatatiga = strings.Replace(kataKetiga, find1, replaceWith1, 1)

	var find2 = "golang"
	var replaceWith2 = "GOLANG"

	var newkataempat = strings.Replace(kataKeempat, find2, replaceWith2, 1)

	fmt.Println(kataPertama, kataKedua, newkatatiga, newkataempat)

	//Soal Nomor 4
	fmt.Println("=========== soal no 4 ==============")
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	newAngka1, _ := strconv.Atoi(angkaPertama)
	newAngka2, _ := strconv.Atoi(angkaKedua)
	newAngka3, _ := strconv.Atoi(angkaKetiga)
	newAngka4, _ := strconv.Atoi(angkaKeempat)

	fmt.Println(newAngka1 + newAngka2 + newAngka3 + newAngka4)

	//Soal nomor 5
	fmt.Println("=========== soal no 5 ==============")
	kalimat := "halo halo bandung"
	angka := 2021

	kalimat = `"` + strings.ReplaceAll(kalimat, "halo", "hi") + `"` + ` -`

	fmt.Println(kalimat, angka)

}
