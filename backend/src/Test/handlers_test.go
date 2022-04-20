// handlers_test.go
package handlers

import (
	auth "attorneyManager/_auth"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func TestHandleLogin(t *testing.T) {

	// create json from userLogin
	userLogin := auth.UserLogin{
		Email:    "akshay@gmail.com",
		Password: "akshay",
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(userLogin)
	if err != nil {
		t.Fatal(err)
	}

	// make request
	req, _ := http.NewRequest("POST", "http://localhost:8080/api/login", &buf)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	// Check the status code is what we expect.
	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestHandleGetContract(t *testing.T) {
	u, err := url.Parse("http://localhost:8080/api/get-contract?attorneyID=akshay@gmail.com&contractID=00000001")
	if err != nil {
		panic(err)
	}

	// make request
	req, _ := http.NewRequest("POST", u.String(), nil)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	// add query params
	// q := req.URL.Query()
	// q.Add("attorneyID", "akshay@gmail.com")
	// q.Add("contractID", "00000001")
	// req.URL.RawQuery() = q.Encode()

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	// Check the status code is what we expect.
	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
