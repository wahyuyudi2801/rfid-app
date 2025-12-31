package attendancecontroller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rfid-app/config"
	"rfid-app/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[GET] menampilkan data attendance")
	var attendances []models.Attendance

	rows, err := config.DB.Query("SELECT * FROM attendances")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var attendance models.Attendance
		rows.Scan(&attendance.ID, &attendance.RfidID, &attendance.LocationID, &attendance.TapType, &attendance.ValidationStatus, &attendance.CreatedAt)

		attendances = append(attendances, attendance)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"status":  true,
		"message": "berhasil mengambil data attendances",
		"data":    attendances,
	})
}

func Show(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Println("[GET] menampilkan detail attendance, id: " + id)
	var attendance models.Attendance

	if err := find(id, &attendance); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"status":  true,
		"message": "berhasil mengambil data attendance detail. id: " + id,
		"data":    attendance,
	})
}

func Store(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[POST] menambah attendance")
	var attendance models.Attendance
	if err := json.NewDecoder(r.Body).Decode(&attendance); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	querySql := "INSERT INTO attendances (rfid_id, location_id, tap_type, validation_status, created_at) VALUES (?, ?, ?, ?, ?)"
	result, err := config.DB.Exec(querySql, attendance.RfidID, attendance.LocationID, attendance.TapType, attendance.ValidationStatus, attendance.CreatedAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id, _ := result.LastInsertId()
	attendance.ID = uint(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"status":  true,
		"message": "berhasil absen",
		"data":    attendance,
	})
}

func find(id string, attendance *models.Attendance) error {
	row := config.DB.QueryRow("SELECT * FROM attendances WHERE id=?", id)
	err := row.Scan(&attendance.ID, &attendance.RfidID, &attendance.LocationID, &attendance.TapType, &attendance.ValidationStatus, &attendance.CreatedAt)

	return err
}
