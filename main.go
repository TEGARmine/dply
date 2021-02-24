package main

import (
	"fmt"
	"net/http"
)

func about(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("halaman about"))
}
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Selamat Datang")
	})

	http.HandleFunc("/about", about)

	http.ListenAndServe(":3000", nil)

}
