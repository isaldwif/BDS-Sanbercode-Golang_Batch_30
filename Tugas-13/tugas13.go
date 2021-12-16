package main

//ini ga berjalan sama sekali
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type NilaiMahasiswa struct {
	Nama        string `json:"nama"`
	MataKuliah  string `json:"matakuliah"`
	IndeksNilai string `json:"indeks_nilai"`
	Nilai       uint   `json:"nilai"`
	ID          uint   `json: "id"`
}

var nilaiNilaiMahasiswa = []NilaiMahasiswa{}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()

		// exception user dan password
		if !ok {
			rw.Write([]byte("username dan password tidak boleh kosong!"))
			return
		}

		// validasi auth
		if username == "admin" && password == "admin" {
			next.ServeHTTP(rw, r)
			return
		}

		rw.Write([]byte("username dan password salah!"))
	})
}

//Post
func postMahasiswa(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		var newNilai = NilaiMahasiswa{
			ID: uint(len(nilaiNilaiMahasiswa) + 1),
		}
		if r.Header.Get("Content-Type") == "application/json" {
			decodeJSON := json.NewDecoder(r.Body)
			//by body json
			if err := decodeJSON.Decode(&newNilai); err != nil {
				log.Fatal(err)
			}
		} else {
			//by form
			nama := r.FormValue("nama")
			matkul := r.FormValue("matakuliah")
			getNilai := r.FormValue("nilai")
			nilai, _ := strconv.Atoi(getNilai)

			newNilai.Nama = nama
			newNilai.MataKuliah = matkul
			newNilai.Nilai = uint(nilai)

		}

		if newNilai.Nilai > 100 {
			http.Error(rw, "Nilai tidak boleh diinput lebih dari 100", http.StatusBadRequest)
			return
		}

		//get indeks berdasarkan nilai
		newNilai.IndeksNilai = getIndeksNilai(newNilai.Nilai)

		nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, newNilai)
		dataNilai, _ := json.Marshal(newNilai)
		rw.Write(dataNilai)
		return

	}

	http.Error(rw, "NOT FOUND", http.StatusMethodNotAllowed)
	return

}

//fungsiGET
func getMahasiswa(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		dataNilai, err := json.Marshal(nilaiNilaiMahasiswa)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(dataNilai)
		return
	}
	http.Error(w, "ERROR....", http.StatusNotFound)
}

//fungsi indeks
func getIndeksNilai(nilai uint) string {
	if nilai >= 80 {
		return "A"
	} else if nilai >= 70 {
		return "B"
	} else if nilai >= 60 {
		return "C"
	} else if nilai >= 50 {
		return "D"
	} else if nilai < 50 {
		return "E"
	}
	return ""
}

func main() {
	//method
	http.HandleFunc("/get_mahasiswa", getMahasiswa)
	http.Handle("/post_nilai", Auth(http.HandlerFunc(postMahasiswa)))

	fmt.Println("server running at http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
