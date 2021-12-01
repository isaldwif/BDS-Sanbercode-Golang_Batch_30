package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	//Soal nomor 1
	nama1 := "Bootcamp"
	nama2 := "Digital"
	nama3 := "Skill"
	nama4 := "Sanbercode"
	nama5 := "Golang"

	nama := nama1 + " " + nama2 + " " + nama3 + " " + nama4 + " " + nama5

	fmt.Println(nama)

	//soal nomor 2
	halo := "Halo Dunia"
	var find = "Dunia"
	var replaceWith = "Golang"

	var newHalo = strings.Replace(halo, find, replaceWith, 1)
	fmt.Println(newHalo)

	//Soal nomor 3
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

	kataSemua := kataPertama + " " + kataKedua + " " + newkatatiga + " " + newkataempat

	fmt.Println(kataSemua)

	//Soal Nomor 4
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	var num, err = strconv.Atoi(angkaPertama + angkaKedua + angkaKetiga + angkaKeempat)

	if err == nil {
		num = 8 + 5 + 6 + 7
		fmt.Println(num)
	}

	//Soal nomor 5
	kalimat := "halo halo bandung"
	angka := 2021

	var find3 = "halo halo"
	var replaceWith3 = "Hi Hi"

	var newKalimat = strings.Replace(kalimat, find3, replaceWith3, 1)

	fmt.Print(" '' ")
	fmt.Print(newKalimat)
	fmt.Print(" '' ")
	fmt.Print(" - ")
	fmt.Println(angka)

}
