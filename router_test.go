package main // Package utama aplikasi Go

import (
	"fmt" // Digunakan untuk menulis output ke response
	"github.com/julienschmidt/httprouter" // Library router HTTP yang cepat dan ringan
	"github.com/stretchr/testify/assert" // Library untuk melakukan assertion pada unit test
	"io" // Digunakan untuk membaca isi response body
	"net/http" // Package standar untuk HTTP server dan request/response
	"net/http/httptest" // Package untuk membuat HTTP request dan response palsu (mock) saat testing
	"testing" // Package bawaan Go untuk membuat unit test
)

func TestRouter(t *testing.T) { // Fungsi unit test, harus diawali dengan "Test"

	router := httprouter.New() // Membuat instance router baru dari httprouter

	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		// Mendaftarkan handler untuk method GET pada path "/"
		fmt.Fprintf(writer, "Hello HTTP Router") // Menulis response ke client
	})

	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)
	// Membuat HTTP request palsu (GET /) untuk keperluan testing

	recorder := httptest.NewRecorder()
	// Membuat recorder untuk menangkap response dari server

	router.ServeHTTP(recorder, request)
	// Menjalankan request ke router dan menyimpan hasilnya ke recorder

	response := recorder.Result()
	// Mengambil hasil response HTTP dari recorder

	body, _ := io.ReadAll(response.Body)
	// Membaca seluruh isi body response

	assert.Equal(t, "Hello HTTP Router", string(body))
	// Memastikan response body sesuai dengan yang diharapkan
}

// KESIMPULAN:
// Kode ini merupakan unit test untuk memastikan router dari package httprouter bekerja dengan benar. Test ini membuat router, mendaftarkan route GET "/", lalu mensimulasikan HTTP request menggunakan httptest. Response yang dihasilkan ditangkap dan dibandingkan dengan output yang diharapkan menggunakan assertion, sehingga dapat dipastikan handler dan routing berjalan sesuai tujuan.
