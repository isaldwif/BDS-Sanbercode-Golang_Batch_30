package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

//soal no 1
func iniKalimat() {
	kalimat := "Golang Backend Development"
	tahun := 2021
	fmt.Println(kalimat, tahun)
}

func runKalimat() {
	defer iniKalimat()
}

//soal no 2
func getPanic() {
	result := recover()
	fmt.Println(result)
}

func handlePanic(message string) {
	panic(message)
}

func kelilingSegitigaSamaSisi(sisi uint, showText bool) interface{} {

	defer getPanic()
	switch {

	case sisi > 0 && showText:
		return "keliling segitiga sama sisinya dengan sisi " + strconv.Itoa(int(sisi)) + " cm adalah " + strconv.Itoa(int(sisi*3)) + " cm"
	case sisi > 0 && !showText:
		return `"` + strconv.Itoa(int(sisi*3)) + `"`
	case sisi == 0 && showText:
		return "Maaf anda belum menginput sisi dari segitiga sama sisi"
	case sisi == 0 && !showText:
		handlePanic("Maaf anda belum menginput sisi dari segitiga sama sisi")
	default:
		return nil
	}
	return " "
}

//soal3
func tambahAngka(angka int, buatAngka *int) {

	*buatAngka = *buatAngka + angka
}

func cetakAngka(buatAngka *int) {
	fmt.Println("=========== soal no 3 ==============", "\n", *buatAngka)
}

//soal4
func tambahHp(nama string, phones *[]string) {
	*phones = append(*phones, nama)
}

//soal5
func luas(jari_jari float64, luasLigkaran *float64) {
	*luasLigkaran = math.Phi * jari_jari * jari_jari
	return
}

func keliling(jari_jari float64, kelilingLigkaran *float64) {
	*kelilingLigkaran = math.Phi * jari_jari * 2
	return
}

func main() {

	//soal no 1
	fmt.Println(" ")
	fmt.Println("=========== soal no 1 ==============")
	//ini akan tampildibawah
	runKalimat()

	//soal no 2
	fmt.Println(" ")
	fmt.Println("=========== soal no 2 ==============")

	fmt.Println(kelilingSegitigaSamaSisi(4, true))
	fmt.Println(kelilingSegitigaSamaSisi(8, false))
	fmt.Println(kelilingSegitigaSamaSisi(0, true))
	fmt.Println(kelilingSegitigaSamaSisi(0, false))

	//soal no 3
	angka := 1
	defer cetakAngka(&angka)
	fmt.Println(" ")
	// deklarasi variabel angka ini simpan di baris pertama func main
	//isi akan ada dipaling bawah deng

	tambahAngka(7, &angka)

	tambahAngka(6, &angka)

	tambahAngka(-1, &angka)

	tambahAngka(9, &angka)

	//soal no 4
	fmt.Println(" ")
	fmt.Println("=========== soal no 4 ==============")
	time.Sleep(time.Second * 1)

	var phones = []string{}

	tambahHp("Xiaomi", &phones)
	tambahHp("Asus", &phones)
	tambahHp("Iphone", &phones)
	tambahHp("Oppo", &phones)
	tambahHp("Realme", &phones)
	tambahHp("Vivo", &phones)

	for i, item := range phones {
		time.Sleep(time.Second * 1)
		fmt.Print(i + 1)
		fmt.Println(".", item)
		fmt.Println()
	}

	//soal no 5
	fmt.Println(" ")
	fmt.Println("=========== soal no 5 ==============")
	var luasLigkaran float64
	var kelilingLingkaran float64

	luas(7.10, &luasLigkaran)
	keliling(10, &kelilingLingkaran)
	fmt.Println("luas lingkaran adalah :", luasLigkaran, "cm")
	fmt.Println("keliling lingkaran adalah :", kelilingLingkaran, "cm")

}
