package rolecontroller

import (
	"database/sql"
	"fmt"
	"log"
	"rfid-app/config"
)

type RoleStruct struct {
	ID   uint   `json:"id"`
	Role string `json:"role"`
}

func CreateRoles() {
	db := config.DB

	// Daftar role yang ingin dipastikan keberadaannya
	rolesToCreate := []struct {
		ID   int
		Name string
	}{
		{ID: 1, Name: "admin"},
		{ID: 2, Name: "user"},
	}

	for _, r := range rolesToCreate {
		var existingRole RoleStruct

		// 1. Cek apakah role dengan ID tersebut ada
		err := db.QueryRow("SELECT id, role FROM roles WHERE id = ?", r.ID).Scan(&existingRole.ID, &existingRole.Role)

		if err != nil {
			if err == sql.ErrNoRows {
				// 2. Jika tidak ditemukan (ErrNoRows), lakukan INSERT
				fmt.Printf("Role %s tidak ditemukan, membuat role baru...\n", r.Name)

				_, insertErr := db.Exec("INSERT INTO roles (id, role) VALUES (?, ?)", r.ID, r.Name)
				if insertErr != nil {
					log.Printf("Gagal membuat role %s: %v", r.Name, insertErr)
				} else {
					fmt.Printf("Role %s berhasil dibuat.\n", r.Name)
				}
			} else {
				// Jika ada error lain (misal koneksi terputus)
				log.Printf("Error saat mengecek role %d: %v", r.ID, err)
			}
		} else {
			fmt.Printf("Role %s sudah ada (ID: %d)\n", existingRole.Role, existingRole.ID)
		}
	}
}
