package controller

import (
	"database/sql"
	"net/http"
	"path/filepath"
	"strconv"
	"text/template"
)

func NewUpdateEmployeeController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			id := r.URL.Query().Get("id")
			r.ParseForm()

			name := r.Form["name"][0]
			npwp := r.Form["npwp"][0]
			address := r.Form["address"][0]
			//query := "INSERT INTO employee (name, npwp, address) VALUES ($1, $2, $3);"
			_, err := db.Exec("UPDATE employee SET name=$1, npwp=$2, address=$3 WHERE id=$4", name, npwp, address, id)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/employee", http.StatusMovedPermanently)
			return
		} else if r.Method == "GET" {
			idStr := r.URL.Query().Get("Id")

			if idStr == "" {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("ID tidak boleh kosong"))
				return
			}

			id, err := strconv.Atoi(idStr)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("ID tidak valid"))
				return
			}

			row := db.QueryRow("SELECT name, npwp, address FROM employee WHERE id = $1 ", id)
			if row.Err() != nil {
				w.Write([]byte(row.Err().Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			var employee Employee
			err = row.Scan(
				&employee.Name,
				&employee.NPWP,
				&employee.Address,
			)
			employee.Id = idStr
			if err != nil {
				w.Write([]byte(row.Err().Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			fp := filepath.Join("views", "update.html")
			tmpl, err := template.ParseFiles(fp)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			data := make(map[string]any)
			data["employee"] = employee

			err = tmpl.Execute(w, data)
			if err != nil {
				w.Write([]byte(err.Error()))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
}
