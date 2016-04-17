package main

import (
	"fmt"
	"net/http"
	"github.com/dorothy-ordogh/NanaKala/api"
	"github.com/dorothy-ordogh/NanaKala/models"
)

func main() {
	models.ConnectDB()

	defer models.DB_CONNECTION.Close()
	fmt.Println("Starting server")
	http.ListenAndServe(":8080", api.Handlers())
}
