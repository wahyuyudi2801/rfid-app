package rfidcardcontroller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rfid-app/config"
	"rfid-app/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[GET] menampilkan data RFID Card")
	var cards []models.RfidCard
	rows, err := config.DB.Query("SELECT id, rfid_tag, activation_date, status, employee_id FROM rfid_cards")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rows.Close()

	if rows.Next() {
		var card models.RfidCard
		rows.Scan(&card.ID, &card.RfidTag, &card.ActivationDate, &card.Status, &card.EmployeeID)

		cards = append(cards, card)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"status":  true,
		"message": "berhasil mengambil data RFID Card",
		"data":    cards,
	})
}

func Show(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Println("[GET] menampilkan detail RFID Card, id: " + id)
	var card models.RfidCard

	if err := find(id, &card); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"status":  true,
		"message": "berhasil mengambil data detail RFID Card dengan id: " + id,
		"data":    card,
	})
}

func find(id string, card *models.RfidCard) error {
	row := config.DB.QueryRow("SELECT id, rfid_tag, activation_date, status, employee_id FROM rfid_cards WHERE id=?", id)
	err := row.Scan(&card.ID, &card.RfidTag, &card.ActivationDate, &card.Status, &card.EmployeeID)
	return err
}

func Store(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[POST] menambah RFID Card")
	var card models.RfidCard

	if err := json.NewDecoder(r.Body).Decode(&card); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := config.DB.Exec("INSERT INTO rfid_cards (rfid_tag, activation_date, status, employee_id) VALUES (?, ?, ?, ?)", card.RfidTag, card.ActivationDate, card.Status, card.EmployeeID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	card.ID = uint(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"status":  true,
		"message": "berhasil menambahkan RFID Card",
		"data":    card,
	})
}

func Update(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Println("[PATCH] mengupdate RFID Card id: " + id)
	var card models.RfidCard

	if err := find(id, &card); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&card); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := config.DB.Exec("UPDATE rfid_cards SET status=? WHERE id=?", card.Status, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"status":  true,
		"message": "berhasil mengupdate data card dengan id: " + id,
	})
}

func Destroy(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Println("[DELETE] menghapus RFID Card id: " + id)
	var card models.RfidCard

	if err := find(id, &card); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err := config.DB.Exec("DELETE FROM rfid_cards WHERE id=?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"status":  true,
		"message": "berhasil menghapus rfid card dengan id " + id,
	})
}
