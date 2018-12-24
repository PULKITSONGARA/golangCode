package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

type Person struct {
	Email_id string
	Password string
}

var a App

func TestMain(m *testing.M) {
	a = App{}
	a.Run()
	code := m.Run()

	// clearTable()

	os.Exit(code)
}
func TestSignUp(t *testing.T) {
	var person Person
	person.Email_id = "p@gmail.com"
	person.Password = "pulkit"

	jsonPerson, err := json.Marshal(person)

	if err != nil {
		t.Fatalf("error aa gyi bhiya")
	}
	req, err := http.NewRequest("POST", "http://localhost/signup", bytes.NewBuffer(jsonPerson))
	if err != nil {
		t.Fatalf("error occur in req")
	}
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)
	fmt.Println(m)
	if m["Email_id"] != "" {
		t.Errorf("Expected the id to remain the same. Got %v", m["Email_id"])
	}
	// rec := httptest.NewRecorder()
	// a.Router.SignUp(rec, req)
	// res := rec.Result()
	// b, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	t.Fatalf("sadjflsajsahgh")
	// }
	// d, err := strcov.Atoi(string(bytes.TrimSpace(b)))
	// if err != nil {
	// 	t.Fatalf("dsfjlfj")
	// }
	// fmt.Println(b)

}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rec := httptest.NewRecorder()

	a.Router.ServeHTTP(rec, req)

	return rec
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
