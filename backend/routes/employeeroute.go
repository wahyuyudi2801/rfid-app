package routes

import (
	"net/http"
	"rfid-app/controllers/employeecontroller"
)

func EmployeeRoute(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/employees", employeecontroller.Index)
	mux.HandleFunc("GET /api/employees/{id}", employeecontroller.Show)
	mux.HandleFunc("POST /api/employees", employeecontroller.Store)
	mux.HandleFunc("PATCH /api/employees/{id}", employeecontroller.Update)
	mux.HandleFunc("DELETE /api/employees/{id}", employeecontroller.Destroy)
}
