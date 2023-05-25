package Test

import (
	"example.com"
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
	main.AddNumbers(recorder, req)
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
