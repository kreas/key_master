package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kreas/key_master/actions"
)

const port string = ":8000"

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/{id}", actions.GetDataHandler).Methods("GET")
	r.HandleFunc("/{id}", actions.AddDataHandler).Methods("POST")

	http.Handle("/", &KeyMasterServer{r})
	log.Printf("Server started at http://localhost" + port)
	log.Fatal("Server started at http://localhost" + port)
}

// KeyMasterServer struct
type KeyMasterServer struct {
	r *mux.Router
}

// ServeHTTP updates the headers of our applicaiton
func (s *KeyMasterServer) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// Set up CORS
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	rw.Header().Set("Access-Control-Allow-Headers",
		"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	// Since this is an API only application let's set the content type to application/json
	rw.Header().Set("Content-Type", "application/json")

	// Stop here if its Preflighted OPTIONS request
	if req.Method == "OPTIONS" {
		return
	}

	// Lets Gorilla work
	s.r.ServeHTTP(rw, req)
}
