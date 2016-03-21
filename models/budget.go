package models

import (
	// "encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	// "time"
	// "math/rand"
)

type Budget struct {
	BudgetID int64 `json:"id"`
	BudgetAmount float64 `json:"amt"`
	BudgetName string `json:"name"`
	BudgetCategories []string `json:"categories"`
	BudgetGID int64 `json:"gid"`
	BudgetUID int64 `json:"uid"`
}

func HandleBudget(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	switch req.Method {
	case "POST":
		// budget stuff
	case "GET", "DELETE", "PUT":
		res.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func HandleBudgetWithID(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	userid := vars["id"]

	fmt.Println(userid)

	switch req.Method {
	case "GET":
		// lookup budget in db by id and return
	case "PUT":
		// update budget in db by first lookup and then posting
	case "POST":
		res.WriteHeader(http.StatusMethodNotAllowed)
	case "DELETE":
		// delete budget in db
	}
}