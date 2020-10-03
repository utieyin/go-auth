package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var a App

func TestMain(m *testing.M) {
	var err error
	err = godotenv.Load(os.ExpandEnv(".env"))
	if err != nil {
		log.Fatalf("Error getting env '%v'", err)
	}

	a.Initialize(
		os.Getenv("TestDbUser"),
		os.Getenv("TestDbPassword"),
		os.Getenv("TestDbName"),
		os.Getenv("TestDbPort"),
		"127.0.0.1",
	)
	creatTable()
	code := m.Run()
	clearTable()
	os.Exit(code)

}

func creatTable() {
	a.DB.Migrator().CreateTable(&User{})
}
func ensureTableExists() {
	var tables []string
	rows, _ := a.DB.Raw("SELECT concat(table_schema, '.', table_name)  FROM information_schema.tables WHERE table_schema = ?", "public").Rows()
	defer rows.Close()
	for rows.Next() {
		var tableName string
		rows.Scan(&tableName)
		tables = append(tables, tableName)
	}
	if len(tables) > 0 {
		fmt.Println("Table exists")
	} else {
		fmt.Println("No table")
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM users")
	a.DB.Exec("ALTER SEQUENCE users_id_seq RESTART WITH 1")
}

func dropTable(table string) {
	statement := fmt.Sprintf("DROP TABLE IF EXISTS %s", table)
	if err := a.DB.Exec(statement); err != nil {
		fmt.Println(err.Error)
	}
}

func TestEmptyTable(t *testing.T) {
	clearTable()
	req, _ := http.NewRequest("GET", "/users", nil)
	response := executeRequest(req)
	checkStatusCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected empty array got %s ", body)
	}
}

func TestGetUser(t *testing.T) {
	clearTable()
	req, _ := http.NewRequest("GET", "/users/1", nil)
	response := executeRequest(req)
	checkStatusCode(t, http.StatusOK, response.Code)
}
func TestCreateUser(t *testing.T) {
	clearTable()
	var jsonStr = []byte(`{"email": "kan@gmail.com", "password": ""}`)
	req, _ := http.NewRequest("POST", "/users/1", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	response := executeRequest(req)
	checkStatusCode(t, http.StatusOK, response.Code)
	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["email"] != "kan@gmail.com" {
		t.Errorf("Expected email to be 'kan@gmail.com' but got '%v'", m["email"])
	}

}
func TestUserDelete(t *testing.T) {
	clearTable()
	AddUser()
	req, _ := http.NewRequest("GET", "/user/1", nil)
	response := executeRequest(req)
	checkStatusCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("DELETE", "/user/1", nil)
	response = executeRequest(req)
	checkStatusCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("GET", "/user/1", nil)
	response = executeRequest(req)
	checkStatusCode(t, http.StatusOK, response.Code)
}
func TestUserUpdate(t *testing.T) {
	clearTable()
	_, err := AddUser()
	if err != nil {
		t.Errorf("Expected user object, received '%s' ", err)
	}
	req, _ := http.NewRequest("GET", "/user/1", nil)
	response := executeRequest(req)
	var originalUser map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &originalUser)

	var jsonStr = []byte(`{"email": "uti@uti.com", "password": "test"}`)
	req, _ = http.NewRequest("PUT", "/user/1", bytes.NewBuffer(jsonStr))
	response = executeRequest(req)
	checkStatusCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)
	if originalUser["id"] != m["id"] {
		t.Errorf("Expected initial id '%v' to equal '%v'", originalUser["id"], m["id"])
	}
	if originalUser["email"] != m["email"] {
		t.Errorf("Expected initial email '%v' to equal '%v'", originalUser["email"], m["email"])

	}
	if originalUser["password"] != m["password"] {
		t.Errorf("Password '%v' changed to '%v'", originalUser["password"], m["password"])

	}

}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)
	return rr
}
func checkStatusCode(t *testing.T, expected, got int) {
	if expected != got {
		t.Errorf("Expected %d, got %d ", expected, got)
	}
}
