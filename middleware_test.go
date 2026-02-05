package main

import (
	"fmt"                               // untuk mencetak log ke console dan menulis response
	"github.com/julienschmidt/httprouter" // router HTTP ringan
	"github.com/stretchr/testify/assert" // library assertion untuk pengecekan hasil test
	"io"                                // untuk membaca isi response body
	"net/http"                          // package HTTP utama Go
	"net/http/httptest"                 // package untuk simulasi HTTP request saat testing
	"testing"                           // package bawaan Go untuk unit test
)

// struct middleware yang membungkus http.Handler
type LogMiddleware struct {
	http.Handler // handler asli yang akan dipanggil setelah middleware
}

// implementasi interface http.Handler → membuat LogMiddleware menjadi middleware valid
func (middleware *LogMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Received a request")               // log setiap request masuk
	middleware.Handler.ServeHTTP(writer, request)   // teruskan request ke handler berikutnya (router)
}

func TestMiddleware(t *testing.T) { // unit test untuk menguji middleware
	router := httprouter.New() // buat router baru

	// definisikan route GET "/"
	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		fmt.Fprint(writer, "Middleware") // tulis response body
	})

	middleware := LogMiddleware{router} // bungkus router dengan middleware

	// buat request palsu ke path "/"
	request := httptest.NewRequest("GET", "http://localhost:3000/", nil)

	// recorder untuk menangkap response
	recorder := httptest.NewRecorder()

	// jalankan request melalui middleware (bukan langsung ke router)
	middleware.ServeHTTP(recorder, request)

	response := recorder.Result()        // ambil hasil response
	body, _ := io.ReadAll(response.Body) // baca isi response body

	// validasi bahwa response tetap dihasilkan oleh handler router
	assert.Equal(t, "Middleware", string(body))
}

// Kesimpulan:
// Kode ini membuat middleware sederhana dengan membungkus http.Handler dan menambahkan log setiap request sebelum diteruskan ke router. Middleware mengimplementasikan ServeHTTP lalu memanggil handler asli di dalamnya. Pada unit test, router dibungkus oleh LogMiddleware dan request dijalankan melalui middleware tersebut. Test memastikan bahwa middleware tidak mengganggu response handler utama dan output tetap sesuai.
