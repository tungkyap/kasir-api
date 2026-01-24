package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateCategoryRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

var category = []Category{
	{ID: 1, Name: "Makanan", Description: "Kategori Makanan"},
	{ID: 2, Name: "Minuman", Description: "Kategori Minuman"},
	{ID: 3, Name: "Mainan", Description: "Kategori Mainan"},
}

// sesuaikan dengan data di kategori, untuk handle ketika hapus dan update data kategori
var lastCategoryID = 3

// POST localhost:8080/api/category
func createCategory(w http.ResponseWriter, r *http.Request) {
	// Buat header untuk terima data JSON
	w.Header().Set("Content-Type", "application/json")
	// buat variabel objek kategori baru untuk simpan data r.Body dari FrontEnd
	var newCategory Category
	// baca request body dan masukan hasil request body ke variabel err
	err := json.NewDecoder(r.Body).Decode(&newCategory)
	// cek apakah err kosong atau tidak, jika kosong maka return error invalid request,
	// jika tidak kosong maka abaikan dan lanjut kebawah
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	// tambah id baru rumus lastCategoryID + 1 (cek deskripsi variabel global lastCategoryID di atas)
	lastCategoryID++
	// masukan data ke dalam variabel kategori baru yang dikirim r.Body dari FrontEnd
	// masukan id kategori baru ke variabel objek newCategory
	newCategory.ID = lastCategoryID
	// masukan sisa data baru dari r.Body ke variabel objek global category
	category = append(category, newCategory)
	w.WriteHeader(http.StatusCreated) // 201
	json.NewEncoder(w).Encode(newCategory)
}

// GET localhost:8080/api/category
func getAllCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Panggil semua kategori di dalam global variabel objek kategori dengan format JSON
	json.NewEncoder(w).Encode(category)
}

// GET localhost:8080/api/category/{id}
func getCategoryByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// trim string url /api/category/{id} untuk hanya dapatkan {id} dari path param
	// path param dikirim dari FrontEnd dengan trim prefix
	idStr := strings.TrimPrefix(r.URL.Path, "/api/category/")
	// ubah string menjadi int dan masukan ke dalam 2 variabel id dan err
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}

	// looping variabel array of objek category
	for _, c := range category {
		// cek id dari category dicocokan dengan string dari variabel id setelah di ubah dari string ke int
		if c.ID == id {
			// ambil data yang sudah cocok dan ditemukan kemudian tampilkan data dengan format JSON
			json.NewEncoder(w).Encode(c)
			return
		}
	}

	http.Error(w, "Category belum ada", http.StatusNotFound)
}

// PUT localhost:8080/api/category/{id}
func updateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := strings.TrimPrefix(r.URL.Path, "/api/category/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Produk ID", http.StatusBadRequest)
		return
	}
	// get data dari request buat variabel baru dengan nama updateCategory
	var updateReq UpdateCategoryRequest
	err = json.NewDecoder(r.Body).Decode(&updateReq)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	// loop category untuk cari id yang sesuai kemudian ganti sesuai data dari request body
	for i := range category {
		if category[i].ID == id {
			if updateReq.Name != nil {
				category[i].Name = *updateReq.Name
			}

			if updateReq.Description != nil {
				category[i].Description = *updateReq.Description
			}
			json.NewEncoder(w).Encode(category[i])
			return
		}
	}
	http.Error(w, "Category belum ada", http.StatusNotFound)
}

// DELETE localhost:8080/api/category/{id}
func deleteCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idStr := strings.TrimPrefix(r.URL.Path, "/api/category/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Category ID", http.StatusBadRequest)
		return
	}

	for i, c := range category {
		if c.ID == id {
			category = append(category[:i], category[i+1:]...)
			json.NewEncoder(w).Encode(map[string]string{
				"message": "sukses hapus category",
			})
			return
		}
	}
}
