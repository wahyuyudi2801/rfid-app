package employeecontroller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rfid-app/config"
	"rfid-app/models"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[GET] mengambil data karyawan")
	var employees []models.Employee

	rows, err := config.DB.Query("SELECT * FROM employees")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	for rows.Next() {
		var employee models.Employee
		rows.Scan(&employee.ID, &employee.Name, &employee.Phone, &employee.Email, &employee.RegistrationNumber, &employee.Position, &employee.WorkUnit, &employee.Status)

		employees = append(employees, employee)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"status":  true,
		"message": "Berhasil mengambil data karyawan",
		"data":    employees,
	})
}

func Show(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[GET] mengambil detail data karyawan")
	id := r.PathValue("id")
	var employee models.Employee

	if err := findById(id, &employee); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]any{
		"status":  true,
		"message": "berhasil mengambil data karyawan id " + id,
		"data":    employee,
	})
}

func Store(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[POST] membuat data karyawan")
	var employee models.Employee

	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := config.DB.Exec("INSERT INTO employees (name, phone, email, registration_number, position, work_unit, status) VALUES (?, ?, ?, ?, ?, ?, ?)", employee.Name, employee.Phone, employee.Email, employee.RegistrationNumber, employee.Position, employee.WorkUnit, employee.Status)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	employee.ID = uint(id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{
		"status":  true,
		"message": "berhasil menambahkan karyawan baru",
		"data":    employee,
	})
}
func Update(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[PATCH] memperbarui data karyawan")
	var employee models.Employee
	id := r.PathValue("id")

	if err := findById(id, &employee); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := config.DB.Exec("UPDATE employees SET name=?, phone=?, email=?, registration_number=?, position=?, work_unit=?, status=? WHERE id = ?", employee.Name, employee.Phone, employee.Email, employee.RegistrationNumber, employee.Position, employee.WorkUnit, employee.Status, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"status":  true,
		"message": "berhasil memperbarui data karyawan id " + id,
	})
}
func Destroy(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[DELETE] menghapus data karyawan")
	id := r.PathValue("id")
	var employee models.Employee

	if err := findById(id, &employee); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	_, err := config.DB.Exec("DELETE FROM employees WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]any{
		"status":  true,
		"message": "berhasil menghapus data karyawan id " + id,
	})
}

func findById(id string, employee *models.Employee) error {
	err := config.DB.QueryRow("SELECT * FROM employees WHERE id = ?", id).Scan(
		&employee.ID, &employee.Name, &employee.Phone, &employee.Email, &employee.RegistrationNumber, &employee.Position, &employee.WorkUnit, &employee.Status,
	)
	return err
}
