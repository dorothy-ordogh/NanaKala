package models

import (
	// "encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	// "time"
	// "math/rand"
)

type Group struct {
	GroupID int64 `json:"id"`
	GroupName string `json:"gname"`
	GroupUsers []User `json:"users"`
}


func HandleGroup(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	switch req.Method {
	case "POST":
		// group stuff
	case "GET", "DELETE", "PUT":
		res.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func HandleGroupWithID(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	userid := vars["id"]

	fmt.Println(userid)

	switch req.Method {
	case "GET":
		// lookup user in db by id and return
	case "PUT":
		// update user in db by first lookup and then posting
	case "POST":
		res.WriteHeader(http.StatusMethodNotAllowed)
	case "DELETE":
		// delete user in db
	}
}

func HandleGroupExpenses(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	userid := vars["id"]

	fmt.Println(userid)

	switch req.Method {
	case "GET":
		// lookup user expenses and return all
	case "PUT":
		res.WriteHeader(http.StatusMethodNotAllowed)
	case "POST":
		res.WriteHeader(http.StatusMethodNotAllowed)
	case "DELETE":
		// delete all user expenses
	}
}

func HandleGroupBudgets(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	userid := vars["id"]

	fmt.Println(userid)

	switch req.Method {
	case "GET":
		// lookup user budgets and return all
	case "PUT":
		res.WriteHeader(http.StatusMethodNotAllowed)
	case "POST":
		res.WriteHeader(http.StatusMethodNotAllowed)
	case "DELETE":
		// delete all user budgets
	}
}