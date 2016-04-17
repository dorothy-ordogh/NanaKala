package api

import (
	"github.com/gorilla/mux"
	"github.com/dorothy-ordogh/NanaKala/models"
)

func Handlers() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// USER
	router.HandleFunc("/user", models.HandleUser).Methods("GET", "PUT", "POST", "DELETE")
	router.HandleFunc("/user/{id:[0-9]+}", models.HandleUserWithID).Methods("GET", "PUT", "POST", "DELETE")
	router.HandleFunc("/user/{id:[0-9]+}/expenses", models.HandleUserExpenses).Methods("GET", "PUT", "POST", "DELETE")
	router.HandleFunc("/user/{id:[0-9]+}/budgets", models.HandleUserBudgets).Methods("GET", "PUT", "POST", "DELETE")

	// GROUP
	router.HandleFunc("/group", models.HandleGroup).Methods("GET", "PUT", "POST", "DELETE")
	router.HandleFunc("/group/{id:[0-9]+}", models.HandleGroupWithID).Methods("GET", "PUT", "POST", "DELETE")
	router.HandleFunc("/group/{id:[0-9]+}/expenses", models.HandleGroupExpenses).Methods("GET", "PUT", "POST", "DELETE")
	router.HandleFunc("/group/{id:[0-9]+}/budgets", models.HandleGroupBudgets).Methods("GET", "PUT", "POST", "DELETE")

	// EXPENSE
	router.HandleFunc("/expense", models.HandleExpense).Methods("GET", "PUT", "POST", "DELETE")
	router.HandleFunc("/expense/{id:[0-9]+}", models.HandleExpenseWithID).Methods("GET", "PUT", "POST", "DELETE")
	
	// BUDGET
	router.HandleFunc("/budget", models.HandleBudget).Methods("GET", "PUT", "POST", "DELETE")
	router.HandleFunc("/budget/{id:[0-9]+}", models.HandleBudgetWithID).Methods("GET", "PUT", "POST", "DELETE")
	
	return router
}