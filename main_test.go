package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test POST and GET routes
func TestAPI(t *testing.T) {
	router := SetupRouter() // Use the exported function

	// Create a test user payload
	userPayload := map[string]interface{}{
		"name": "Sheila",
		"age":  25,
	}
	jsonPayload, _ := json.Marshal(userPayload)

	// Test POST request
	req, _ := http.NewRequest("POST", "/post", bytes.NewBuffer(jsonPayload))
	req.Header.Set("Content-Type", "application/json")
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("POST request failed: expected status 200, got %d", resp.Code)
	}

	// Extract the user ID from response
	var result map[string]string
	json.Unmarshal(resp.Body.Bytes(), &result)
	userID, exists := result["id"]
	if !exists {
		t.Fatal("Failed to get user ID from POST response")
	}

	// Test GET request
	req, _ = http.NewRequest("GET", "/get/"+userID, nil)
	resp = httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("GET request failed: expected status 200, got %d", resp.Code)
	}
}
