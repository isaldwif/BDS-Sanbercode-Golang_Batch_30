package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
)

//no 1
type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

func (s segitigaSamaSisi) luas() int {
	return s.alas * s.tinggi / 2
}

func (s segitigaSamaSisi) keliling() int {
	return s.alas * s.alas * s.alas
}

func (p persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (p persegiPanjang) keliling() int {
	return 2 * (p.panjang + p.lebar)
}

func (t tabung) volume() float64 {
	return math.Phi * (t.jariJari * t.jariJari) * t.tinggi
}

func (t tabung) luasPermukaan() float64 {
	//L = 2 × π × r × (r + t)
	return 2 * math.Phi * t.jariJari * (t.jariJari * t.tinggi)
}

func (b balok) volume() int {
	return b.panjang * b.lebar * b.tinggi
}

func (b balok) luasPermukaan() int {
	//L = 2 × (p.l + p.t +l.t)
	return 2 * (b.panjang*b.lebar + b.panjang*b.tinggi + b.lebar*b.tinggi)
}

//no2
type merekHp interface {
	panggilHp() string
}

type phone struct {
	name, brand string
	year        int
	colors      []string
}

func (po phone) panggilHp() string {

	for _, s := range po.colors {
		return "name : " + po.name + "\n" + "brand : " + po.brand + "\n" + "Year : " + strconv.Itoa(po.year) + "\n" + "Colors : " + s
	}
	return ""
}

//no3
func iniSwitch(v interface{}) (int, error) {
	switch u := v.(type) {
	case int:
		return u, nil
	case float64:
		return int(u), nil
	case string:
		return strconv(u)
	default:
		return 0, errors.New("Unsupported type")
	}
}

func main() {

	//ini soal no1
	fmt.Println(" ")
	fmt.Println("=========== soal no 1 ==============")

	var bangunDatar hitungBangunDatar

	fmt.Println("===== Segitiga Sama Sisi")
	bangunDatar = segitigaSamaSisi{
		alas:   5,
		tinggi: 5,
	}
	fmt.Println("luas      :", bangunDatar.luas(), "cm")
	fmt.Println("keliling  :", bangunDatar.keliling(), "cm")

	fmt.Println("===== Persegi Panjang")
	bangunDatar = persegiPanjang{
		panjang: 10,
		lebar:   5,
	}
	fmt.Println("luas      :", bangunDatar.luas(), "cm")
	fmt.Println("keliling  :", bangunDatar.keliling(), "cm")

	//ini soal no2
	fmt.Println(" ")
	fmt.Println("=========== soal no 2 ==============")

	var counterHp merekHp

	counterHp = phone{
		name:   "Samsung Galaxy Note 20",
		brand:  "Samsung Galaxy Note 20",
		year:   2020,
		colors: []string{"Mystic Bronze ,Mystic White, Mystic Black"},
	}

	fmt.Printf(counterHp.panggilHp())

	//ini soal no3
	fmt.Println("\n")
	fmt.Println("=========== soal no 3 ==============")
	var iniSwitch interface{}

}
