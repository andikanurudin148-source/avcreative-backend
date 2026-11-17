package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	// 1. Ambil port dari environment variable (Render menyediakan variabel PORT secara otomatis)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Port default jika dijalankan secara lokal di komputer
	}

	// 2. Buat router atau handler sederhana
	mux := http.NewServeMux()

	// Contoh endpoint API sederhana
	mux.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "Halo dari Backend Go AV Creative!"}`))
	})

	// 3. Bungkus handler dengan Middleware CORS
	// CORS ini mengizinkan akses dari domain frontend avcreative.my.id
	handlerWithCORS := enableCORS(mux)

	// 4. Jalankan server pada port yang ditentukan
	log.Printf("Server berjalan di port %s...\n", port)
	if err := http.ListenAndServe(":"+port, handlerWithCORS); err != nil {
		log.Fatalf("Gagal menjalankan server: %v", err)
	}
}

// Fungsi Middleware untuk mengatur header CORS
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Batasi hanya domain frontend Anda yang diizinkan mengakses
		w.Header().Set("Access-Control-Allow-Origin", "https://avcreative.my.id")
		
		// Metode HTTP yang diizinkan
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		
		// Header yang diizinkan dikirim oleh client
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Tangani preflight request (OPTIONS) dari browser
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}