package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"lightningeverywhere_backend/internal/eshop"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Completed in %v", time.Since(start))
	})
}

func main() {
	// Choose store implementation — memory for now
	store := eshop.NewMemoryStore()

	// Setup routes
	mux := http.NewServeMux()
	mux.Handle("/api/eshops", eshop.MakeHandler(store)) // all methods (GET, POST, etc.)

	// Optional: Middleware (basic logger)
	handler := loggingMiddleware(mux)

	// Start server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on http://localhost:%s", port)
	err := http.ListenAndServe(":"+port, handler)
	if err != nil {
		log.Fatal("Server failed:", err)
	}
}
