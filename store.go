package gotrackr

import "database/sql"

// Storage defines methods to interact with the database.
type Storage interface {
	// CreateTask creates a new task in the database.
	CreateTask(t *Task) (*Task, error)
	GetTaskByID(id string) (*Task, error)
	// CreateUser creates a new user in the database.
	CreateUser(u *User) (*User, error)
}

// Store implements the Storage interface.
type Store struct {
	db *sql.DB
}

// NewStore creates a new Store instance.
func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

// CreateTask inserts a new task into the database.
func (s *Store) CreateTask(t *Task) (*Task, error) {
	rows, err := s.db.Exec(`INSERT INTO tasks (name, status, project_id, assigned_to)
	VALUES (?, ?, ?, ?)`, t.Name, t.Status, t.ProjectID, t.AssignedTo)
	if err != nil {
		return nil, err
	}

	id, err := rows.LastInsertId()
	if err != nil {
		return nil, err
	}

	t.ID = id
	return t, nil
}

func (s *Store) GetTaskByID(id string) (*Task, error) {
	var t Task
	err := s.db.QueryRow(`
		SELECT task_id, name, status, project_id, assigned_to, created_at
		FROM tasks
		WHERE task_id = ?`, id).Scan(&t.ID, &t.Name, &t.ProjectID, &t.AssignedTo, &t.CreatedAt)

	return &t, err
}

// CreateUser inserts a new user into the database.
func (s *Store) CreateUser(u *User) (*User, error) {
	// TODO: Implement user creation.
	return nil, nil
}
