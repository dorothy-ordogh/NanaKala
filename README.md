# NanaKala
What is NanaKala? In Hawaiian, "nana" is to watch or observe, and "kala" is money. This project is for UBC's CPSC 448 (directed studies) course. I proposed to develop a REST API which would help track, and split expenses amongst users and groups, as well as budget money. 

### Setting up
#### Setting up the Server
NanaKala runs on Go 1.5.3 and uses MySql 5.7.11. Please make sure your system is running these versions to guarantee successful integration.

Once your system is running the right version, fork the repo and make sure to get the following libraries:
The Gorilla Web Toolkit by running 
```
$ go get github.com/gorilla/mux
```
and the MySql driver for Go by running
```
$ go get github.com/go-sql-driver/mysql
```

Now you are ready to set up the database.

#### Setting up the Database
I have provided scripts that will set up all the tables you will need to run the code provided. These scripts are located in the folder "scripts" under "dbsetup.sql".

There is one thing you will need to complete manually: the setting up of the categories the budgets and expenses can belong to. The table must have some content in order to work, so I have included one row in its creation, but please make sure to customize it in the way you see fit. Of course, because the categories being sent to the server are numerical values, these are the integer IDs for the category, I suggest you have some kind of enumeration on the front end to make selecting categories more user-friendly.

In order to easily accomplish adding categories to the table, I have included a sample script called "addcategories.sql" in the scripts folder that can be altered and used.

### Components of the system:
- User
- Group (an extension of the user model with one or several users)
- Expense
- Budget (contains a desired spending limit)

Budgets and expenses can apply to a single user, a group, a user within a group, a subset of users in a group or any combination. 

|   Resource   |   GET   |   PUT   |   POST   |  DELETE   |  NOTES  |
|----|----|----|----|----|----|
| ```/user``` | ```405```: method not allowed | ```405```: method not allowed | ```200```: creates a new user <br/>OR <br/>```400```: bad request if error | ```405```: method not allowed | Does not need an ID when posting. Will respond with full user doc with ID |
| ```/user/:id``` | ```200```: gets the user with specified id <br/>OR <br/>```400```: bad request if error | ```200```: updates the specified user <br/>OR <br/>```400```: bad request if error | ```405```: method not allowed | ```200```: deletes the user with the specified id and all their budgets and expenses <br/> OR <br/>```400```: bad request if error |  |
| ```/user/:id/expenses``` | ```200```: gets the specified user's expenses <br/>OR <br/>```400```: bad request if error | ```405```: method not allowed | ```405```: method not allowed | ```200```: deletes the specified user's expenses <br/>OR <br/>```400```: bad request if error |  |
| ```/user/:id/budgets``` | ```200```: gets the specified user's budgets <br/>OR <br/>```400```: bad request if error | ```405```: method not allowed | ```405```: method not allowed | ```200```: deletes the specified user's budgets <br/>OR <br/>```400```: bad request if error |  |
| ```/group``` | ```405```: method not allowed | ```405```: method not allowed | ```200```: creates a new group <br/>OR <br/>```400```: bad request if error | ```405```: method not allowed | Does not need an ID when posting. Will respond with full group doc with ID |
| ```/group/:id``` | ```200```: gets the group with specified id <br/>OR <br/>```400```: bad request if error | ```200```: updates the specified group <br/>OR <br/>```400```: bad request if error | ```405```: method not allowed | ```200```: deletes the group with the specified id and all their budgets and expenses <br/>OR <br/>```400```: bad request if error |
| ```/group/:id/expenses``` | ```200```: gets the specified groups's expenses <br/>OR <br/>```400```: bad request if error | ```405```: method not allowed | ```405```: method not allowed | ```200```: deletes the specified group's expenses <br/>OR <br/>```400```: bad request if error |
| ```/group/:id/budgets``` | ```200```: gets the specified group's budgets <br/>OR <br/>```400```: bad request if error | ```405```: method not allowed | ```405```: method not allowed | ```200```: deletes the specified group's budgets <br/>OR <br/>```400```: bad request if error |
| ```/expense``` | ```405```: method not allowed | ```405```: method not allowed | ```200```: creates a new expense <br/>OR <br/>```400```: bad request if error | ```405```: method not allowed |
| ```/expense/:id``` | ```200```: gets the expense with specified id <br/>OR <br/>```400```: bad request if error | ```200```: updates the specified expense <br/>OR <br/>```400```: bad request if error | ```405```: method not allowed | ```200```: deletes the expense with the specified id <br/>OR <br/>```400```: bad request if error |
| ```/budget``` | ```405```: method not allowed | ```405```: method not allowed | ```200```: creates a new budget <br/>OR <br/>```400```: bad request if error | ```405```: method not allowed |
| ```/budget/:id``` | ```200```: gets the budget with specified id <br/>OR <br/>```400```: bad request if error | ```200```: updates the specified budget <br/>OR <br/>```400```: bad request if error | ```405```: method not allowed | ```200```: deletes the budget with the specified id <br/>OR <br/>```400```: bad request if error |

### Representations of Users, Groups, Expenses, and Budgets
A json object will represent each item in the requests made to the server.

#### User
Users work as you would expect them to. On posting a user, you will not have the ID for that user because it is automatically generated by the database. It will be returned in the response to the post request.
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
The users field in the group objects represents the users that are in the group. Again, the ID is created on inserting into the database, and will be returned in the response to the post request. 
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
A note on how expenses work: 
If you are submitting an expense for a single user, you must add the user information and the information required in the split object (i.e. user and amount) in order to correctly assign the information to that one user. The same goes for several users. 

If you are submitting an expense for a group, you must include the "gid" field with the group's ID and then the expense will be assigned to the group BUT NOT TO EACH INDIVIDUAL USER IN THE GROUP.

The "budgetid" field is to associate the expense with a certian budget identified by its ID.
```
{
	"id": 90909090,
	"amt": 900.00,
	"category": 9,
	"name": "flight to Spain",
	"budgetid": 8,
	"split": [
		{
			"user": {
				"id": 128398281,
				"fname": "John",
				"lname": "Doe",
				"email": "jd@hotmail.com",
				"phone": 7185555555
			},
			"amt": 50.00
		},
		{
			"user": {
				"id": 123456789,
				"fname": "Jane",
				"lname": "Doe",
				"email": "janed@hotmail.com",
				"phone": 71855555555
			},
			"amt": 850.00,
		}
	]
}
```
or for a group expense:
```
{
	"id": 90909090,
	"amt": 900.00,
	"category": 9,
	"name": "flight to Spain",
	"budgetid": 8,
	"gid":0987654321
}
```
### Budget
Budgets work as you would expect them to. 

Note:
If you include a "uid" field with a value, the budget will be associated with a user, and if you include the "gid" field, it will be associated with a group. IT CANNOT BE ASSOCIATED WITH BOTH AT THE SAME TIME. 
```
{
	"id": 128398281,
	"amt": 900.00,
	"name": "travel",
	"category": 8
	"uid": 0987654321
}
```
or for a group budget:
```
{
	"id": 01919283757,
	"amt": 900.00,
	"name": "travel",
	"category": 9,
	"gid": 0987654321
}
```
