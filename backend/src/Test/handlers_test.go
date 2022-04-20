// handlers_test.go
package handlers

import (
	auth "attorneyManager/_auth"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
