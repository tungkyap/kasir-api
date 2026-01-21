package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Produk struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	Harga int    `json:"harga"`
	Stok  int    `json:"stok"`
}

var produk = []Produk{
	{ID: 1, Nama: "Indomie Godog", Harga: 3500, Stok: 10},
	{ID: 2, Nama: "Aquviva", Harga: 3000, Stok: 40},
	{ID: 3, Nama: "Kecap", Harga: 12000, Stok: 20},
}

var lastID = 3

// GET localhost:8080/api/produk/{id}
func getProdukByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	// Merubah string ke int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	for _, p := range produk {
		if p.ID == id {
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	http.Error(w, "Produk belum ada", http.StatusNotFound)
}

// PUT localhost:8080/api/produk/{id}
func updateProduk(w http.ResponseWriter, r *http.Request) {
	// dapat get id dari request
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	// ganti string ke int
	id, err := strconv.Atoi(idStr) // strconv.Atoi() function sama kayak Number() di JS
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}
	// get data dari request
	var updateProduk Produk
	err = json.NewDecoder(r.Body).Decode(&updateProduk)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	// loop produk, cari id, ganti sesuai data dari request
	for i := range produk {
		if produk[i].ID == id {
			updateProduk.ID = id
			produk[i] = updateProduk
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updateProduk)
			return
		}
	}
	http.Error(w, "Produk belum ada", http.StatusNotFound)
}

// DELETE localhost:8080/api/produk/{id}
func deleteProduk(w http.ResponseWriter, r *http.Request) {
	// get id
	idStr := strings.TrimPrefix(r.URL.Path, "/api/produk/")
	// ganti id dari string ke int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	// loop produk cari id, dapat index yang dihapus
	for i, p := range produk {
		if p.ID == id {
			// bikin array slice baru dengan data sebelum dan sesudah index
			produk = append(produk[:i], produk[i+1:]...)
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "sukses delete",
			})
			return
		}
	}
}

func main() {
	// GET localhost:8080/api/produk/{id}
	// PUT localhost:8080/api/produk/{id}
	// DELETE localhost:8080/api/produk/{id}
	http.HandleFunc("/api/produk/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getProdukByID(w, r)
		} else if r.Method == "PUT" {
			updateProduk(w, r)
		} else if r.Method == "DELETE" {
			deleteProduk(w, r)
		}
	})

	// GET localhost:8080/api/produk
	// POST localhost:8080/api/produk
	http.HandleFunc("/api/produk", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(produk)
		} else if r.Method == "POST" {
			// baca dari request data
			var produkBaru Produk
			err := json.NewDecoder(r.Body).Decode(&produkBaru)
			if err != nil {
				http.Error(w, "Invalid request", http.StatusBadRequest)
				return
			}
			// masukan data ke dalam variabel produk
			lastID++
			produkBaru.ID = lastID
			produk = append(produk, produkBaru)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated) // 201
			json.NewEncoder(w).Encode(produkBaru)
		}
	})

	// localhost:8080/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "API running",
		})
	})
	fmt.Println(("Server running di localhost:8080"))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("gagal running server")
	}
}
