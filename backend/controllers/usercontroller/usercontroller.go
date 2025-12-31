package usercontroller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"rfid-app/config"
	"rfid-app/utilities"
)

type User struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	RoleID     int    `json:"role_id"`
	EmployeeID int    `json:"employee_id"`
}

type InputUser struct {
	Username             string `json:"username"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
	RoleID               int    `json:"role_id"`
	EmployeeID           int    `json:"employee_id"`
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[GET] Mengambil data user")
	rows, err := config.DB.Query("SELECT id, username, role_id, employee_id FROM users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.RoleID, &user.EmployeeID); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"status":  true,
		"message": "Berhasil mengambil data user",
		"data":    users,
	})
}

func Show(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[GET] Mengambil data user detail")

	id := r.PathValue("id")
	var user User

	row := config.DB.QueryRow("SELECT id, username, role_id, employee_id FROM users WHERE id = ?", id)
	if err := row.Scan(&user.ID, &user.Username, &user.RoleID, &user.EmployeeID); err == sql.ErrNoRows {
		http.Error(w, "User tidak ditemukan", http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"status":  true,
		"message": "Berhasil mengambil data user dengan ID: " + id,
		"data":    user,
	})
}
func Store(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[POST] Membuat data user")
	var inputUser InputUser
	if err := json.NewDecoder(r.Body).Decode(&inputUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if inputUser.Username == "" {
		http.Error(w, "Username harus diisi", http.StatusBadRequest)
		return
	}

	if inputUser.RoleID == 0 {
		http.Error(w, "Role harus diisi", http.StatusBadRequest)
		return
	}

	if inputUser.EmployeeID == 0 {
		http.Error(w, "Employee harus diisi", http.StatusBadRequest)
		return
	}

	if inputUser.Password == "" {
		http.Error(w, "Password harus diisi", http.StatusBadRequest)
		return
	}

	if inputUser.PasswordConfirmation == "" {
		http.Error(w, "Konfirmasi Password harus diisi", http.StatusBadRequest)
		return
	}

	if inputUser.Password != inputUser.PasswordConfirmation {
		http.Error(w, "Password confirmation do not match", http.StatusBadRequest)
		return
	}

	hash, _ := utilities.HashPassword(inputUser.Password)
	inputUser.Password = hash

	querySql := "INSERT INTO users (username, password, role_id, employee_id) VALUES (?, ?, ?, ?)"
	result, err := config.DB.Exec(querySql, inputUser.Username, hash, inputUser.RoleID, inputUser.EmployeeID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	user := User{
		ID:         int(id),
		Username:   inputUser.Username,
		RoleID:     inputUser.RoleID,
		EmployeeID: inputUser.EmployeeID,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

func Update(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[PATCH] Memperbarui data user")

	id := r.PathValue("id")
	var inputUser InputUser
	var user User
	var message string

	if err := json.NewDecoder(r.Body).Decode(&inputUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var querySql string

	querySql = "SELECT id FROM users WHERE id = ?"
	if err := config.DB.QueryRow(querySql, id).Scan(&user.ID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if inputUser.Password != "" {
		if inputUser.Password != inputUser.PasswordConfirmation {
			http.Error(w, "Password confirmation do not match", http.StatusBadRequest)
			return
		}

		hash, _ := utilities.HashPassword(inputUser.Password)
		inputUser.Password = hash

		querySql = "UPDATE users SET username=?, password=?, role_id=?, employee_id=? WHERE id = ?"
		config.DB.Exec(querySql, inputUser.Username, inputUser.Password, inputUser.RoleID, inputUser.EmployeeID, id)
		message = "berhasil memperbarui data user dan password. ID: "
	} else {
		querySql = "UPDATE users SET username=?, role_id=?, employee_id=? WHERE id = ?"
		config.DB.Exec(querySql, inputUser.Username, inputUser.RoleID, inputUser.EmployeeID, id)
		message = "berhasil memperbarui data user. ID: "
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"status":  true,
		"message": message + id,
	})
}
func Destroy(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[DELETE] Menghapus data user")
	id := r.PathValue("id")
	var querySql string

	querySql = "DELETE FROM users WHERE id = ?"
	_, err := config.DB.Exec(querySql, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"status":  true,
		"message": "Berhasil menghapus user id: " + id,
	})
}
