package routes

import (
	"net/http"
	"rfid-app/controllers/attendancecontroller"
)

func AttendanceRoute(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/attendances", attendancecontroller.Index)
	mux.HandleFunc("GET /api/attendances/{id}", attendancecontroller.Show)
	mux.HandleFunc("POST /api/attendances", attendancecontroller.Store)
	mux.HandleFunc("PATCH /api/attendances/{id}", attendancecontroller.Update)
	mux.HandleFunc("DELETE /api/attendances/{id}", attendancecontroller.Destroy)
}
