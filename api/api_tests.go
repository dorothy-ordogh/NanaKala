package api

import (
	"fmt"
	"net/http/httptest"
	"net/http"
	"strings"
	"testing"
)

var server *httptest.Server 
	
func init() {
	server = httptest.NewServer(Handlers())
}

/*===================================================================================*/


/*===================================================================================*/


/*===================================================================================*/



/*===================================================================================*/
func TestCreateExpense(t *testing.T) {
	expenseUrl := fmt.Sprintf("%s/expense", server.URL)
	expenseJson := `{"amt": 120.00, "name":"money", "category": 8, "split":[{"user": {"id":5, "fname":"Dot", "lname":"Orgado", "email":"dot@hotmail.com", "phone":"9999319990"},"splitamt": 60.00}, {"user": {"id": 7, "fname":"Shris", "lname":"Home", "email":"shris@hotmail.com", "phone":"1112342222"}, "splitamt": 60.00}]}`
	// expenseJson := `{}`
	reader := strings.NewReader(expenseJson)

	request, err := http.NewRequest("POST", expenseUrl, reader)

	res, err := http.DefaultClient.Do(request)
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