package helper

import "strconv"

//struct
type SegitigaSamaSisi struct {
	Alas, Tinggi int
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

//interface
type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

//func
func (s SegitigaSamaSisi) Luas() int {
	return s.Alas * s.Tinggi / 2
}

func (s SegitigaSamaSisi) Keliling() int {
	return s.Alas * 3
}

func (p PersegiPanjang) Luas() int {
	return p.Panjang * p.Lebar
}

func (p PersegiPanjang) Keliling() int {
	return 2 * (p.Panjang + p.Lebar)
}

func LuasPersegi(Sisi uint, ShowText bool) interface{} {
	switch {
	case Sisi > 0 && ShowText:
		return "Luas persegi dengan Sisi " + strconv.Itoa(int(Sisi)) + " cm adalah " + strconv.Itoa(int(Sisi*Sisi)) + " cm"
	case Sisi > 0 && !ShowText:
		return Sisi * Sisi
	case Sisi == 0 && ShowText:
		return "Maaf anda belum menginput Sisi dari persegi"
	default:
		return nil
	}
}
