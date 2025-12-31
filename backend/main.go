package main

import (
	"fmt"
	"log"
	"net/http"
	"rfid-app/config"
	"rfid-app/controllers/rolecontroller"
	"rfid-app/routes"
)

func main() {
	config.InitDB()
	db := config.DB

	defer db.Close()

	port := ":8080"
	rolecontroller.CreateRoles()
	mux := http.NewServeMux()

	// Routing User
	routes.UserRoute(mux)

	// Routing Employee
	routes.EmployeeRoute(mux)

	// Routing Location
	routes.LocationRoute(mux)

	// Routing RFID Card
	routes.RfidCardRoute(mux)

	// Routing Attendance
	routes.AttendanceRoute(mux)

	fmt.Printf("Server berjalan di port: %s \n", port)
	log.Fatal(http.ListenAndServe(port, mux))
}
