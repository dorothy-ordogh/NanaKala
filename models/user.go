package models

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type User struct {
	UserID int64 `json:"id"`
	FirstName string `json:"fname"`
	LastName string `json:"lname"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}


func HandleUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.Header().Add("Access-Control-Allow-Origin", "*")

	switch req.Method {
	case "POST":

		user := new(User)
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&user)
		checkErr(err, res)

		result, err := DB_CONNECTION.Exec("INSERT INTO user (user_id, user_fname, user_lname, user_email, user_phone) VALUES (?, ?, ?, ?, ?)", nil, user.FirstName, user.LastName, user.Email, user.Phone)
		checkErr(err, res)

		id, err := result.LastInsertId()
		user.UserID = id

		outgoingJson, err := json.Marshal(user)
		checkErr(err, res)

		res.WriteHeader(http.StatusOK)
		fmt.Fprint(res, string(outgoingJson))

	case "GET", "DELETE", "PUT":
		res.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func HandleUserWithID(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.Header().Add("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(req)
	useridstr := vars["id"]
	userid, err := strconv.ParseInt(useridstr, 10, 64)
	checkErr(err, res)

	switch req.Method {
	case "GET":
		// lookup user in db by id and return
		prep, err := DB_CONNECTION.Prepare("SELECT * FROM user WHERE user_id = ?")
		checkErr(err, res)
		
		var u User
		err = prep.QueryRow(userid).Scan(&u.UserID, &u.FirstName, &u.LastName, &u.Email, &u.Phone)
		checkErr(err, res)

		outgoingJson, err := json.Marshal(u)
		checkErr(err, res)

		res.WriteHeader(http.StatusOK)
		fmt.Fprint(res, string(outgoingJson))
	case "PUT":
		// update user in db by first lookup and then posting
		user := new(User)
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&user)
		checkErr(err, res)

		prep, err := DB_CONNECTION.Prepare("UPDATE user SET user_fname = ?, user_lname = ?, user_email = ?, user_phone = ? WHERE user_id = ?")
		checkErr(err, res)
		
		result, err := prep.Exec(user.FirstName, user.LastName, user.Email, user.Phone, userid)
		checkErr(err, res)

		affected, err := result.RowsAffected()

		if affected > 1 {
			err = fmt.Errorf("Too many rows were affected, please verify userID: %d", userid)
			checkErr(err, res)
		} else if affected < 1 {
			err = fmt.Errorf("User does not exist, please verify user ID: %d", userid)
			checkErr(err, res)
		}

		res.WriteHeader(http.StatusOK)

	case "POST":
		res.WriteHeader(http.StatusMethodNotAllowed)
	case "DELETE":
		// delete user in db
		prep, err := DB_CONNECTION.Prepare("DELETE FROM user WHERE user_id = ?")
		checkErr(err, res)
		
		result, err := prep.Exec(userid)
		checkErr(err, res)

		affected, err := result.RowsAffected()

		if affected > 1 {
			err = fmt.Errorf("Too many rows were affected, please verify userID: %d", userid)
			checkErr(err, res)
		}

		res.WriteHeader(http.StatusOK)
	}
}

func HandleUserExpenses(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.Header().Add("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(req)
	useridstr := vars["id"]
	userid, err := strconv.ParseInt(useridstr, 10, 64)
	checkErr(err, res)

	switch req.Method {
	case "GET":
		// lookup user expenses and return all

		expenseSlice := make([]Expense, 0)

		prep, err := DB_CONNECTION.Prepare("SELECT T1.expense_id, T1.expense_amt, T1.split_id, COALESCE(T1.expense_name, '') as name FROM expense T1 INNER JOIN user_expenses T2 ON T1.expense_id = T2.expense_id WHERE T2.user_id = ?")
		checkErr(err, res)
		
		rows, err := prep.Query(userid)
		checkErr(err, res)

		for rows.Next() {
			var exp Expense 
			var sid int64
			err = rows.Scan(&exp.ExpenseID, &exp.ExpenseAmount, &sid, &exp.ExpenseName)
			checkErr(err, res)

			// How splitting works:
			// If an expense is split, it will have a split ID
			// All expenses that have the same split ID was a split
			// of a single expense between multiple people. Essentially,
			// a split expense forms several expenses for different users

			expenseSlice = append(expenseSlice, exp)
		}

		outgoingJson, err := json.Marshal(expenseSlice)
		checkErr(err, res)

		res.WriteHeader(http.StatusOK)
		fmt.Fprint(res, string(outgoingJson))

	case "PUT":
		res.WriteHeader(http.StatusMethodNotAllowed)
	case "POST":
		res.WriteHeader(http.StatusMethodNotAllowed)
	case "DELETE":
		// delete all user expenses

		prep, err := DB_CONNECTION.Prepare("DELETE T1 FROM expense T1 INNER JOIN user_expenses T2 ON T1.expense_id = T2.expense_id WHERE T2.user_id = ?")
		checkErr(err, res)
		
		result, err := prep.Exec(userid)
		checkErr(err, res)

		affected, err := result.RowsAffected()

		if affected < 1 {
			err = fmt.Errorf("Failed to delete expenses for user: %d", userid)
			checkErr(err, res)
		}

		res.WriteHeader(http.StatusOK)
	}
}

func HandleUserBudgets(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	res.Header().Add("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(req)
	useridstr := vars["id"]
	userid, err := strconv.ParseInt(useridstr, 10, 64)
	checkErr(err, res)

	switch req.Method {
	case "GET":
		// lookup user budgets and return all
		budgetSlice := make([]Budget, 0)

		prep, err := DB_CONNECTION.Prepare("SELECT T1.budget_id, T1.budget_amt, T1.budget_name FROM budget T1 INNER JOIN user_budgets T2 ON T1.budget_id = T2.budget_id WHERE T2.user_id = ?")
		checkErr(err, res)
		fmt.Println(userid)
		rows, err := prep.Query(userid)
		checkErr(err, res)
		fmt.Println(rows)
		for rows.Next() {
			var b Budget 
			err = rows.Scan(&b.BudgetID, &b.BudgetAmount, &b.BudgetName)
			checkErr(err, res)

			budgetSlice = append(budgetSlice, b)
		}

		outgoingJson, err := json.Marshal(budgetSlice)
		checkErr(err, res)

		res.WriteHeader(http.StatusOK)
		fmt.Fprint(res, string(outgoingJson))
	case "PUT":
		res.WriteHeader(http.StatusMethodNotAllowed)
	case "POST":
		res.WriteHeader(http.StatusMethodNotAllowed)
	case "DELETE":
		// delete all user budgets

		prep, err := DB_CONNECTION.Prepare("DELETE T1 FROM budget T1 INNER JOIN user_budgets T2 ON T1.budget_id = T2.budget_id WHERE T2.user_id = ?")
		checkErr(err, res)
		
		result, err := prep.Exec(userid)
		checkErr(err, res)

		affected, err := result.RowsAffected()

		if affected < 1 {
			err = fmt.Errorf("Failed to delete budgets for user: %d", userid)
			checkErr(err, res)
		}

		res.WriteHeader(http.StatusOK)
	}
}

func GetUserID(res *http.Response) int64 {
	user := new(User)
	decoder := json.NewDecoder(res.Body)
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Println("User could not be unmarshalled")
	}

	return user.UserID
}