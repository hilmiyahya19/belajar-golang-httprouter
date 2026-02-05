package main

import (
	"fmt"                               // untuk menulis output string ke response
	"github.com/julienschmidt/httprouter" // router HTTP ringan dengan performa tinggi
	"github.com/stretchr/testify/assert" // library assertion untuk mempermudah validasi hasil test
	"io"                                // untuk membaca isi response body
	"net/http"                          // package HTTP utama Go
	"net/http/httptest"                 // package untuk simulasi request/response saat testing
	"testing"                           // package bawaan untuk unit test
)

func TestNotFound(t *testing.T) { // fungsi unit test untuk menguji handler NotFound
	router := httprouter.New() // membuat instance router baru

	// set custom handler jika route tidak ditemukan
	router.NotFound = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Not Found") // kirim teks "Not Found" ke response body
	})
	
	// membuat request palsu ke path yang tidak terdaftar di router
	request := httptest.NewRequest("GET", "http://localhost:3000/404", nil)

	// recorder untuk menangkap response dari router
	recorder := httptest.NewRecorder()

	// jalankan router dengan request → akan memicu NotFound handler
	router.ServeHTTP(recorder, request)

	response := recorder.Result()        // ambil hasil response
	body, _ := io.ReadAll(response.Body) // baca seluruh isi body response

	// verifikasi bahwa response sesuai dengan output NotFound handler
	assert.Equal(t, "Not Found", string(body))
}

// Kesimpulan:
// Kode ini adalah unit test untuk memastikan custom NotFound handler pada httprouter berjalan benar. Router tidak memiliki route "/404", sehingga handler NotFound dipanggil dan menulis "Not Found". Dengan httptest, request dan response disimulasikan tanpa server asli. Assertion di akhir test memastikan bahwa isi response body sesuai dengan yang diharapkan.
