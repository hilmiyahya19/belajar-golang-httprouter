package main // Package utama aplikasi Go

import (
	"fmt" // Digunakan untuk menulis output ke response
	"github.com/julienschmidt/httprouter" // Router HTTP dengan dukungan named & catch-all parameter
	"github.com/stretchr/testify/assert" // Library assertion untuk unit testing
	"io" // Digunakan untuk membaca response body
	"net/http" // Package standar HTTP (request dan response)
	"net/http/httptest" // Package untuk simulasi HTTP request/response saat testing
	"testing" // Package bawaan Go untuk membuat unit test
)

func TestRouterPatternNamedParameter(t *testing.T) { // Unit test untuk named parameter lebih dari satu
	router := httprouter.New() // Membuat instance router baru

	router.GET("/products/:id/items/:itemId", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		// Mendefinisikan route dengan dua named parameter (:id dan :itemId)

		id := params.ByName("id")
		// Mengambil nilai parameter "id" dari URL

		itemId := params.ByName("itemId")
		// Mengambil nilai parameter "itemId" dari URL

		text := "Product ID: " + id + ", Item ID: " + itemId
		// Membentuk response text dari kedua parameter

		fmt.Fprint(writer, text)
		// Mengirim response ke client
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/products/1/items/1", nil)
	// Membuat request palsu dengan dua parameter URL

	recorder := httptest.NewRecorder()
	// Recorder untuk menangkap response

	router.ServeHTTP(recorder, request)
	// Menjalankan request ke router

	response := recorder.Result()
	// Mengambil hasil response

	body, _ := io.ReadAll(response.Body)
	// Membaca seluruh isi body response

	assert.Equal(t, "Product ID: 1, Item ID: 1", string(body))
	// Memastikan response sesuai dengan yang diharapkan
}

func TestRouterPatternCatchAllParameter(t *testing.T) { // Unit test untuk catch-all parameter (*parameter)
	router := httprouter.New() // Membuat instance router baru

	router.GET("/images/*image", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		// Mendefinisikan route dengan catch-all parameter "*image"

		image := params.ByName("image")
		// Mengambil seluruh sisa path setelah "/images/"

		text := "Image: " + image
		// Membentuk response text dari path yang ditangkap

		fmt.Fprint(writer, text)
		// Mengirim response ke client
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/images/characters/dottore.png", nil)
	// Membuat request palsu dengan path bertingkat (nested path)

	recorder := httptest.NewRecorder()
	// Recorder untuk menangkap response

	router.ServeHTTP(recorder, request)
	// Menjalankan request ke router

	response := recorder.Result()
	// Mengambil hasil response

	body, _ := io.ReadAll(response.Body)
	// Membaca seluruh isi body response

	assert.Equal(t, "Image: /characters/dottore.png", string(body))
	// Memastikan catch-all parameter menangkap path dengan benar
}

// KESIMPULAN:
// Kode ini berisi dua unit test untuk menguji kemampuan httprouter dalam menangani pola routing yang kompleks, yaitu named parameter lebih dari satu dan catch-all parameter. Test pertama memastikan parameter dinamis pada URL dapat diambil berdasarkan nama, sedangkan test kedua memastikan seluruh sisa path dapat ditangkap menggunakan wildcard (*). Pengujian dilakukan dengan mensimulasikan HTTP request menggunakan httptest dan memverifikasi response menggunakan assertion agar routing berjalan sesuai harapan.
