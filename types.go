package gotrackr

import "time"

type CreateProjectPayload struct {
	Name string `json:"name"`
}

type Project struct {
	ProjectID int64     `json:"projectId"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

type RegisterPayload struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
}

type User struct {
	UserID    int64     `json:"userId"`
	Email     string    `json:"email"`
	FirstName string    `json:"firstNaame"`
	LastName  string    `json:"lastName"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreateTaskPayload struct {
	Name         string `json:"name"`
	ProjectID    int64  `json:"projectId"`
	AssignedToID int64  `json:"assignedTo"`
}

type Task struct {
	ID         int64     `json:"taskId"`
	Name       string    `json:"name"`
	Status     string    `json:"status"`
	ProjectID  int64     `json:"projectId"`
	AssignedTo int64     `json:"assignedTo"`
	CreatedAt  time.Time `json:"createdAt"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
