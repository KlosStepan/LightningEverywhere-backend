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
		var start time.Time = time.Now()
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Completed in %v", time.Since(start))
	})
}

func main() {
	// Choose store implementation — memory for now
	var store eshop.Store = eshop.NewMemoryStore()

	// Setup routes
	var mux *http.ServeMux = http.NewServeMux()
	mux.Handle("/api/eshops", eshop.MakeHandler(store)) // all methods (GET, POST, etc.)

	// Optional: Middleware (basic logger)
	var handler http.Handler = loggingMiddleware(mux)

	// Start server
	var port string = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Server running on http://localhost:%s", port)
	var err error = http.ListenAndServe(":"+port, handler)
	if err != nil {
		log.Fatal("Server failed:", err)
	}
}
