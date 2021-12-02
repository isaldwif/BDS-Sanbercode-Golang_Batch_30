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

	switch bulan {
	case 1:
		fmt.Println("Januari")
	case 2:
		fmt.Println("Februari")
	case 3:
		fmt.Println("Maret")
	case 4:
		fmt.Println("April")
	case 5:
		fmt.Println("Mei")
	case 6:
		fmt.Println("Juni")
	case 7:
		fmt.Println("Juli")
	case 8:
		fmt.Println("Agustus")
	case 9:
		fmt.Println("September")
	case 10:
		fmt.Println("Oktober")
	case 11:
		fmt.Println("November")
	case 12:
		fmt.Println("Desember")
	}

	fmt.Println(tanggal, bulan, tahun)

}
