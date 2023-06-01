package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"os"

	routesAuth "github.com/Laodeus/glt/routes/auth"
)

func TestRegisterUser(t *testing.T) {

	os.Setenv("SECRET", "test")
	os.Setenv("SERVER_PORT", "1234")
	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USER", "usr")
	os.Setenv("DB_PASSWORD", "passwd")
	os.Setenv("DB_NAME", "glt")
	os.Setenv("ADMINER_PORT", "9002")

	user := &routesAuth.DbLogin{
		Login:    "TestUser",
		Password: "TestPassword",
	}
	userJson, _ := json.Marshal(user)
	requestBody := bytes.NewBuffer(userJson)

	req, err := http.NewRequest("POST", "/api/v1/users/register", requestBody)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(routesAuth.RegisterUser)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "User recorded"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestLoginUser(t *testing.T) {
	user := &routesAuth.DbLogin{
		Login:    "TestUser",
		Password: "TestPassword",
	}
	userJson, _ := json.Marshal(user)
	requestBody := bytes.NewBuffer(userJson)

	req, err := http.NewRequest("POST", "/api/v1/users/login", requestBody)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(routesAuth.LoginUser)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusUnauthorized {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusUnauthorized)
	}

	if len(rr.Body.String()) == 0 || rr.Body.String()[:7] != "Bearer " {
		t.Errorf("handler returned unexpected body: got %v want a string that starts with 'Bearer '",
			rr.Body.String())
	}
}
