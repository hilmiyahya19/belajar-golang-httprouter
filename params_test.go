package main // Package utama aplikasi Go

import (
	"fmt" // Digunakan untuk menulis output ke response
	"github.com/julienschmidt/httprouter" // Library router HTTP dengan dukungan parameter URL
	"github.com/stretchr/testify/assert" // Library assertion untuk unit testing
	"io" // Digunakan untuk membaca response body
	"net/http" // Package standar HTTP (request & response)
	"net/http/httptest" // Package untuk membuat request dan response palsu saat testing
	"testing" // Package bawaan Go untuk unit test
)

func TestParams(t *testing.T) { // Fungsi unit test untuk menguji parameter pada route

	router := httprouter.New() // Membuat instance router baru

	router.GET("/products/:id", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		// Mendefinisikan route GET dengan parameter dinamis ":id"

		id := params.ByName("id")
		// Mengambil nilai parameter "id" dari URL

		text := "Product ID: " + id
		// Membuat response text berdasarkan parameter id

		fmt.Fprint(writer, text)
		// Mengirim response ke client
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/products/1", nil)
	// Membuat HTTP request palsu ke endpoint dengan parameter id = 1

	recorder := httptest.NewRecorder()
	// Membuat recorder untuk menangkap response server

	router.ServeHTTP(recorder, request)
	// Menjalankan request ke router dan menyimpan hasilnya

	response := recorder.Result()
	// Mengambil hasil response dari recorder

	body, _ := io.ReadAll(response.Body)
	// Membaca seluruh isi body response

	assert.Equal(t, "Product ID: 1", string(body))
	// Memastikan response body sesuai dengan output yang diharapkan
}

// KESIMPULAN:
// Kode ini adalah unit test untuk memastikan httprouter dapat menangani parameter URL dengan benar. Route "/products/:id" diuji dengan mensimulasikan request HTTP menggunakan httptest, kemudian nilai parameter diambil melalui params.ByName(). Response yang dihasilkan dibandingkan dengan nilai yang diharapkan menggunakan assertion untuk memastikan routing dan parameter parsing berjalan dengan baik.
