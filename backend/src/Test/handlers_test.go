// handlers_test.go
package handlers

import (
	auth "attorneyManager/_auth"
	contract "attorneyManager/_contract"
	messages "attorneyManager/_messages"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))

	// Check the status code is what we expect.
	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestHandleRegister(t *testing.T) {
	// create json from userLogin
	userRegister := auth.UserRegister{
		Name:     "sample_user",
		Email:    "sample@gmail.com",
		Username: "sample",
		Password: "sample",
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(userRegister)
	if err != nil {
		t.Fatal(err)
	}

	// make request
	req, _ := http.NewRequest("POST", "http://localhost:8080/api/create-account", &buf)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))

	// Check the status code is what we expect.
	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestGetUserEmail(t *testing.T) {
	u, err := url.Parse("http://localhost:8080/api/getuser")
	if err != nil {
		panic(err)
	}

	// make request
	req, _ := http.NewRequest("POST", u.String(), nil)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))

	// Check the status code is what we expect.
	if status := resp.StatusCode; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

func TestHandleLogout(t *testing.T) {

	// make request
	req, _ := http.NewRequest("POST", "http://localhost:8080/api/logout", nil)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))

	// Check the status code is what we expect.
	if status := resp.StatusCode; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

func TestHandleGetContract(t *testing.T) {
	u, err := url.Parse("http://localhost:8080/api/get-contract?attorneyID=Bob@gmail.com&contractID=00000000")
	if err != nil {
		panic(err)
	}

	// make request
	req, _ := http.NewRequest("POST", u.String(), nil)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Check the status code is what we expect.
	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// test contract
	contractTest := contract.Contract{
		ContractID:      "00000000",
		ContractType:    "lease",
		ContractName:    "test.pdf",
		DateCreated:     "3/2/2022",
		TerminationDate: "3/2/2023",
		ValidSigniture:  true,
		PaymentType:     "cash",
		AmountPaid:      0.0,
		AmountOwed:      100.0,
		AttorneyName:    "Bob",
		AttorneyEmail:   "Bob@gmail.com",
		ClientName:      "Alice",
		ClientEmail:     "alice@yahoo.com"}

	// check the response body
	var contract contract.Contract
	err = json.NewDecoder(resp.Body).Decode(&contract)
	if err != nil {
		t.Error("response body difformed")
	}

	// compare with test contract
	if contract != contractTest {
		t.Errorf("response contract does not match the test contract")
	}
}

// akshay
func TestHandleFileUpload(t *testing.T) {
	// request url
	u, err := url.Parse("/api/upload?contract_type=UNIT_TEST&termination_date=00/00/0000&payment_type=GOLD&ammount_paid=1.5&ammount_owed=0&client_email=alice@yahoo.com&client_name=alice")
	if err != nil {
		panic(err)
	}

	// test file to upload

	// make request
	req, _ := http.NewRequest("POST", u.String(), nil)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Check the status code is what we expect.
	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestHandleFileDownload(t *testing.T) {
	u, err := url.Parse("http://localhost:8080/api/download?attorney_email=a@a.a&contract_id=00000001")
	if err != nil {
		panic(err)
	}

	// make request
	req, _ := http.NewRequest("POST", u.String(), nil)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Check the status code is what we expect.
	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// get test file & read file to byte arrray
	testFile, err := ioutil.ReadFile("../contract_store/a@a.a_00000001.pdf")
	if err != nil {
		t.Errorf("Error loading test file")
	}

	// get the response body
	bodyFile, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error loading file from response body")
	}

	// compare with test file
	res := bytes.Compare(testFile, bodyFile)
	if res != 0 {
		t.Errorf("response contract does not match the test contract")
	}
}

func TestHandleCountContracts(t *testing.T) {
	u, err := url.Parse("http://localhost:8080/api/count-contracts?attorney_email=Bob@gmail.com")
	if err != nil {
		panic(err)
	}

	// make request
	req, _ := http.NewRequest("POST", u.String(), nil)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Check the status code is what we expect.
	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error reading response body")
	}
	bodyString := string(bodyBytes)

	// compare with test file
	testStr := "2\n"
	if bodyString != testStr {
		fmt.Println(bodyString)
		fmt.Println(testStr)
		t.Errorf("response contract does not match the test contract")
	}
}

func TestHandleCountMessages(t *testing.T) {
	u, err := url.Parse("http://localhost:8080/api/count-messages?attorney_email=bob@bob.bob")
	if err != nil {
		panic(err)
	}

	// make request
	req, _ := http.NewRequest("POST", u.String(), nil)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Check the status code is what we expect.
	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Error reading response body")
	}
	bodyString := string(bodyBytes)

	// compare with test file
	testStr := "3\n"
	if bodyString != testStr {
		fmt.Println(bodyString)
		fmt.Println(testStr)
		t.Errorf("response contract does not match the test contract")
	}
}

func TestHandleSendMessage(t *testing.T) {

	// create json from userLogin
	message := messages.Message{
		Sender:   "frank@gmail.com",
		Receiver: "john@gmail.com",
		Message:  "UNIT TEST MESSAGE",
		Time:     "1650319345",
	}
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(message)
	if err != nil {
		t.Fatal(err)
	}

	// make request
	req, _ := http.NewRequest("POST", "http://localhost:8080/api/send-message", &buf)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//fmt.Println("response Status:", resp.Status)
	//fmt.Println("response Headers:", resp.Header)
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("response Body:", string(body))

	// Check the status code is what we expect.
	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
