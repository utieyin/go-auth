package utils

import (
	"fmt"


	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func TestPasswordHash(t *testing.T) {
	result, err := HashPassword("tony")
	got := []byte(result)
	name := []byte("tony")
	err = bcrypt.CompareHashAndPassword(got, name)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}

func TestTokenValidation(t *testing.T) {
	req, err := http.NewRequest("GET", "https://graph.facebook.com/debug_token", nil)
	if err != nil {
		log.Println("Something went wrong ", err)
	}
	q := req.URL.Query()
	q.Add("input_token", "something")
	q.Add("access_token", "something")
	req.URL.RawQuery = q.Encode()
	fmt.Println(req.URL.String())
	response := executeRequest(req)
	log.Print(response.Body)
	got := TokenValidation("is_valid")
	wanted := "is_valid"
	if got != wanted {
		t.Errorf("Expected %v, got %s", wanted, got)
	}

}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router := *mux.NewRouter()
	router.ServeHTTP(rr, req)
	return rr
}

func TestFacebookToken(t *testing.T) {
	server := httptest.NewServer(http.HandleFunc(func(rw http.ResponseWriter, req *http.Request) {
		equals(t, req.URL.String(), "/some/path")
		rw.Write([]byte(`OK`))
	}))
	defer server.Close()
	api := API{server.Client(), server.URL}
	body, err := api.DoStuff()
	ok(t, err)
	equals(t, []byte(`OK`), body)
}

func ok(tb testing.TB, err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d: unexpected error: %s\033[39m\n\n", filepath.Base(file), line, err.Error())
		tb.FailNow()
	}
}

// equals fails the test if exp is not equal to act.
func equals(tb testing.TB, exp, act interface{}) {
	if !reflect.DeepEqual(exp, act) {
		_, file, line, _ := runtime.Caller(1)
		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n", filepath.Base(file), line, exp, act)
		tb.FailNow()
	}
}
