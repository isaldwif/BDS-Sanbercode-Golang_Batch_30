package main

import (
	"Tugas-9/helper"
	"fmt"
	"strconv"
)

func main() {
	// soal 1
	fmt.Println("------SOAL 1-------")
	var segitiga1 helper.HitungBangunDatar = helper.SegitigaSamaSisi{Alas: 7, Tinggi: 8}
	var persegiPanjang1 helper.HitungBangunDatar = helper.PersegiPanjang{Panjang: 7, Lebar: 8}

	fmt.Println("------bangun datar-------")
	fmt.Println("luas segitiga sama sisi dari Alas 7 dan tinggi 8 adalah", segitiga1.Luas())
	fmt.Println("Keliling segitiga sama sisi dari alas 7 dan tinggi 8 adalah", segitiga1.Keliling())

	fmt.Println("Luas persegi panjang sama dari panjang 7 dan lebar 8 adalah", persegiPanjang1.Luas())
	fmt.Println("Keliling persegi panjang sama dari panjang 7 dan lebar adalah", persegiPanjang1.Keliling())

	fmt.Println("------bangun ruang-------")

	var tabung1 helper.HitungBangunRuang = helper.Tabung{JariJari: 7, Tinggi: 10}
	var balok1 helper.HitungBangunRuang = helper.Balok{Panjang: 7, Lebar: 5, Tinggi: 6}

	fmt.Println("volume tabung dari jari-jari 7 dan tinggi 10 adalah", tabung1.Volume())
	fmt.Println("Luas permukaan tabung dari jari-jari 7 dan tinggi 10 adalah", tabung1.LuasPermukaan())

	fmt.Println("Volume balok dari Panjang 7, Lebar 5 dan tinggi 6 adalah", balok1.Volume())
	fmt.Println("Luas permukaan balok dari Panjang 7, lebar 5 dan tinggi 6 adalah", balok1.LuasPermukaan())

	// soal 2
	fmt.Println(" ")
	fmt.Println("------SOAL 2-------")
	var samsung helper.Description = helper.Phone{
		Name:   "Samsung Galaxy Note 20",
		Brand:  "Samsung",
		Year:   2020,
		Colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}

	fmt.Println(samsung.GetDescription())

	// soal 3
	fmt.Println(" ")
	fmt.Println("------SOAL 3-------")
	fmt.Println(helper.LuasPersegi(4, true))

	fmt.Println(helper.LuasPersegi(8, false))

	fmt.Println(helper.LuasPersegi(0, true))

	fmt.Println(helper.LuasPersegi(0, false))

	// soal 4
	fmt.Println(" ")
	fmt.Println("------SOAL 4-------")

	kumpulanAngka := helper.KumpulanAngkaPertama.([]int)

	Kalimat := helper.Prefix.(string)

	angkaKedua := helper.KumpulanAngkaKedua.([]int)
	kumpulanAngka = append(kumpulanAngka, angkaKedua...)

	for index, item := range kumpulanAngka {
		helper.Jumlah += item
		if index == len(kumpulanAngka)-1 {
			Kalimat += strconv.Itoa(item) + "=" + strconv.Itoa(helper.Jumlah)
		} else {
			Kalimat += strconv.Itoa(item) + "+"
		}
	}

	fmt.Println(Kalimat)

}
