package models

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	// "time"
	// "math/rand"
)

type User struct {
	UserID int64 `json:"id"`
	FirstName string `json:"fname"`
	LastName string `json:"lname"`
	Email string `json:"email"`
	Phone int64 `json:"phone"`
}


func HandleUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	switch req.Method {
	case "POST":
		user := new(User)
		decoder := json.NewDecoder(req.Body)
		err := decoder.Decode(&user)
		checkErr(err, res)
		// if user.UserID == 0 {
		// 	seed := rand.NewSource(time.Now().UnixNano())
  //   		randomNumber := rand.New(seed)
  //   		user.UserID = randomNumber.Int63()
		// }
		// fmt.Println(user)

		prep, err := DB_CONNECTION.Prepare("INSERT INTO user (?, ?, ?, ?, ?)")
		checkErr(err, res)
		
		result, err := prep.Exec(nil, user.FirstName, user.LastName, user.Email, user.Phone)
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
	vars := mux.Vars(req)
	userid := vars["id"]

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
			// error!! too many rows effected!
		}

		outgoingJson, err := json.Marshal(user)
		checkErr(err, res)

		res.WriteHeader(http.StatusOK)
		fmt.Fprint(res, string(outgoingJson))
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
			// error!! too many rows effected!
		}

		res.WriteHeader(http.StatusOK)
	}
}

func HandleUserExpenses(res http.ResponseWriter, req *http.Request) {
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

func HandleUserBudgets(res http.ResponseWriter, req *http.Request) {
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

func checkErr(err error, res http.ResponseWriter) {
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}