package api

import (
	"fmt"
	"net/http/httptest"
	"net/http"
	"strings"
	"testing"
	"github.com/dorothy-ordogh/NanaKala/models"
)

var server *httptest.Server 
	
func init() {
	server = httptest.NewServer(Handlers())
}

/*===================================================================================*/
// USER

func TestCreateUser(t *testing.T) {
	usrUrl := fmt.Sprintf("%s/user", server.URL)
	usrJson := `{ "fname": "John", "lname": "Doe", "phone":"5555555555", "email": "jd@hotmail.com"}`
	
	reader := strings.NewReader(usrJson)

	request, err := http.NewRequest("POST", usrUrl, reader)

	res, err := http.DefaultClient.Do(request)
	checkTestErr(err, t)

	userid := models.GetUserID(res)

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}

	//============ DELETE ==============//
	usrUrl = fmt.Sprintf("%s/user/%v", server.URL, userid)
	usrJson = `{ }`
	
	reader = strings.NewReader(usrJson)

	request, err = http.NewRequest("DELETE", usrUrl, reader)

	res, err = http.DefaultClient.Do(request)
	checkTestErr(err, t)

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}

func TestUser(t *testing.T) {
	//=========== POST ================//
	usrUrl := fmt.Sprintf("%s/user", server.URL)
	usrJson := `{ "fname": "John", "lname": "Doe", "phone":"5555555555", "email": "jd@hotmail.com"}`
	
	reader := strings.NewReader(usrJson)

	request, err := http.NewRequest("POST", usrUrl, reader)

	res, err := http.DefaultClient.Do(request)
	checkTestErr(err, t)

	userid := models.GetUserID(res)

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}

	//========== GET =================//
	usrUrl = fmt.Sprintf("%s/user/%v", server.URL, userid)

	request, err = http.NewRequest("GET", usrUrl, nil)

	res, err = http.DefaultClient.Do(request)
	checkTestErr(err, t)

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}

	//============ PUT =================//
	usrUrl = fmt.Sprintf("%s/user/%v", server.URL, userid)
	usrJson = `{ "fname": "John", "lname": "Doe", "phone":"1111111111", "email": "johndoe@hotmail.com"}`
	
	reader = strings.NewReader(usrJson)

	request, err = http.NewRequest("PUT", usrUrl, reader)

	res, err = http.DefaultClient.Do(request)
	checkTestErr(err, t)

	userid = models.GetUserID(res)

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}

	//============ DELETE ==============//
	usrUrl = fmt.Sprintf("%s/user/%v", server.URL, userid)
	usrJson = `{ "fname": "John", "lname": "Doe", "phone":"1111111111", "email": "johndoe@hotmail.com"}`
	
	reader = strings.NewReader(usrJson)

	request, err = http.NewRequest("DELETE", usrUrl, reader)

	res, err = http.DefaultClient.Do(request)
	checkTestErr(err, t)

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}

/*===================================================================================*/
// GROUP

func TestCreateGroup(t *testing.T) {
	// first create users to include in group
	usrUrl := fmt.Sprintf("%s/user", server.URL)
	usrJson1 := '{ "fname": "John", "lname": "Doe", "phone":"5555555555", "email": "jd@hotmail.com"}'
	
	reader := strings.NewReader(usrJson1)

	request, err := http.NewRequest("POST", usrUrl, reader)

	res, err := http.DefaultClient.Do(request)
	checkTestErr(err, t)

	userid1 := models.GetUserID(res)

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}

	usrUrl = fmt.Sprintf("%s/user", server.URL)
	usrJson2 := '{ "fname": "Jane", "lname": "Doe", "phone":"5555555555", "email": "jd@hotmail.com"}'
	
	reader = strings.NewReader(usrJson2)

	request, err = http.NewRequest("POST", usrUrl, reader)

	res, err = http.DefaultClient.Do(request)
	checkTestErr(err, t)

	userid2 := models.GetUserID(res)

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}

	// Now create group with users
	groupUrl := fmt.Sprintf("%s/group", server.URL)
	groupJson := fmt.Sprintf("{ 'gname': 'mortgage', [{'id':%v, 'fname': 'John', 'lname': 'Doe', 'phone':'5555555555', 'email': 'jd@hotmail.com'}, {'id':%v, 'fname': 'Jane', 'lname': 'Doe', 'phone':'5555555555', 'email': 'jd@hotmail.com'}]}", userid1, userid2)
	
	reader := strings.NewReader(groupJson)

	request, err := http.NewRequest("POST", groupUrl, reader)

	res, err := http.DefaultClient.Do(request)
	checkTestErr(err, t)

	groupid := models.GetUserID(res)

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}

	//============ DELETE ==============//
	groupUrl = fmt.Sprintf("%s/group/%v", server.URL, groupid)
	groupJson = `{ }`
	
	reader = strings.NewReader(groupJson)

	request, err = http.NewRequest("DELETE", groupUrl, reader)

	res, err = http.DefaultClient.Do(request)
	checkTestErr(err, t)

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}

