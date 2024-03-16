package gotrackr

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestCreateTask(t *testing.T) {
	ms := &MockStore{}
	service := NewTasksService(ms)

	t.Run("should return error if name is empty", func(t *testing.T) {
		payload := &CreateTaskPayload{
			Name: "",
		}

		testCreateTaskWithPayload(t, service, payload, http.StatusBadRequest, errNameRequired.Error())
	})

	t.Run("should return error if project ID is zero", func(t *testing.T) {
		payload := &CreateTaskPayload{
			Name:       "Task with zero project ID",
			ProjectID:  0,
			AssignedTo: 42,
		}

		testCreateTaskWithPayload(t, service, payload, http.StatusBadRequest, errProjectIDRequired.Error())
	})

	t.Run("should return error if assigned to user ID is zero", func(t *testing.T) {
		payload := &CreateTaskPayload{
			Name:       "Task with zero assigned user ID",
			ProjectID:  1,
			AssignedTo: 0,
		}

		testCreateTaskWithPayload(t, service, payload, http.StatusBadRequest, errUserIDRequired.Error())
	})

	t.Run("should create a task", func(t *testing.T) {
		payload := &CreateTaskPayload{
			Name:       "Creating a REST API in Go",
			ProjectID:  2,
			AssignedTo: 1,
		}

		testCreateTaskWithPayload(t, service, payload, http.StatusCreated, "")
	})
}

func testCreateTaskWithPayload(t *testing.T, service *TasksService, payload *CreateTaskPayload, expectedStatusCode int, expectedErrorMessage string) {
	b, err := json.Marshal(payload)
	if err != nil {
		t.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(b))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()

	router.HandleFunc("/tasks", service.handleCreateTask)
	router.ServeHTTP(rr, req)

	if rr.Code != expectedStatusCode {
		t.Errorf("expected status code %d, got %d", expectedStatusCode, rr.Code)
	}

	if expectedStatusCode == http.StatusBadRequest {
		var response ErrorResponse
		err = json.NewDecoder(rr.Body).Decode(&response)
		if err != nil {
			t.Fatal(err)
		}

		if response.Error != expectedErrorMessage {
			t.Errorf("expected error message %s, got %s", expectedErrorMessage, response.Error)
		}
	}
}
