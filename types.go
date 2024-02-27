package gotrackr

import "time"

type CreateProjectPayload struct {
	Name string `json:"name"`
}

type Project struct {
	ProjectID int64     `json:"project_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type RegisterPayload struct {
	Email     string `json:"email"`
	FirstName string `json:"first_ame"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
}

type User struct {
	UserID    int64     `json:"id"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreateTaskPayload struct {
	Name         string `json:"name"`
	ProjectID    int64  `json:"project_id"`
	AssignedToID int64  `json:"assigned_to"`
}

type Task struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Status     string    `json:"status"`
	ProjectID  int64     `json:"project_id"`
	AssignedTo int64     `json:"assigned_to"`
	CreatedAt  time.Time `json:"created_at"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
