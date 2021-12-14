package main

import (
	"fmt"
	"math"
	"strconv"
	"sync"
	"time"
)

//no1
func cetakHp(hitung int, phones *[]string, wg *sync.WaitGroup) {
	for i, phone := range *phones {
		fmt.Printf("%d. %s\n", i+1, phone)
		time.Sleep(time.Second)
	}
	wg.Done()
}

//no2
func getMovies(ch chan string, mv ...string) {
	for i, value := range mv {
		ch <- strconv.Itoa(i+1) + ". " + value
	}

	close(ch)
}

//no3
func luasLigkaran(luas chan float64, jariJari float64) {
	nilaiLuas := math.Pi * jariJari * jariJari
	luas <- nilaiLuas
}

func kelilingLingkaran(keliling chan float64, jariJari float64) {
	nilaiKeliling := 2 * math.Pi * jariJari
	keliling <- nilaiKeliling
}

func volumeTabung(volume chan float64, jariJari float64, tinggi float64) {
	nilaiVolume := math.Pi * jariJari * jariJari * tinggi
	volume <- nilaiVolume
}

func luasPersegiPanjang(luas chan int, panjang int, lebar int) {
	nilaiLuas := panjang * lebar
	luas <- nilaiLuas
}

//no4
func kelilingPersegiPanjang(keliling chan int, panjang int, lebar int) {
	nilaiKeliling := 2 * (panjang + lebar)
	keliling <- nilaiKeliling
}

func volumeBalok(volume chan int, panjang int, lebar int, tinggi int) {
	nilaiVolume := panjang * lebar * tinggi
	volume <- nilaiVolume
}

func main() {
	//Soal nomor 1
	fmt.Println(" ")
	fmt.Println("[=========== soal no 1 ==============]")

	var phones = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}
	var wg sync.WaitGroup

	wg.Add(1)
	go cetakHp(0, &phones, &wg)

	wg.Wait()

	//Soal nomor 2
	fmt.Println(" ")
	fmt.Println("[=========== soal no 2 ==============]")

	var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}

	moviesChannel := make(chan string)

	go getMovies(moviesChannel, movies...)

	for value := range moviesChannel {
		fmt.Println(value)
	}

	//Soal nomor 3
	fmt.Println(" ")
	fmt.Println("[=========== soal no 3 ==============]")

	luasLing := make(chan float64)
	kelilingLing := make(chan float64)
	volumeTabung1 := make(chan float64)

	go luasLigkaran(luasLing, 8)
	nilaiLuas := <-luasLing
	fmt.Println("Luas lingkaran :", nilaiLuas)

	go kelilingLingkaran(kelilingLing, 14)
	nilaiKeliling := <-kelilingLing
	fmt.Println("Keliling lingkaran :", nilaiKeliling)

	go volumeTabung(volumeTabung1, 20, 10)
	nilaiVolume := <-volumeTabung1
	fmt.Println("Volume tabung :", nilaiVolume)

	//Soal nomor 4
	fmt.Println(" ")
	fmt.Println("[=========== soal no 4 ==============]")

	luasPersegipanjang := make(chan int)
	kelilingPersegipanjang := make(chan int)
	volumePersegipanjang := make(chan int)

	go luasPersegiPanjang(luasPersegipanjang, 10, 11)
	go kelilingPersegiPanjang(kelilingPersegipanjang, 6, 7)
	go volumeBalok(volumePersegipanjang, 10, 11, 12)

	for i := 0; i < 3; i++ {
		select {
		case nilaiLuas := <-luasPersegipanjang:
			fmt.Println("Luas persegi panjang :", nilaiLuas)
		case nilaiKeliling := <-kelilingPersegipanjang:
			fmt.Println("Keliling persegi panjang :", nilaiKeliling)
		case nilaiVolume := <-volumePersegipanjang:
			fmt.Println("Volume tabung :", nilaiVolume)
		}
	}
}
