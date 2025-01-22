package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	//"os"
	"testing"

	logger "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

const baseURL = "http://localhost:8080/api/applicants"
var sample_id string
/*var logger = logger.New()

func init() {
	// Set up logger to write to a file
	file, err := os.OpenFile("test_logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Failed to open log file: %v\n", err)
		return
	}
	logger.Out = file
	logger.SetFormatter(&logger.TextFormatter{
		FullTimestamp: true,
	})
}*/
// Sample Applicant JSON
var sampleApplicant = map[string]interface{}{
	"id":                "e1f456d2-4877-4ffb-9f66-f248b2569d28",
	"nric":              "S8279633T",
	"name":              "Jessie",
	"employment_status": "unemployed",
	"sex":               "female",
	"date_of_birth":     "14-06-1991",
	"household":         "",
}

// Test GET /api/applicants
func TestGetApplicants(t *testing.T) {
	resp, err := http.Get(baseURL)
	assert.NoError(t, err, "GET request failed")
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected status 200")

	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err, "Failed to read response body")

	var applicants []map[string]interface{}
	err = json.Unmarshal(body, &applicants)
	assert.NoError(t, err, "Failed to parse JSON response")

	logger.WithFields(logger.Fields{
		"endpoint": "GET /api/applicants",
		"response": applicants,
	}).Info("GET /api/applicants success")
}

// Test POST /api/applicants
func TestCreateApplicant(t *testing.T) {
	applicantJSON, err := json.Marshal(sampleApplicant)
	assert.NoError(t, err, "Failed to marshal JSON")

	resp, err := http.Post(baseURL, "application/json", bytes.NewBuffer(applicantJSON))
	assert.NoError(t, err, "POST request failed")
	defer resp.Body.Close()

	assert.Equal(t, http.StatusCreated, resp.StatusCode, "Expected status 201")
	
	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err, "Failed to read response body")

	var createdApplicant map[string]interface{}
	err = json.Unmarshal(body, &createdApplicant)
	assert.NoError(t, err, "Failed to parse JSON response")
	sample_id = createdApplicant["id"].(string)

	logger.WithFields(logger.Fields{
		"endpoint": "POST /api/applicants",
		"request":  sampleApplicant,
		"response": createdApplicant,
	}).Info("POST /api/applicants success")
}

// Test PUT /api/applicants/{id}
func TestUpdateApplicant(t *testing.T) {
	updateData := map[string]interface{}{

	"nric":              "S9139327T",
	"name":              "Jonathan",
	"employment_status": "employed",
	"sex":               "male",
	"date_of_birth":     "02-11-1990",
	"household":         "",
	}
	updateJSON, err := json.Marshal(updateData)
	assert.NoError(t, err, "Failed to marshal JSON")

	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/%s", baseURL, sample_id), bytes.NewBuffer(updateJSON))
	assert.NoError(t, err, "Failed to create PUT request")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(t, err, "PUT request failed")
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected status 200")

	body, err := ioutil.ReadAll(resp.Body)
	assert.NoError(t, err, "Failed to read response body")

	var updatedApplicant map[string]interface{}
	err = json.Unmarshal(body, &updatedApplicant)
	assert.NoError(t, err, "Failed to parse JSON response")

	logger.WithFields(logger.Fields{
		"endpoint": "PUT /api/applicants/{id}",
		"request":  updateData,
		"response": updatedApplicant,
	}).Info("PUT /api/applicants/{id} success")
}

// Test DELETE /api/applicants/{id}
func TestDeleteApplicant(t *testing.T) {
	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/%s", baseURL,sample_id), nil)
	assert.NoError(t, err, "Failed to create DELETE request")

	client := &http.Client{}
	resp, err := client.Do(req)
	assert.NoError(t, err, "DELETE request failed")
	defer resp.Body.Close()

	assert.Equal(t, http.StatusNoContent, resp.StatusCode, "Expected status 204")

	logger.WithFields(logger.Fields{
		"endpoint": "DELETE /api/applicants/{id}",
		"id":       sampleApplicant["id"],
	}).Info("DELETE /api/applicants/{id} success")
}
