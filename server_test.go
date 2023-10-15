package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

var (
	validEmail      = "correct.email123@yahoo.com"
	validPassword   = "Ab1@cdEf"
	invalidEmail    = "taro@p12"
	invalidPassword = "onlylowercase123@."
)

func TestHandlePOSTUser(t *testing.T) {
	t.Run("user with valid email and password", func(t *testing.T) {
		formData := url.Values{
			"email":    {validEmail},
			"password": {validPassword},
		}

		request := createRequestWithFormData(formData)
		response := httptest.NewRecorder()

		UserServer(response, request)

		gotStatus := response.Code
		expectedStatus := http.StatusFound
		if gotStatus != expectedStatus {
			t.Errorf("Expected status code %d, but got %d", expectedStatus, gotStatus)
		}

		expectedLocation := "/home"
		if gotLocation := response.Header().Get("Location"); gotLocation != expectedLocation {
			t.Errorf("Expected redirect location %s, but got %s", expectedLocation, gotLocation)
		}

		numOfCookies := len(response.Result().Cookies())
		if numOfCookies == 0 {
			t.Errorf("Expected to have auth cookie, got %d cookies", numOfCookies)
		}
	})

	t.Run("user with invalid email", func(t *testing.T) {
		formData := url.Values{
			"email":    {invalidEmail},
			"password": {validPassword},
		}

		request := createRequestWithFormData(formData)
		response := httptest.NewRecorder()

		UserServer(response, request)

		gotStatus := response.Code
		expectedStatus := http.StatusUnprocessableEntity
		if gotStatus != expectedStatus {
			t.Errorf("Expected status code %d, but got %d", expectedStatus, gotStatus)
		}

		expectedErrorMessage := fmt.Sprintf("Email %s is invalid", invalidEmail)
		gotResult := response.Body.String()
		if !strings.Contains(gotResult, expectedErrorMessage) {
			t.Errorf("Expected error message containing %q, but got %q", expectedErrorMessage, gotResult)
		}
	})

	t.Run("user with invalid password", func(t *testing.T) {
		formData := url.Values{
			"email":    {validEmail},
			"password": {invalidPassword},
		}

		request := createRequestWithFormData(formData)
		response := httptest.NewRecorder()

		UserServer(response, request)

		gotStatus := response.Code
		expectedStatus := http.StatusUnprocessableEntity
		if gotStatus != expectedStatus {
			t.Errorf("Expected status code %d, but got %d", expectedStatus, gotStatus)
		}

		expectedErrorMessage := "Invalid password, missing uppercase letter"
		gotResult := response.Body.String()
		if !strings.Contains(gotResult, expectedErrorMessage) {
			t.Errorf("Expected error message containing %q, but got %q", expectedErrorMessage, gotResult)
		}
	})
}

func createRequestWithFormData(formData url.Values) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, "/users", strings.NewReader(formData.Encode()))
	req.Header.Set("Content-Type", ContentTypeForm)
	return req
}
