//models/movie.go, digunakan untuk membuat struct/ struktur.

package models

import "time"

type (
	NilaiMahasiswa struct {
		Nama        string    `json:"nama"`
		MataKuliah  string    `json:"matakuliah"`
		IndeksNilai string    `json:"indeks_nilai"`
		Nilai       uint      `json:"nilai"`
		ID          uint      `json: "id"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
)
