package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddNumbers(t *testing.T) {
	// Define la solicitud para sumar 2 y 3
	request := AddRequest{Num1: 3, Num2: 3}
	jsonReq, _ := json.Marshal(request)

	// Crea una solicitud HTTP POST con la solicitud JSON
	req, err := http.NewRequest("POST", "/add", bytes.NewBuffer(jsonReq))
	if err != nil {
		t.Fatal(err)
	}

	// Crea un objeto "ResponseRecorder" para recibir la respuesta
	rr := httptest.NewRecorder()

	// Llama a la función "AddNumbers" con la solicitud y el objeto "ResponseRecorder"
	handler := http.HandlerFunc(AddNumbers)
	handler.ServeHTTP(rr, req)

	// Verifica que el código de estado sea 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Decodifica la respuesta JSON y verifica el resultado
	var response AddResponse
	if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
		t.Fatal(err)
	}
	expected := 5
	if response.Result != expected {
		t.Errorf("Handler returned unexpected result: got %v want %v", response.Result, expected)
	}
}