func TestGroup(t *testing.T) {
	//=========== POST ================//
	// first create users to include in group
	usrUrl := fmt.Sprintf("%s/user", server.URL)
	usrJson1 := '{ "fname": "John", "lname": "Doe", "phone":"5555555555", "email": "jd@hotmail.com"}'
	
	reader := strings.NewReader(usrJson1)

	request, err := http.NewRequest("POST", usrUrl, reader)

	res, err := http.DefaultClient.Do(request)
	checkTestErr(err, t)

	userid1 := models.GetUserID(res)

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}

	usrUrl = fmt.Sprintf("%s/user", server.URL)
	usrJson2 := '{ "fname": "Jane", "lname": "Doe", "phone":"5555555555", "email": "jd@hotmail.com"}'
	
	reader = strings.NewReader(usrJson2)

	request, err = http.NewRequest("POST", usrUrl, reader)

	res, err = http.DefaultClient.Do(request)
	checkTestErr(err, t)

	userid2 := models.GetUserID(res)

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}

	// Now create group with users
	groupUrl := fmt.Sprintf("%s/group", server.URL)
	groupJson := fmt.Sprintf("{ 'gname': 'mortgage', [{'id':%v, 'fname': 'John', 'lname': 'Doe', 'phone':'5555555555', 'email': 'jd@hotmail.com'}, {'id':%v, 'fname': 'Jane', 'lname': 'Doe', 'phone':'5555555555', 'email': 'jd@hotmail.com'}]}", userid1, userid2)
	
	reader := strings.NewReader(groupJson)

	request, err := http.NewRequest("POST", groupUrl, reader)

	res, err := http.DefaultClient.Do(request)
	checkTestErr(err, t)

	groupid := models.GetUserID(res)

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}

	//========== GET =================//
	groupUrl = fmt.Sprintf("%s/group/%v", server.URL, groupid)

	request, err = http.NewRequest("GET", groupUrl, nil)

	res, err = http.DefaultClient.Do(request)
	checkTestErr(err, t)

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}

	//============ PUT =================//
	groupUrl = fmt.Sprintf("%s/group/%v", server.URL, groupid)
	groupJson = `{ "fname": "John", "lname": "Doe", "phone":"1111111111", "email": "johndoe@hotmail.com"}`
	
	reader = strings.NewReader(groupJson)

	request, err = http.NewRequest("PUT", groupUrl, reader)

	res, err = http.DefaultClient.Do(request)
	checkTestErr(err, t)

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}

	//============ DELETE ==============//
	groupUrl = fmt.Sprintf("%s/group/%v", server.URL, groupid)
	groupJson = `{ "fname": "John", "lname": "Doe", "phone":"1111111111", "email": "johndoe@hotmail.com"}`
	
	reader = strings.NewReader(groupJson)

	request, err = http.NewRequest("DELETE", groupUrl, reader)

	res, err = http.DefaultClient.Do(request)
	checkTestErr(err, t)

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}

/*===================================================================================*/
// BUDGET

func TestCreateBudget(t *testing.T) {
	// first create users to include in budget
	usrUrl := fmt.Sprintf("%s/user", server.URL)
	usrJson1 := '{ "fname": "John", "lname": "Doe", "phone":"5555555555", "email": "jd@hotmail.com"}'
	
	reader := strings.NewReader(usrJson1)

	request, err := http.NewRequest("POST", usrUrl, reader)

	res, err := http.DefaultClient.Do(request)
	checkTestErr(err, t)

	userid := models.GetUserID(res)

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
	// now create budget for user
	budgetUrl := fmt.Sprintf("%s/budget", server.URL)
	budgetJson := fmt.Sprintf("{'amt': 120.00, 'name':'money', 'uid':%v}", userid)
	
	reader = strings.NewReader(budgetJson)

	request, err = http.NewRequest("POST", budgetUrl, reader)

	res, err = http.DefaultClient.Do(request)
	checkTestErr(err, t)

	bid := models.GetExpenseID(res)

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}

	//============ DELETE ==============//
	budgetUrl = fmt.Sprintf("%s/budget/%v", server.URL, bid)
	budgetJson = `{ }`
	
	reader = strings.NewReader(budgetJson)

	request, err = http.NewRequest("DELETE", budgetUrl, reader)

	res, err = http.DefaultClient.Do(request)
	checkTestErr(err, t)

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}

