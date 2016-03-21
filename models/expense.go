package models

import (
	// "encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	// "time"
	// "math/rand"
)

type Expense struct {
	ExpenseID int64 `json:"id"`
	ExpenseAmount float64 `json:"amt"`
	ExpenseCategory string `json:"category"`
	SplitWith []Split `json:"split"`
}

type Split struct {
	SplitUser User `json:"user"`
	SplitAmount float64 `json:"splitamt"`
	SplitPercent float64 `json:"splitpercent"`
	SplitID int64 `json:splitid`
}

func HandleExpense(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	switch req.Method {
	case "POST":
		// expense stuff
	case "GET", "DELETE", "PUT":
		res.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func HandleExpenseWithID(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	userid := vars["id"]

	fmt.Println(userid)
	
	switch req.Method {
	case "GET":
		// lookup expense in db by id and return
	case "PUT":
		// update expense in db by first lookup and then posting
	case "POST":
		res.WriteHeader(http.StatusMethodNotAllowed)
	case "DELETE":
		// delete expense in db
	}
}