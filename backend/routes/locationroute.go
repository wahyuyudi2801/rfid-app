package routes

import (
	"net/http"
	"rfid-app/controllers/locationcontroller"
)

func LocationRoute(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/locations", locationcontroller.Index)
	mux.HandleFunc("GET /api/locations/{id}", locationcontroller.Show)
	mux.HandleFunc("POST /api/locations", locationcontroller.Store)
	mux.HandleFunc("PATCH /api/locations/{id}", locationcontroller.Update)
	mux.HandleFunc("DELETE /api/locations/{id}", locationcontroller.Destroy)
}
