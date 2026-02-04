package main // Package utama aplikasi Go

import (
	"embed" // Digunakan untuk menanamkan file statis ke dalam binary Go
	"github.com/julienschmidt/httprouter" // Router HTTP dengan performa tinggi
	"github.com/stretchr/testify/assert" // Library assertion untuk unit testing
	"io" // Digunakan untuk membaca response body
	"io/fs" // Digunakan untuk mengelola filesystem virtual (FS)
	"net/http" // Package standar HTTP
	"net/http/httptest" // Package untuk simulasi HTTP request/response
	"testing" // Package bawaan Go untuk unit test
)

//go:embed resources
// Directive untuk menyertakan seluruh folder "resources" ke dalam binary
var resources embed.FS // File system virtual hasil embed

func TestServeFile(t *testing.T) { // Unit test untuk menguji serving file hello.txt
	router := httprouter.New() // Membuat instance router baru

	directory, _ := fs.Sub(resources, "resources")
	// Mengambil sub-folder "resources" dari embed FS

	router.ServeFiles("/files/*filepath", http.FS(directory))
	// Mengatur router agar dapat melayani file statis dari directory

	request := httptest.NewRequest("GET", "http://localhost:3000/files/hello.txt", nil)
	// Membuat request palsu untuk mengambil file hello.txt

	recorder := httptest.NewRecorder()
	// Recorder untuk menangkap response server

	router.ServeHTTP(recorder, request)
	// Menjalankan request ke router

	response := recorder.Result()
	// Mengambil hasil response

	body, _ := io.ReadAll(response.Body)
	// Membaca isi file dari response body

	assert.Equal(t, "Hello HttpRouter", string(body))
	// Memastikan isi file sesuai dengan yang diharapkan
}

func TestServeFileGoodBye(t *testing.T) { // Unit test untuk menguji serving file goodbye.txt
	router := httprouter.New() // Membuat instance router baru

	directory, _ := fs.Sub(resources, "resources")
	// Mengambil sub-folder "resources" dari embed FS

	router.ServeFiles("/files/*filepath", http.FS(directory))
	// Mengatur router agar melayani file statis dari directory

	request := httptest.NewRequest("GET", "http://localhost:3000/files/goodbye.txt", nil)
	// Membuat request palsu untuk mengambil file goodbye.txt

	recorder := httptest.NewRecorder()
	// Recorder untuk menangkap response server

	router.ServeHTTP(recorder, request)
	// Menjalankan request ke router

	response := recorder.Result()
	// Mengambil hasil response

	body, _ := io.ReadAll(response.Body)
	// Membaca isi file dari response body

	assert.Equal(t, "Good Bye HttpRouter", string(body))
	// Memastikan isi file sesuai dengan yang diharapkan
}

// KESIMPULAN:
// Kode ini menguji kemampuan httprouter dalam melayani file statis menggunakan fitur embed Go. Folder "resources" ditanam langsung ke dalam binary aplikasi, kemudian diakses sebagai filesystem virtual. Router dikonfigurasi menggunakan ServeFiles untuk menangani request file statis, dan pengujian dilakukan dengan mensimulasikan HTTP request menggunakan httptest serta memverifikasi isi file melalui assertion untuk memastikan file yang dilayani sesuai dengan harapan.