func TestBudget(t *testing.T) {
	//=========== POST ================//
	// first create users to include in budget
	usrUrl := fmt.Sprintf("%s/user", server.URL)
	usrJson1 := '{ "fname": "John", "lname": "Doe", "phone":"5555555555", "email": "jd@hotmail.com"}'
	
	reader := strings.NewReader(usrJson1)

	request, err := http.NewRequest("POST", usrUrl, reader)

	res, err := http.DefaultClient.Do(request)
	checkTestErr(err, t)

	userid := models.GetUserID(res)

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
	// now create budget for user
	budgetUrl := fmt.Sprintf("%s/budget", server.URL)
	budgetJson := fmt.Sprintf("{'amt': 120.00, 'name':'Road Trip', 'uid':%v}", userid)
	
	reader = strings.NewReader(budgetJson)

	request, err = http.NewRequest("POST", budgetUrl, reader)

	res, err = http.DefaultClient.Do(request)
	checkTestErr(err, t)

	bid := models.GetExpenseID(res)

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}

	//========== GET =================//
	budgetUrl = fmt.Sprintf("%s/budget/%v", server.URL, bid)

	request, err = http.NewRequest("GET", budgetUrl, nil)

	res, err = http.DefaultClient.Do(request)
	checkTestErr(err, t)

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}

	//============ PUT =================//
	budgetUrl = fmt.Sprintf("%s/budget/%v", server.URL, bid)
	budgetJson = `{"name": "Trippin on the road", "amt": 100.00, "category": 3}`
	
	reader = strings.NewReader(budgetJson)

	request, err = http.NewRequest("PUT", budgetUrl, reader)

	res, err = http.DefaultClient.Do(request)
	checkTestErr(err, t)

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}

	//============ DELETE ==============//
	budgetUrl = fmt.Sprintf("%s/budget/%v", server.URL, bid)
	budgetJson = `{"name": "Trippin on the road", "amt": 100.00, "category": 3}`
	
	reader = strings.NewReader(budgetJson)

	request, err = http.NewRequest("DELETE", budgetUrl, reader)

	res, err = http.DefaultClient.Do(request)
	checkTestErr(err, t)

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}

/*===================================================================================*/
// EXPENSE

func TestCreateExpense(t *testing.T) {
	expenseUrl := fmt.Sprintf("%s/expense", server.URL)
	expenseJson := `{"amt": 120.00, "name":"money", "category": 8, "split":[{"user": {"id":5, "fname":"Dot", "lname":"Orgado", "email":"dot@hotmail.com", "phone":"9999319990"},"splitamt": 60.00}, {"user": {"id": 7, "fname":"Shris", "lname":"Home", "email":"shris@hotmail.com", "phone":"1112342222"}, "splitamt": 60.00}]}`
	
	reader := strings.NewReader(expenseJson)

	request, err := http.NewRequest("POST", expenseUrl, reader)

	res, err := http.DefaultClient.Do(request)
	checkTestErr(err, t)

	expenseID := models.GetExpenseID(res)

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}

	//============ DELETE ==============//
	expenseUrl = fmt.Sprintf("%s/expense/%v", server.URL, expenseID)
	expenseJson = `{ }`
	
	reader = strings.NewReader(expenseJson)

	request, err = http.NewRequest("DELETE", expenseUrl, reader)

	res, err = http.DefaultClient.Do(request)
	checkTestErr(err, t)

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}

func TestExpense(t *testing.T) {
	//=========== POST ================//
	expenseUrl := fmt.Sprintf("%s/expense", server.URL)
	expenseJson := `{"amt": 120.00, "name":"Euro-trip", "category": 8, "gid": 9}`
	
	reader := strings.NewReader(expenseJson)

	request, err := http.NewRequest("POST", expenseUrl, reader)

	res, err := http.DefaultClient.Do(request)
	checkTestErr(err, t)

	expenseID := models.GetExpenseID(res)

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}

	//========== GET =================//
	expenseUrl = fmt.Sprintf("%s/expense/%v", server.URL, expenseID)

	request, err = http.NewRequest("GET", expenseUrl, nil)

	res, err = http.DefaultClient.Do(request)
	checkTestErr(err, t)

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}

	//============ PUT =================//
	expenseUrl = fmt.Sprintf("%s/expense/%v", server.URL, expenseID)
	expenseJson = `{"amt": 10.00, "name":"Euro-trip", "category": 3, "gid": 9}`
	
	reader = strings.NewReader(expenseJson)

	request, err = http.NewRequest("PUT", expenseUrl, reader)

	res, err = http.DefaultClient.Do(request)
	checkTestErr(err, t)

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}

	//============ DELETE ==============//
	expenseUrl = fmt.Sprintf("%s/expense/%v", server.URL, expenseID)
	expenseJson = `{"amt": 120.00, "name":"Euro-trip", "category": 8, "gid": 9}`
	
	reader = strings.NewReader(expenseJson)

	request, err = http.NewRequest("DELETE", expenseUrl, reader)

	res, err = http.DefaultClient.Do(request)
	checkTestErr(err, t)

	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}

func checkTestErr(err error, t *testing.T) {
	if err != nil {
		t.Error(err)
	}
}