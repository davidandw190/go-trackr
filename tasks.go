package gotrackr

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

var errNameRequired = errors.New("name is required")
var errProjectIDRequired = errors.New("project id is required")
var errUserIDRequired = errors.New("user id is required")

type TasksService struct {
	store Store
}

func NewTasksService(s Store) *TasksService {
	return &TasksService{store: s}
}

func (ts *TasksService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/tasks", ts.handleCreateTask).Methods("POST")
	r.HandleFunc("/tasks/{id}", ts.handleGetTask).Methods("GET")

}

func (ts *TasksService) handleCreateTask(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}

	defer r.Body.Close()

	var task *Task
	err = json.Unmarshal(body, &task)
	if err != nil {
		// TODO
		return
	}

	if err = validateTaskPayload(task); err != nil {
		// TODO
		return
	}
}

func (ts *TasksService) handleGetTask(w http.ResponseWriter, r *http.Request) {

}

func validateTaskPayload(task *Task) error {
	if task.Name == "" {
		return errNameRequired
	}

	if task.ProjectID == 0 {
		return errProjectIDRequired
	}

	if task.AssignedTo == 0 {
		return errUserIDRequired
	}

	return nil
}
