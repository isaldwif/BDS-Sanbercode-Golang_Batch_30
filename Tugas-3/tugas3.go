package main

import (
	"fmt"
	"strconv"
)

func main() {

	//Soal nomor 1
	fmt.Println("=========== soal no 1 ==============")
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"

	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	panPersegi, _ := strconv.Atoi(panjangPersegiPanjang)
	lebPersegi, _ := strconv.Atoi(lebarPersegiPanjang)
	alSegi, _ := strconv.Atoi(alasSegitiga)
	tingSegi, _ := strconv.Atoi(tinggiSegitiga)

	//rumus keliling Persegi panjang 2p*2l
	//rumus luas persegi panjang p*l
	//rumus luas segitiga  a*t/2

	fmt.Println("---- Luas Persegi panjang ----")
	var luasPersegiPanjang int = panPersegi * lebPersegi
	fmt.Println(luasPersegiPanjang)

	fmt.Println("---- Keliling Persegi panjang ----")
	var kelilingPersegiPanjang int = 2*panPersegi + 2*lebPersegi
	fmt.Println(kelilingPersegiPanjang)

	fmt.Println("---- Luas Segitiga ----")
	var luasSegitiga int = alSegi * tingSegi / 2
	fmt.Println(luasSegitiga)

	//Soal nomor 2
	fmt.Println("=========== soal no 2 ==============")
	var nilaiJohn = 80
	var nilaiDoe = 50

	fmt.Println("---- Niliai John ----")
	fmt.Println("nilai john adalah : ", nilaiJohn)
	fmt.Print("maka nilai dia mendakapatkan ")
	if nilaiJohn >= 80 {
		fmt.Println("A")
	} else if nilaiJohn >= 70 {
		fmt.Println("B")
	} else if nilaiJohn >= 60 {
		fmt.Println("C")
	} else if nilaiJohn >= 50 {
		fmt.Println("D")
	} else if nilaiJohn < 50 {
		fmt.Println("E")
	}

	fmt.Println("---- Niliai Doe ----")
	fmt.Println("nilai Doe adalah : ", nilaiDoe)
	fmt.Print("maka nilai dia mendakapatkan ")
	if nilaiDoe >= 80 {
		fmt.Println("A")
	} else if nilaiDoe >= 70 {
		fmt.Println("B")
	} else if nilaiDoe >= 60 {
		fmt.Println("C")
	} else if nilaiDoe >= 50 {
		fmt.Println("D")
	} else if nilaiDoe < 50 {
		fmt.Println("E")
	}

	//Soal nomor 3
	fmt.Println("=========== soal no 3 ==============")
	var tanggal = 5
	var bulan = 10
	var tahun = 1998

	fmt.Print(tanggal, " ")

	switch bulan {
	case 1:
		fmt.Print("Januari ")
	case 2:
		fmt.Print("Februari ")
	case 3:
		fmt.Print("Maret ")
	case 4:
		fmt.Print("April ")
	case 5:
		fmt.Print("Mei")
	case 6:
		fmt.Print("Juni ")
	case 7:
		fmt.Print("Juli ")
	case 8:
		fmt.Print("Agustus ")
	case 9:
		fmt.Print("September ")
	case 10:
		fmt.Print("Oktober ")
	case 11:
		fmt.Print("November ")
	case 12:
		fmt.Print("Desember ")
	}

	fmt.Println(tahun)

}
