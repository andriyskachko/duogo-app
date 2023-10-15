package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestIsValidEmail(t *testing.T) {
	t.Run("validates correct email", func(t *testing.T) {
		valid, err := IsValidEmail("correct.email123@yahoo.com")

		if err != nil {
			t.Errorf("expected error to be nil, received %+v", err)
		}

		if !valid {
			t.Error("expected email to be valid")
		}
	})

	t.Run("validates incorrect email", func(t *testing.T) {
		email := "hello@p"
		valid, err := IsValidEmail(email)

		if err == nil {
			t.Errorf("expected to error, received nil")
		}

		expectedErrorMessage := fmt.Sprintf("Email %s is invalid", email)
		receivedErrorMessage := err.Error()

		if expectedErrorMessage != receivedErrorMessage {
			t.Errorf("Incorrect error message, expected %q, received %q", expectedErrorMessage, receivedErrorMessage)
		}

		if valid {
			t.Error("expected email to be invalid")
		}
	})
}

func TestIsValidPassword(t *testing.T) {
	testCases := []struct {
		password      string
		expectedValid bool
		expectedError error
	}{
		{"Ab1@cdEf", true, nil},
		{"abc", false, errors.New("Invalid password, less than 7 characters")},
		{"AbcdEfk", false, errors.New("Invalid password, missing special character")},
		{"onlylowercase123@.", false, errors.New("Invalid password, missing uppercase letter")},
		{"ONLYUPPERCASE123@.", false, errors.New("Invalid password, missing lowercase letter")},
		{"1234567@.", false, errors.New("Invalid password, missing uppercase letter")},
		{"@#$%^&", false, errors.New("Invalid password, less than 7 characters")},
	}

	for _, testCase := range testCases {
		valid, err := IsValidPassword(testCase.password)
		if valid != testCase.expectedValid {
			t.Errorf("For password '%s', expected valid: %v, got: %v", testCase.password, testCase.expectedValid, valid)
		}

		if err == nil && testCase.expectedError != nil {
			t.Errorf("For password '%s', expected error: %v, got: %v", testCase.password, testCase.expectedError, err)
		}

		if err != nil && testCase.expectedError == nil {
			t.Errorf("For password '%s', expected no error, but got error: %v", testCase.password, err)
		}

		if err != nil && err.Error() != testCase.expectedError.Error() {
			t.Errorf("For password '%s', expected error message: %v, got: %v", testCase.password, testCase.expectedError, err)
		}
	}
}
