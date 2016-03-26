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
	GroupUsers []*User `json:"users"`
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