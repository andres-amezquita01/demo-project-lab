package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	// Create a test HTTP request with the necessary parameters.
	form := url.Values{}
	form.Add("num1", "2")
	form.Add("num2", "3")
	req, err := http.NewRequest("POST", "/add", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}
	// Create a test ResponseWriter to capture the response.
	recorder := httptest.NewRecorder()

	// Call the AddNumbers function passing the test Request and ResponseWriter.
	AddNumbers(recorder, req)
	// Check the expected status code.
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, recorder.Code)
	}
	// Check the body of the expected JSON response.
	expectedResponse := `{"result":5}`
	if recorder.Body.String() != expectedResponse {
		t.Errorf("Expected response %s but got %s", expectedResponse, recorder.Body.String())
	}
}

func TestBin(t *testing.T) {
	// Create a test HTTP request with the necessary parameters.
	form := url.Values{}
	form.Add("num1", "2023")
	req, err := http.NewRequest("POST", "/bin", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}
	// Create a test ResponseWriter to capture the response.
	recorder := httptest.NewRecorder()

	// Call the AddNumbers function passing the test Request and ResponseWriter.
	ConvertIntToBinary(recorder, req)
	// Check the expected status code.
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, recorder.Code)
	}
	// Check the body of the expected JSON response.
	expectedResponse := `{"Message":"11111100111"}`
	if recorder.Body.String() != expectedResponse {
		t.Errorf("Expected response %s but got %s", expectedResponse, recorder.Body.String())
	}
}
func TestDiv(t *testing.T) {
	// Create a test HTTP request with the necessary parameters.
	form := url.Values{}
	form.Add("num1", "6")
	form.Add("num2", "2")
	req, err := http.NewRequest("POST", "/div", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}
	// Create a test ResponseWriter to capture the response.
	recorder := httptest.NewRecorder()

	// Call the AddNumbers function passing the test Request and ResponseWriter.
	DivisionNumbers(recorder, req)
	// Check the expected status code.
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, recorder.Code)
	}
	// Check the body of the expected JSON response.
	expectedResponse := `{"result":3}`
	if recorder.Body.String() != expectedResponse {
		t.Errorf("Expected response %s but got %s", expectedResponse, recorder.Body.String())
	}
}
func TestMul(t *testing.T) {
	// Create a test HTTP request with the necessary parameters.
	form := url.Values{}
	form.Add("num1", "6")
	form.Add("num2", "2")
	req, err := http.NewRequest("POST", "/mul", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}
	// Create a test ResponseWriter to capture the response.
	recorder := httptest.NewRecorder()

	// Call the AddNumbers function passing the test Request and ResponseWriter.
	MultiplyNumbers(recorder, req)
	// Check the expected status code.
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, recorder.Code)
	}
	// Check the body of the expected JSON response.
	expectedResponse := `{"result":12}`
	if recorder.Body.String() != expectedResponse {
		t.Errorf("Expected response %s but got %s", expectedResponse, recorder.Body.String())
	}
}
func TestSub(t *testing.T) {
	// Create a test HTTP request with the necessary parameters.
	form := url.Values{}
	form.Add("num1", "6")
	form.Add("num2", "2")
	req, err := http.NewRequest("POST", "/sub", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		t.Fatal(err)
	}
	// Create a test ResponseWriter to capture the response.
	recorder := httptest.NewRecorder()

	// Call the AddNumbers function passing the test Request and ResponseWriter.
	SubtractionNumbers(recorder, req)
	// Check the expected status code.
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, recorder.Code)
	}
	// Check the body of the expected JSON response.
	expectedResponse := `{"result":4}`
	if recorder.Body.String() != expectedResponse {
		t.Errorf("Expected response %s but got %s", expectedResponse, recorder.Body.String())
	}
}
