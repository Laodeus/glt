package apiCall

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func Login(username, password string) (string, error) {

	apiAddress := os.Getenv("API_ADDRESS")
	apiPort := os.Getenv("SERVER_PORT")
	apiEndpoint := "/api/v1/users/login"
	enpoint := fmt.Sprintf("%s:%s%s", apiAddress, apiPort, apiEndpoint)

	fmt.Println(enpoint)

	requestBody, err := json.Marshal(map[string]string{
		"username": username,
		"password": password,
	})
	if err != nil {
		return "", err
	}

	// Make the API request
	resp, err := http.Post(fmt.Sprintf(enpoint), "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	var responseBody map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&responseBody)
	if err != nil {
		return "", err
	}

	// Process the API response and return the result
	result := fmt.Sprintf("Status: %s, Message: %s", resp.Status, responseBody["message"])
	return result, nil
}
