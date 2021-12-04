package main

import "fmt"

//FUNGSI Soal nomor 1
func luasPersegiPanjang(panjang int, lebar int) (rumusLuas int) {
	rumusLuas = panjang * lebar
	return
}

func kelilingPersegiPanjang(panjang int, lebar int) (rumusKeliling int) {
	rumusKeliling = 2 * (panjang + lebar)
	return
}

func volumeBalok(panjang int, lebar int, tinggi int) (rumusVolume int) {
	rumusVolume = panjang * lebar * tinggi
	return
}

//FUNGSI Soal nomor 2
func introduce(nama string, jenis string, pekerjaan string, umur string) string {
	perkenalan := nama + " adalah seorang " + pekerjaan + " yang berusia " + umur + " tahun"

	if jenis == "laki-laki" {
		fmt.Println("Pak " + perkenalan)
	} else {
		fmt.Println("Bu " + perkenalan)
	}

	return ""
}

//FUNGSI Soal nomor 3
func buahFavorit(nama string, kalimat ...string) (kalimat1 string) {
	kalimat1 = "halo nama saya " + nama + " dan buah favorit saya adalah "

	for _, kumpul := range kalimat {
		kalimat1 = kalimat1 + `"` + kumpul + `"` + " "
	}
	return
}

func main() {

	//soal no 1
	fmt.Println("\n")
	fmt.Println("=========== soal no 1 ==============")
	panjang := 12
	lebar := 4
	tinggi := 8

	Luas := luasPersegiPanjang(panjang, lebar)
	Keliling := kelilingPersegiPanjang(panjang, lebar)
	Volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println("Luas persegi panjang adalah : ", Luas, "cm")
	fmt.Println("Keliling persegi panjang adalah : ", Keliling, "cm")
	fmt.Println("Volume Balok adalah : ", Volume, "cm")

	//soal no 2
	fmt.Println("\n")
	fmt.Println("=========== soal no 2 ==============")
	john := introduce("john", "laki-laki", "penulis", "30")

	sarah := introduce("sarah", "perempuan", "model", "29")

	fmt.Println(john)
	fmt.Println(sarah)

	//soal no 3
	fmt.Println("=========== soal no 3 ==============")

	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}

	var buahFavoritJohn = buahFavorit("John", buah...)

	fmt.Println(buahFavoritJohn)

	//soal no4
	fmt.Println("\n")
	fmt.Println("=========== soal no 4 ==============")

	var dataFilm = []map[string]string{}
	// closure function
	tambahDataFilm := func(title string, durasi string, genre string, tahun string) {
		dataFilm = append(dataFilm, map[string]string{"genre": genre, "jam": durasi, "tahun": tahun, "title": title})

	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}

}
