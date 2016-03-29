package models

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
	"math/rand"
	"io/ioutil"
)

type Expense struct {
	ExpenseID int64 `json:"id"`
	ExpenseAmount float64 `json:"amt"`
	ExpenseCategory int64 `json:"category"`
	SplitWith []Split `json:"split"`
	UnderBudgetID int64 `json:"budgetid"`
	ExpenseGID int64 `json:"gid"`
	ExpenseName string `json:"name"`
}

type Split struct {
	SplitUser User `json:"user"`
	SplitAmount float64 `json:"splitamt"`
	SplitID int64 `json:splitid`
}

func HandleExpense(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	switch req.Method {
	case "POST":
		// expense stuff

		// Splits will be represented in the DB as new expenses for each
		// user with a special ID value in the split column of the expenses
		// table. This way we can query for every expense with the same 
		// split ID and thus

		// If a group ID is specified, then the expense will only be 
		// associated with the group, and not each individual user in the 
		// group. 
		exp := new(Expense)
		body, err := ioutil.ReadAll(req.Body)
		checkErr(err, res)
		err = json.Unmarshal(body, &exp)
		checkErr(err, res)

		if exp.ExpenseGID != 0 {
			// insert into group expense 

			result, err := DB_CONNECTION.Exec("INSERT INTO expense (expense_id, expense_name, expense_amt) VALUES (?, ?, ?)", nil, exp.ExpenseName, exp.ExpenseAmount)
			checkErr(err, res)

			id, err := result.LastInsertId()
			exp.ExpenseID = id

			result, err = DB_CONNECTION.Exec("INSERT INTO group_expenses (expense_id, group_id) VALUES (?, ?)", exp.ExpenseID, exp.ExpenseGID)
			checkErr(err, res)

			affected, err := result.RowsAffected()

			if affected < 1 {
				err = fmt.Errorf("No insert, please verify group ID: %d", exp.ExpenseGID)
				checkErr(err, res)
			}

			if exp.ExpenseCategory != 0 {
				result, err := DB_CONNECTION.Exec("INSERT INTO expense_cat (cat_id, expense_id) VALUES (?, ?)", exp.ExpenseCategory, exp.ExpenseID)
				checkErr(err, res)

				affected, err := result.RowsAffected()

				if affected < 1 {
					err = fmt.Errorf("No insert, please verify category: %d", exp.ExpenseCategory)
					checkErr(err, res)
				}
			}

			if exp.UnderBudgetID != 0 {
				result, err := DB_CONNECTION.Exec("INSERT INTO budget_expenses (expense_id, budget_id) VALUES (?, ?)", exp.ExpenseID, exp.UnderBudgetID)
				checkErr(err, res)

				affected, err := result.RowsAffected()

				if affected < 1 {
					err = fmt.Errorf("No insert, please verify budget ID %d exists", exp.UnderBudgetID)
					checkErr(err, res)
				}
			}
		} else {
			seed := rand.NewSource(time.Now().UnixNano())
    		randomNumber := rand.New(seed)
    		splitid := randomNumber.Int63()

			for _, split := range exp.SplitWith {

				result, err := DB_CONNECTION.Exec("INSERT INTO expense (expense_id, expense_amt, split_id) VALUES (?, ?, ?)", nil, split.SplitAmount, splitid)
				checkErr(err, res)

				expid, err := result.LastInsertId()

				usr := split.SplitUser
				fmt.Println(usr)

				result, err = DB_CONNECTION.Exec("INSERT INTO user_expenses (expense_id, user_id) VALUES (?, ?)", expid, usr.UserID)
				checkErr(err, res)

				affected, err := result.RowsAffected()

				if affected < 1 {
					err = fmt.Errorf("No insert, please verify user ID: %d", usr.UserID)
					checkErr(err, res)
				}

				if exp.ExpenseCategory != 0 {
					result, err := DB_CONNECTION.Exec("INSERT INTO expense_cat (cat_id, expense_id) VALUES (?, ?)", exp.ExpenseCategory, expid)
					checkErr(err, res)

					affected, err := result.RowsAffected()

					if affected < 1 {
						err = fmt.Errorf("No insert, please verify category: %d", exp.ExpenseCategory)
						checkErr(err, res)
					}
				}

				if exp.UnderBudgetID != 0 {
					result, err := DB_CONNECTION.Exec("INSERT INTO budget_expenses (expense_id, budget_id) VALUES (?, ?)", expid, exp.UnderBudgetID)
					checkErr(err, res)

					affected, err := result.RowsAffected()

					if affected < 1 {
						err = fmt.Errorf("No insert, please verify budget ID %d exists", exp.UnderBudgetID)
						checkErr(err, res)
					}
				}
			}
		}

		outgoingJson, err := json.Marshal(exp)
		checkErr(err, res)

		res.WriteHeader(http.StatusOK)
		fmt.Fprint(res, string(outgoingJson))

	case "GET", "DELETE", "PUT":
		res.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func HandleExpenseWithID(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	expid := vars["id"]
	
	switch req.Method {
	case "GET":
		// lookup expense in db by id and return

	case "PUT":
		// update expense in db by first lookup and then posting
	case "POST":
		res.WriteHeader(http.StatusMethodNotAllowed)
	case "DELETE":
		// delete expense in db
		prep, err := DB_CONNECTION.Prepare("DELETE FROM expense WHERE expense_id = ?")
		checkErr(err, res)
		
		result, err := prep.Exec(expid)
		checkErr(err, res)

		affected, err := result.RowsAffected()

		if affected > 1 {
			err = fmt.Errorf("Too many rows were affected, please verify expense ID: %d", expid)
			checkErr(err, res)
		}

		res.WriteHeader(http.StatusOK)
	}
}