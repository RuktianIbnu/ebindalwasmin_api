package routes

import (
	"ebindalwasmin_api/controllers"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// NewRouter ...
func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/login", controllers.Login).Methods("POST")
	return r
}

// LoadCors ...
func LoadCors(r http.Handler) http.Handler {
	headers := handlers.AllowedHeaders([]string{"X-Request", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	return handlers.CORS(headers, methods, origins)(r)
}
