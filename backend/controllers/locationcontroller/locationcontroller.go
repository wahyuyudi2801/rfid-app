package locationcontroller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rfid-app/config"
	"rfid-app/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[GET] menampilkan data location")
	var locations []models.Location

	rows, err := config.DB.Query("SELECT id, location_name, reader_code FROM locations")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var location models.Location
		rows.Scan(&location.ID, &location.LocationName, &location.ReaderCode)

		locations = append(locations, location)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"status":  true,
		"message": "berhasil mengambil data locations",
		"data":    locations,
	})
}

func Show(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Println("[GET] menampilkan detail location, id: " + id)
	var location models.Location

	if err := find(id, &location); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"status":  true,
		"message": "berhasil mengambil detail data location. id: " + id,
		"data":    location,
	})
}

func Store(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[POST] menambah location")
	var location models.Location
	json.NewDecoder(r.Body).Decode(&location)

	result, err := config.DB.Exec("INSERT INTO locations (location_name, reader_code) VALUES (?, ?)", location.LocationName, location.ReaderCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	location.ID = uint(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"status":  true,
		"message": "berhasil menambahkan data location",
		"data":    location,
	})
}

func Update(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Println("[PATCH] mengupdate location id: " + id)
	var location models.Location

	if err := find(id, &location); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&location); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := config.DB.Exec("UPDATE locations SET location_name=?, reader_code=? WHERE id=?", location.LocationName, location.ReaderCode, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"status":  true,
		"message": "berhasil mengupdate data location. ID: " + id,
	})
}

func Destroy(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Println("[DELETE] menghapus location id: " + id)
	var location models.Location

	if err := find(id, &location); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	_, err := config.DB.Exec("DELETE FROM locations WHERE id=?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"status":  true,
		"message": "berhasil menghapus data location dengan id " + id,
	})
}

// private func
func find(id string, location *models.Location) error {
	row := config.DB.QueryRow("SELECT id, location_name, reader_code FROM locations WHERE id=?", id)
	err := row.Scan(&location.ID, &location.LocationName, &location.ReaderCode)
	return err
}
