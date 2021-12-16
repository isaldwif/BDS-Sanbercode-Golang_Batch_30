package main

import "fmt"

//struct no 1
type buahBuahan struct {
	nama       string
	warna      string
	adaBijinya bool
	harga      int
}

//struct dan method no 2
type segitiga struct {
	alas, tinggi int
}

func (s segitiga) luasSegitiga() {
	fmt.Println(s.alas * s.tinggi / 2)
}

type persegi struct {
	sisi int
}

func (p persegi) luasPersegi() {
	fmt.Println(p.sisi * p.sisi * p.sisi * p.sisi)
}

type persegiPanjang struct {
	panjang, lebar int
}

func (pe persegiPanjang) luasPersegiPanjang() {
	fmt.Println(pe.panjang * pe.lebar)
}

//struct dan method no 3
type phone struct {
	name, brand string
	year        int
	colors      []string
}

func (ph phone) tambahWarna() {

	fmt.Println(ph.brand, ph.name, ph.year, ph.colors)

}

//struck dan fungsi no 4
type movie struct {
	title, genre   string
	duration, year int
}

func tambahDataFilm(title string,
	duration int,
	genre string,
	year int,
	dataFilm *[]movie) {
	*dataFilm = append(*dataFilm, movie{title, genre, duration, year})

}

func main() {

	//ini soal no1
	fmt.Println(" ")
	fmt.Println("=========== soal no 1 ==============")
	satu := buahBuahan{
		nama:       "Nanas",
		warna:      "Kuning",
		adaBijinya: false,
		harga:      9000,
	}

	dua := buahBuahan{
		nama:       "Jeruk",
		warna:      "Oranye",
		adaBijinya: true,
		harga:      8000,
	}

	tiga := buahBuahan{
		nama:       "Semangka",
		warna:      "Hijau & Merah",
		adaBijinya: true,
		harga:      10000,
	}

	empat := buahBuahan{
		nama:       "Pisang",
		warna:      "Kuning",
		adaBijinya: false,
		harga:      5000,
	}

	fmt.Println("nama", "warna", "ada bijinya", "harga")

	fmt.Println(satu.nama, satu.warna, satu.adaBijinya, satu.harga)
	fmt.Println(dua.nama, dua.warna, dua.adaBijinya, dua.harga)
	fmt.Println(tiga.nama, tiga.warna, tiga.adaBijinya, tiga.harga)
	fmt.Println(empat.nama, empat.warna, empat.adaBijinya, empat.harga)

	//ini soal no2
	fmt.Println(" ")
	fmt.Println("=========== soal no 2 ==============")
	var Segitiga = segitiga{
		alas:   10,
		tinggi: 8,
	}

	var Persegi = persegi{
		sisi: 5,
	}

	var Persegip = persegiPanjang{
		panjang: 10,
		lebar:   5,
	}

	Segitiga.luasSegitiga()
	Persegi.luasPersegi()
	Persegip.luasPersegiPanjang()

	//ini soal no3
	fmt.Println(" ")
	fmt.Println("=========== soal no 3 ==============")

	var hanpon = phone{
		name:   "Nokia-XC120",
		brand:  "Nokia",
		year:   1998,
		colors: []string{"merah", "hijau"},
	}

	hanpon.tambahWarna()

	//ini soal no4
	fmt.Println(" ")
	fmt.Println("=========== soal no 4 ==============")

	var dataFilm = []movie{}

	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

}
