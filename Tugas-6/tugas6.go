package main

import "fmt"

//fungsi nomor 1
func luas(phi float64, jari_jari float64, luasLigkaran *float64) {
	*luasLigkaran = phi * jari_jari * jari_jari
	return
}

func keliling(phi float64, jari_jari float64, kelilingLigkaran *float64) {
	*kelilingLigkaran = phi * jari_jari * 2
	return
}

//fungsi nomor 2
func introduce(sentence *string, nama string, jenis string, pekerjaan string, umur string) {

	if jenis == "laki-laki" {
		*sentence = "Pak " + nama + " adalah seorang " + pekerjaan + " yang berusia " + umur + " tahun"
	} else {
		*sentence = "Bu " + nama + " adalah seorang " + pekerjaan + " yang berusia " + umur + " tahun"
	}

	return
}

//fungsi nomor 3
func buahFavorit(buah *[]string, nama_buah ...string) {
	*buah = append(*buah, nama_buah...)

	for i, kumpul := range *buah {
		fmt.Print(i + 1)
		fmt.Println(".", kumpul)
	}

}

//fungsi no 4
func tambahDataFilm(title string, durasi string, genre string, tahun string, dataFilm *[]map[string]string) {
	*dataFilm = append(*dataFilm, map[string]string{
		"genre": genre,
		"jam":   durasi,
		"tahun": tahun,
		"title": title,
	})

}

func main() {

	//ini soal no1
	fmt.Println(" ")
	fmt.Println("=========== soal no 1 ==============")
	var luasLigkaran float64
	var kelilingLingkaran float64

	luas(3.14, 14, &luasLigkaran)
	keliling(3.14, 14, &kelilingLingkaran)

	fmt.Println("luas lingkaran adalah :", luasLigkaran, "cm")
	fmt.Println("keliling lingkaran adalah :", kelilingLingkaran, "cm")

	//ini soal no2
	fmt.Println(" ")
	fmt.Println("=========== soal no 2 ==============")
	var sentence string
	introduce(&sentence, "John", "laki-laki", "penulis", "30")

	fmt.Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"
	introduce(&sentence, "Sarah", "perempuan", "model", "28")

	fmt.Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	//ini soal no3
	fmt.Println(" ")
	fmt.Println("=========== soal no 3 ==============")
	var buah = []string{}

	buahFavorit(&buah, "Jeruk", "Semangka", "Mangga", "Stawberry", "Durian", "Manggis", "Alpukat")

	// ini soal no4
	fmt.Println(" ")
	fmt.Println("=========== soal no 4 ==============")
	var dataFilm = []map[string]string{}

	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	//fungsinya ada di atas
	for i, item := range dataFilm {
		fmt.Print(i + 1)
		fmt.Print(". ")
		fmt.Println("title :", item["genre"])
		fmt.Println("  ", "duration :", item["jam"])
		fmt.Println("  ", "tahun :", item["tahun"])
		fmt.Println("  ", "title :", item["title"])
	}

}
