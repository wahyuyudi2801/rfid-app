package routes

import (
	"net/http"
	"rfid-app/controllers/usercontroller"
)

func UserRoute(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/users", usercontroller.Index)
	mux.HandleFunc("GET /api/users/{id}", usercontroller.Show)
	mux.HandleFunc("POST /api/users", usercontroller.Store)
	mux.HandleFunc("PATCH /api/users/{id}", usercontroller.Update)
	mux.HandleFunc("DELETE /api/users/{id}", usercontroller.Destroy)
}
