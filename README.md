# NanaKala
What is NanaKala? In Hawaiian, "nana" is to watch or observe, and "kala" is money. This project is for a UBC's CPSC 448 (directed studies) course. I proposed to develop a REST API which would help track, and split expenses amongst users and groups, as well as budget money. 

### Components of the system:
- User
- Group (an extension of the user model with one or several users)
- Expense
- Budget (contains expenses and desired spending limit. Can nest budgets)

Budgets and expenses can apply to a single user, a group, a user within a group, a subset of users in a group or any combination. 

|   Resource   |   GET   |   PUT   |   POST   | DELETE   |
|----|----|----|----|----|
| ```/user``` | ```405```: method not allowed | ```405```: method not allowed | ```200```: creates a new user <br/>OR <br/>```400```: bad request if error | ```405```: method not allowed |
| ```/user/:id``` | ```200```: gets the user with specified id <br/> OR <br/>```404```: id not found <br/>OR <br/>```400```: bad request if error | ```200```: updates the specified user <br/>OR <br/>```404```: id not found <br/>OR <br/>```400```: bad request if error | ```405```: method not allowed | ```200```: deletes the user with the specified id and all their budgets and expenses <br/> OR <br/>```404```: id not found <br/>OR <br/>```400```: bad request if error |
| ```/user/:id/expenses``` | ```200```: gets the specified user's expenses <br/>OR <br/>```404```: id not found <br/>OR <br/>```400```: bad request if error | ```405```: method not allowed | ```405```: method not allowed | ```200```: deletes the specified user's expenses <br/>OR <br/>```404```: id not found <br/>OR <br/>```400```: bad request if error |
| ```/user/:id/budgets``` | ```200```: gets the specified user's budgets <br/>OR <br/>404: id not found <br/>OR <br/>```400```: bad request if error | ```405```: method not allowed | ```405```: method not allowed | ```200```: deletes the specified user's budgets <br/>OR <br/>```404```: id not found <br/>OR <br/>```400```: bad request if error |
| ```/group``` | ```405```: method not allowed | ```405```: method not allowed | ```200```: creates a new group <br/>OR <br/>```400```: bad request if error | ```405```: method not allowed |
| ```/group/:id``` | ```200```: gets the group with specified id <br/>OR <br/>```404```: id not found <br/>OR <br/>```400```: bad request if error | ```200```: updates the specified group <br/>OR <br/>```404```: id not found <br/>OR <br/>```400```: bad request if error | ```405```: method not allowed | ```200```: deletes the group with the specified id and all their budgets and expenses <br/>OR <br/>```404```: id not found <br/>OR <br/>```400```: bad request if error |
| ```/group/:id/expenses``` | ```200```: gets the specified groups's expenses <br/>OR <br/>```404```: id not found <br/>OR <br/>```400```: bad request if error | ```405```: method not allowed | ```405```: method not allowed | ```200```: deletes the specified group's expenses <br/>OR <br/>```404```: id not found <br/>OR <br/>```400```: bad request if error |
| ```/group/:id/budgets``` | ```200```: gets the specified group's budgets <br/>OR <br/>```404```: id not found <br/>OR <br/>```400```: bad request if error | ```405```: method not allowed | ```405```: method not allowed | ```200```: deletes the specified group's budgets <br/>OR <br/>```404```: id not found <br/>OR <br/>```400```: bad request if error |
| ```/expense``` | ```405```: method not allowed | ```405```: method not allowed | ```200```: creates a new expense <br/>OR <br/>```400```: bad request if error | ```405```: method not allowed |
| ```/expense?query``` | ```200```: gets the results of the query on all expenses <br/>OR <br/>```400```: bad request if error | ```405```: method not allowed | ```405```: method not allowed | ```405```: method not allowed |
| ```/expense/:id``` | ```200```: gets the expense with specified id <br/>OR <br/>```404```: id not found <br/>OR <br/>```400```: bad request if error | ```200```: updates the specified expense <br/>OR <br/>```404```: id not found <br/>OR <br/>```400```: bad request if error | ```405```: method not allowed | ```200```: deletes the expense with the specified id <br/>OR <br/>```404```: id not found <br/>OR <br/>```400```: bad request if error |
| ```/budget``` | ```405```: method not allowed | ```405```: method not allowed | ```200```: creates a new budget <br/>OR <br/>```400```: bad request if error | ```405```: method not allowed |
| ```/budget?query``` | ```200```: gets the results of the query on all budgets <br/>OR <br/>```400```: bad request if error | ```405```: method not allowed | ```405```: method not allowed | ```405```: method not allowed |
| ```/budget/:id``` | ```200```: gets the budget with specified id <br/>OR <br/>```404```: id not found <br/>OR <br/>```400```: bad request if error | ```200```: updates the specified budget <br/>OR <br/>```404```: id not found <br/>OR <br/>```400```: bad request if error | ```405```: method not allowed | ```200```: deletes the budget with the specified id <br/>OR <br/>```404```: id not found <br/>OR <br/>```400```: bad request if error |

## Representations of Users, Groups, Expenses, and Budgets
A json object will represent each item in the requests made to the server.

### User
```
{
	"id": 128398281,
	"fname": "John",
	"lname": "Doe",
	"email": "jd@hotmail.com",
	"phone": 7185555555
}
```
### Group
```
{
	"id": 0987654321,
	"gname": "The Awesome Group",
	"users": [
		{
			"id": 128398281,
			"fname": "John",
			"lname": "Doe",
			"email": "jd@hotmail.com",
			"phone": 7185555555
		},
		{
			"id": 123456789,
			"fname": "Jane",
			"lname": "Doe",
			"email": "janed@hotmail.com",
			"phone": 71855555555
		}
	]
}
```
### Expense
```
{
	"id": 90909090,
	"amt": 900.00,
	"category": "fight",
	"split": [
		{
			"user": {
				"id": 128398281,
				"fname": "John",
				"lname": "Doe",
				"email": "jd@hotmail.com",
				"phone": 7185555555
			},
			"amt": 0.00
			"percentage": 0.00
		},
		{
			"user": {
				"id": 123456789,
				"fname": "Jane",
				"lname": "Doe",
				"email": "janed@hotmail.com",
				"phone": 71855555555
			},
			"amt": 900.00,
			"percentage": 100.00
		}
	]
}
```
### Budget
```
{
	"id": 01919283757,
	"amt": 900.00,
	"name": "travel",
	"categories": [
		"fight",
		"hotel"
	],
	"group": 0987654321
}
```
