package main

import (
	"fmt"                               // untuk format output string ke ResponseWriter
	"github.com/julienschmidt/httprouter" // router HTTP ringan dengan dukungan params & panic handler
	"github.com/stretchr/testify/assert" // library assertion untuk mempermudah pengecekan di unit test
	"io"                                // untuk membaca isi response body
	"net/http"                          // package HTTP utama Go
	"net/http/httptest"                 // package untuk membuat HTTP test server/recorder
	"testing"                           // package bawaan Go untuk unit testing
)

func TestPanicHandler(t *testing.T) { // fungsi unit test untuk menguji PanicHandler pada router
	router := httprouter.New() // membuat instance router baru

	// set custom panic handler → akan dipanggil jika terjadi panic di handler
	router.PanicHandler = func(writer http.ResponseWriter, request *http.Request, error interface{}) {
		fmt.Fprint(writer, "Panic : ", error) // kirim pesan panic ke response body
	}

	// definisikan route GET "/" yang sengaja memicu panic
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		panic("Ups") // trigger panic untuk mensimulasikan error
	})

	// membuat HTTP request palsu untuk kebutuhan testing
	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)

	// membuat recorder untuk menangkap response dari handler
	recorder := httptest.NewRecorder()	

	// jalankan router dengan request & recorder (simulasi request masuk)
	router.ServeHTTP(recorder, request)

	response := recorder.Result()        // ambil hasil response
	body, _ := io.ReadAll(response.Body) // baca seluruh isi body response

	// cek apakah output sesuai dengan yang diharapkan
	assert.Equal(t, "Panic : Ups", string(body))
}

// Kesimpulan:
// Kode ini adalah unit test untuk memastikan PanicHandler pada httprouter bekerja dengan benar. Route "/" sengaja dibuat panic, lalu PanicHandler menangkap panic tersebut dan menuliskan pesan ke response. Dengan httptest.NewRequest dan NewRecorder, request HTTP disimulasikan tanpa server sungguhan. Terakhir, assert digunakan untuk memverifikasi bahwa response body berisi teks panic yang diharapkan.
