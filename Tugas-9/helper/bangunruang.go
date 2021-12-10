package helper

type Tabung struct {
	JariJari, Tinggi float64
}

type Balok struct {
	Panjang, Lebar, Tinggi int
}

type HitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}

func (t Tabung) Volume() float64 {
	if int(t.JariJari)%7 == 0 {
		return float64(22 / 7 * t.JariJari * t.JariJari * t.Tinggi)
	} else {
		return float64(3.14 * t.JariJari * t.JariJari * t.Tinggi)
	}
}

func (t Tabung) LuasPermukaan() float64 {
	if int(t.JariJari)%7 == 0 {
		return float64(2 * 22 / 7 * t.JariJari * (t.JariJari + t.Tinggi))
	} else {
		return float64(2 * 3.14 * t.JariJari * (t.JariJari + t.Tinggi))
	}
}

func (b Balok) Volume() float64 {
	return float64(b.Panjang) * float64(b.Lebar) * float64(b.Tinggi)
}

func (b Balok) LuasPermukaan() float64 {
	p := float64(b.Panjang)
	l := float64(b.Lebar)
	t := float64(b.Tinggi)

	return float64((2 * (p + l)) + (2 * (p + t)) + (2 * (l + t)))
}
