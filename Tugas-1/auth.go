package main

import "net/http"

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()

		// exception user dan password
		if !ok {
			rw.Write([]byte("username dan password tidak boleh kosong!"))
			return
		}

		// validasi auth admin
		if username == "admin" && password == "admin" {
			next.ServeHTTP(rw, r)
			return
		}

		if username == "editor" && password == "secret" {
			next.ServeHTTP(rw, r)
			return
		}

		if username == "trainer" && password == "rahasia" {
			next.ServeHTTP(rw, r)
			return
		}

		rw.Write([]byte("username dan password salah!"))
	})
}
