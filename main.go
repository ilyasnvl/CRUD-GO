package main

import (
	"net/http"

	"github.com/ilyasnvl/crud-employee/database"
	"github.com/ilyasnvl/crud-employee/routes"
)

func main() {
	db := database.InitDatabase()

	server := http.NewServeMux()

	routes.MapRoutes(server, db)

	http.ListenAndServe(":8080", server)
}
