package models

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type Budget struct {
	BudgetID int64 `json:"id"`
	BudgetAmount float64 `json:"amt"`
	BudgetName string `json:"name"`
	BudgetCategory int64 `json:"category"`
	BudgetGID int64 `json:"gid"`
	BudgetUID int64 `json:"uid"`
}

func HandleBudget(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.Header().Add("Access-Control-Allow-Origin", "*")

	switch req.Method {
	case "POST":
		// budget stuff
		budget := new(Budget)
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&budget)
		checkErr(err, res)

		if budget.BudgetName == "" || budget.BudgetAmount == 0 {
			err = fmt.Errorf("Required information is not included in submitted budget")
			checkErr(err, res)
		}

		result, err := DB_CONNECTION.Exec("INSERT INTO budget (budget_id, budget_name, budget_amt) VALUES (?, ?, ?)", nil, budget.BudgetName, budget.BudgetAmount)
		checkErr(err, res)

		id, err := result.LastInsertId()
		budget.BudgetID = id

		if budget.BudgetGID != 0 {

			fmt.Println("adding to group budgets")
			// insert into group_budgets
			result, err := DB_CONNECTION.Exec("INSERT INTO group_budgets (group_id, budget_id) VALUES (?, ?)", budget.BudgetGID, budget.BudgetID)
			checkErr(err, res)

			affected, err := result.RowsAffected()

			if affected < 1 {
				err = fmt.Errorf("No insert, please verify group ID: %d", budget.BudgetGID)
				checkErr(err, res)
			}

		} else if budget.BudgetUID != 0 {
			// insert into user_budgets
			result, err := DB_CONNECTION.Exec("INSERT INTO user_budgets (user_id, budget_id) VALUES (?, ?)", budget.BudgetUID, budget.BudgetID)
			checkErr(err, res)

			affected, err := result.RowsAffected()

			if affected < 1 {
				err = fmt.Errorf("No insert, please verify user ID: %d", budget.BudgetUID)
				checkErr(err, res)
			}
		}

		if budget.BudgetCategory != 0 {
			result, err := DB_CONNECTION.Exec("INSERT INTO budget_cat (cat_id, budget_id) VALUES (?, ?)", budget.BudgetCategory, budget.BudgetID)
			checkErr(err, res)

			affected, err := result.RowsAffected()

			if affected < 1 {
				err = fmt.Errorf("No insert, please verify category: %d", budget.BudgetCategory)
				checkErr(err, res)
			}
		}

		outgoingJson, err := json.Marshal(budget)
		checkErr(err, res)

		res.WriteHeader(http.StatusOK)
		fmt.Fprint(res, string(outgoingJson))
	case "GET", "DELETE", "PUT":
		res.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func HandleBudgetWithID(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.Header().Add("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(req)
	bidstr := vars["id"]
	bid, err := strconv.ParseInt(bidstr, 10, 64)
	checkErr(err, res)

	switch req.Method {
	case "GET":
		// lookup budget in db by id and return
		// future work: return expenses filed under requested budget
		var b Budget

		var exists int
		prep, err := DB_CONNECTION.Prepare("SELECT EXISTS(SELECT 1 FROM budget WHERE budget_id = ?)")
		checkErr(err, res)

		err = prep.QueryRow(bid).Scan(&exists)
		checkErr(err, res)

		if exists == 0 {
			err = fmt.Errorf("Requested budget does not exist")
			checkErr(err, res)
		}

		prep, err = DB_CONNECTION.Prepare("SELECT budget_id, budget_amt, budget_name FROM budget WHERE budget_id = ?")
		checkErr(err, res)
		
		err = prep.QueryRow(bid).Scan(&b.BudgetID, &b.BudgetAmount, &b.BudgetName)
		checkErr(err, res)

		prep, err = DB_CONNECTION.Prepare("SELECT cat_id FROM budget_cat WHERE budget_id = ?")
		checkErr(err, res)
		
		err = prep.QueryRow(b.BudgetID).Scan(&b.BudgetCategory)
		checkErr(err, res)

		prep, err = DB_CONNECTION.Prepare("SELECT EXISTS(SELECT 1 FROM user_budgets WHERE budget_id = ?)")
		checkErr(err, res)

		err = prep.QueryRow(bid).Scan(&exists)
		checkErr(err, res)

		if exists == 1 {
			prep, err := DB_CONNECTION.Prepare("SELECT user_id FROM user_budgets WHERE budget_id = ?")
			checkErr(err, res)

			err = prep.QueryRow(bid).Scan(&b.BudgetUID)
			checkErr(err, res)
		} else {
			prep, err := DB_CONNECTION.Prepare("SELECT group_id FROM group_budgets WHERE budget_id = ?")
			checkErr(err, res)

			err = prep.QueryRow(bid).Scan(&b.BudgetGID)
			checkErr(err, res)
		}

		outgoingJson, err := json.Marshal(b)
		checkErr(err, res)

		res.WriteHeader(http.StatusOK)
		fmt.Fprint(res, string(outgoingJson))
	case "PUT":
		// update budget in db by first lookup and then posting
		budget := new(Budget)
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&budget)
		checkErr(err, res)

		if budget.BudgetName == "" || budget.BudgetAmount == 0 {
			err = fmt.Errorf("Required information is not included in submitted budget")
			checkErr(err, res)
		}

		var exists int
		prep, err := DB_CONNECTION.Prepare("SELECT EXISTS(SELECT 1 FROM budget WHERE budget_id = ?)")
		checkErr(err, res)

		err = prep.QueryRow(bid).Scan(&exists)
		checkErr(err, res)

		if exists == 0 {
			err = fmt.Errorf("Requested budget does not exist")
			checkErr(err, res)
		}

		prep, err = DB_CONNECTION.Prepare("UPDATE budget SET budget_name = ?, budget_amt = ? WHERE budget_id = ?")
		checkErr(err, res)
		
		result, err := prep.Exec(budget.BudgetName, budget.BudgetAmount, bid)
		checkErr(err, res)

		affected, err := result.RowsAffected()

		if affected > 1 {
			err = fmt.Errorf("Too many rows were affected, please verify budget ID: %d", bid)
			checkErr(err, res)
		}

		if budget.BudgetCategory != 0 {
			result, err := DB_CONNECTION.Exec("UPDATE budget_cat SET cat_id = ? WHERE budget_id = ?", budget.BudgetCategory, budget.BudgetID)
			checkErr(err, res)

			affected, err := result.RowsAffected()

			if affected > 1 {
				err = fmt.Errorf("Too many rows were affected, please verify category: %d", budget.BudgetCategory)
				checkErr(err, res)
			}
		}

		outgoingJson, err := json.Marshal(budget)
		checkErr(err, res)

		res.WriteHeader(http.StatusOK)
		fmt.Fprint(res, string(outgoingJson))
	case "POST":
		res.WriteHeader(http.StatusMethodNotAllowed)
	case "DELETE":
		// delete budget in db
		prep, err := DB_CONNECTION.Prepare("DELETE FROM budget WHERE budget_id = ?")
		checkErr(err, res)
		
		result, err := prep.Exec(bid)
		checkErr(err, res)

		affected, err := result.RowsAffected()

		if affected > 1 {
			fmt.Println("more than 1 row affected. not sure if this is cascade")
			// err = fmt.Errorf("Too many rows were affected, please verify userID: %d", userid)
			// checkErr(err, res)
		}

		res.WriteHeader(http.StatusOK)
	}
}