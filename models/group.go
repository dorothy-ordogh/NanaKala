package models

import (
	// "encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	// "time"
	// "math/rand"
	"io/ioutil"
)

type Group struct {
	GroupID int64 `json:"id"`
	GroupName string `json:"gname"`
	GroupUsers []User `json:"users"` //GroupUsers []*User `json:"users"`
}


func HandleGroup(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	switch req.Method {
	case "POST":
		// group stuff
		group := new(Group)
		body, err := ioutil.ReadAll(req.Body)
		checkErr(err, res)
		err = json.Unmarshal(body, &group)
		// decoder := json.NewDecoder(req.Body)
		// err := decoder.Decode(&group)
		checkErr(err, res)

		fmt.Println(group.GroupName, group.GroupID, group.GroupUsers)

		result, err := DB_CONNECTION.Exec("INSERT INTO `group` (group_id, group_name) VALUES (?, ?)", nil, group.GroupName)
		checkErr(err, res)

		fmt.Println(result)

		id, err := result.LastInsertId()
		group.GroupID = id

		for _, guser := range group.GroupUsers {
			fmt.Println(guser)
			var exists int

			prep, err := DB_CONNECTION.Prepare("SELECT EXISTS(SELECT 1 FROM user WHERE user_id = ?)")
			checkErr(err, res)

			err = prep.QueryRow(guser.UserID).Scan(&exists)
			checkErr(err, res)

			fmt.Println(exists)

			if exists == 1 {
				result, err := DB_CONNECTION.Exec("INSERT INTO group_members (group_id, member_id) VALUES (?, ?)", group.GroupID, guser.UserID)
				checkErr(err, res)

				if val, _ := result.RowsAffected(); val == 0 {
					err = fmt.Errorf("Did not insert, please try again for member with id: %d", guser.UserID)
					checkErr(err, res)
				}
			} else {
				err = fmt.Errorf("User with id: %d does not exist", guser.UserID)
				checkErr(err, res)
			}
		}

		outgoingJson, err := json.Marshal(group)
		checkErr(err, res)

		res.WriteHeader(http.StatusOK)
		fmt.Fprint(res, string(outgoingJson))

	case "GET", "DELETE", "PUT":
		res.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func HandleGroupWithID(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	gid := vars["id"]

	fmt.Println(gid)

	switch req.Method {
	case "GET":
		// gets the group with the specified ID
		// will likely need a join to join the group members, group, and user tables

		prep, err := DB_CONNECTION.Prepare("SELECT * FROM group WHERE group_id = ?")
		checkErr(err, res)
		
		var g Group
		gusers := make([]User, 1)
		err = prep.QueryRow(gid).Scan(&g.GroupID, &g.GroupName)
		checkErr(err, res)

		prep, err := DB_CONNECTION.Prepare("SELECT user_id, user_fname, user_lname, user_email, user_phone FROM group_members T1 INNER JOIN user T2 ON T1.member_id = T2.user_id WHERE T1.group_id = ?")
		checkErr(err, res)
		
		rows, err = prep.Query(gid)

		for rows.Next() {
			var u User
			err = rows.Scan(&u.UserID, &u.FirstName, &u.LastName, &u.Email, &u.Phone)
			checkErr(err, res)

			gusers = append(gusers, u)
		}

		g.GroupUsers = gusers

		outgoingJson, err := json.Marshal(g)
		checkErr(err, res)

		res.WriteHeader(http.StatusOK)
		fmt.Fprint(res, string(outgoingJson))

	case "PUT":
		// update group name or users
	case "POST":
		res.WriteHeader(http.StatusMethodNotAllowed)
	case "DELETE":
		// delete group in db and group member associations with group 
		prep, err := DB_CONNECTION.Prepare("DELETE FROM group WHERE group_id = ?")
		checkErr(err, res)
		
		result, err := prep.Exec(gid)
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

func HandleGroupExpenses(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	gid := vars["id"]

	fmt.Println(gid)

	switch req.Method {
	case "GET":
		// lookup group expenses and return all

		expenseSlice := make([]Expense, 1)

		prep, err := DB_CONNECTION.Prepare("SELECT expense_id, expense_amt, split_id, expense_name FROM expense T1 INNER JOIN group_expenses T2 ON T1.expense_id = T2.expense_id WHERE T2.group_id = ?")
		checkErr(err, res)
		
		rows, err = prep.Query(gid)

		for rows.Next() {
			var exp Expense 
			var sid int64
			err = rows.Scan(&exp.ExpenseID, &exp.ExpenseAmount, &sid, &exp.ExpenseName)
			checkErr(err, res)

			prep, err := DB_CONNECTION.Prepare("SELECT cat_name FROM category T1 INNER JOIN expense_cat T2 ON T1.cat_id = T2.cat_id WHERE T2.expense_id = ?")
			checkErr(err, res)
		
			err = prep.QueryRow(exp.ExpenseID).Scan(&exp.ExpenseCategory)
			checkErr(err, res)

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
		// delete all group expenses

		rep, err := DB_CONNECTION.Prepare("DELETE T1 FROM expense T1 INNER JOIN group_expenses T2 ON T1.expense_id = T2.expense_id WHERE T2.group_id = ?")
		checkErr(err, res)
		
		result, err := prep.Exec(gid)
		checkErr(err, res)

		affected, err := result.RowsAffected()

		if affected < 1 {
			err = fmt.Errorf("Failed to delete expenses for group: %d", gid)
			checkErr(err, res)
		}

		res.WriteHeader(http.StatusOK)
	}
}

func HandleGroupBudgets(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)
	userid := vars["id"]

	fmt.Println(userid)

	switch req.Method {
	case "GET":
		// lookup group budgets and return all

		budgetSlice := make([]Budget, 1)

		prep, err := DB_CONNECTION.Prepare("SELECT budget_id, budget_amt, budget_name FROM budget T1 INNER JOIN group_budgets T2 ON T1.budget_id = T2.budget_id WHERE T2.group_id = ?")
		checkErr(err, res)
		
		rows, err = prep.Query(gid)

		for rows.Next() {
			var b Budget 
			err = rows.Scan(&b.BudgetID, &b.BudgetAmount, &b.BudgetName)
			checkErr(err, res)

			prep, err := DB_CONNECTION.Prepare("SELECT cat_name FROM category T1 INNER JOIN budget_cat T2 ON T1.cat_id = T2.cat_id WHERE T2.budget_id = ?")
			checkErr(err, res)
		
			err = prep.QueryRow(b.BudgetID).Scan(&b.BudgetCategory)
			checkErr(err, res)

			budgetSlice = append(budgetSlice, exp)
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
		// delete all group budgets

		prep, err := DB_CONNECTION.Prepare("DELETE T1 FROM budget T1 INNER JOIN group_budgets T2 ON T1.budget_id = T2.budget_id WHERE T2.group_id = ?")
		checkErr(err, res)
		
		result, err := prep.Exec(gid)
		checkErr(err, res)

		affected, err := result.RowsAffected()

		if affected < 1 {
			err = fmt.Errorf("Failed to delete budgets for user: %d", gid)
			checkErr(err, res)
		}

		res.WriteHeader(http.StatusOK)
	}
}