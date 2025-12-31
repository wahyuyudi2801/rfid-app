package routes

import (
	"net/http"
	"rfid-app/controllers/rfidcardcontroller"
)

func RfidCardRoute(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/rfid-cards", rfidcardcontroller.Index)
	mux.HandleFunc("GET /api/rfid-cards/{id}", rfidcardcontroller.Show)
	mux.HandleFunc("POST /api/rfid-cards", rfidcardcontroller.Store)
	mux.HandleFunc("PATCH /api/rfid-cards/{id}", rfidcardcontroller.Update)
	mux.HandleFunc("DELETE /api/rfid-cards/{id}", rfidcardcontroller.Destroy)
}
