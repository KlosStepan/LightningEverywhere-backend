package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"lightningeverywhere_backend/internal/eshop"
	"lightningeverywhere_backend/internal/merchant"

	"lightningeverywhere_backend/db"

	"go.mongodb.org/mongo-driver/mongo"
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
	// Connect to MongoDB
	var mongoDB *mongo.Database = db.ConnectMongo()
	//_ = mongoDB // TODO: pass to actual store
	log.Printf("MongoDB initialized for database: %s", mongoDB.Name())

	// Choose store implementation — memory for now
	var store1 eshop.Store = eshop.NewMemoryStore()
	var store2 merchant.Store = merchant.NewMemoryStore()

	// Setup routes
	var mux *http.ServeMux = http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"ok","message":"Welcome to Lightning Everywhere backend"}`))
	})
	mux.Handle("/api/eshops", eshop.MakeHandler(store1)) // all methods (GET, POST, etc.)
	mux.Handle("/api/merchants", merchant.MakeHandler(store2))

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
