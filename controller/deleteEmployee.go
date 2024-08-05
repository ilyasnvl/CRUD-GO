package controller

import (
	"database/sql"
	"net/http"
)

func NewDeleteEmployeeController(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("Id")

		_, err := db.Exec("DELETE FROM employee WHERE id = $1", id)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/employee", http.StatusMovedPermanently)
	}
}
