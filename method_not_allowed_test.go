package main

import (
	"fmt"                               // untuk menulis output string ke response
	"github.com/julienschmidt/httprouter" // router HTTP ringan dengan dukungan routing berbasis method
	"github.com/stretchr/testify/assert" // library assertion untuk validasi hasil unit test
	"io"                                // untuk membaca isi response body
	"net/http"                          // package HTTP utama Go
	"net/http/httptest"                 // package untuk simulasi HTTP request saat testing
	"testing"                           // package bawaan Go untuk unit test
)

func TestMethodNotAllowed(t *testing.T) { // unit test untuk menguji handler MethodNotAllowed
	router := httprouter.New() // membuat router baru

	// set custom handler jika method tidak diizinkan pada path yang ada
	router.MethodNotAllowed = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Method Not Allowed") // tulis pesan ke response body
	})

	// daftarkan route "/" hanya untuk method POST
	router.POST("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "POST") // response jika method POST dipanggil
	})

	// buat request GET ke path "/" → path ada tapi method salah
	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)

	// recorder untuk menangkap response
	recorder := httptest.NewRecorder()

	// jalankan router dengan request → akan memicu MethodNotAllowed handler
	router.ServeHTTP(recorder, request)

	response := recorder.Result()        // ambil hasil response
	body, _ := io.ReadAll(response.Body) // baca isi body response
	
	// pastikan response sesuai dengan handler MethodNotAllowed
	assert.Equal(t, "Method Not Allowed", string(body))
}

// Kesimpulan:
// Kode ini menguji fitur MethodNotAllowed pada httprouter. Route "/" hanya didaftarkan untuk POST, tetapi test mengirim request GET ke path yang sama. Karena path ada namun method tidak cocok, router memanggil MethodNotAllowed handler. Dengan httptest, request disimulasikan dan assert memastikan response body sesuai harapan.
