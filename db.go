package gotrackr

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

type MySQLStorage struct {
	db *sql.DB
}

func NewMySQLStorage(cfg mysql.Config) (*MySQLStorage, error) {
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatalf("[ERROR] Failed to open MySQL connection: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("[ERROR] Failed to ping MySQL: %v", err)
		return nil, err
	}

	log.Println("[INFO] Connected to MySQL successfully")

	return &MySQLStorage{db: db}, nil
}

func (s *MySQLStorage) Init() (*sql.DB, error) {
	// Initialize tables if necessary
	if err := s.createUsersTable(); err != nil {
		return nil, err
	}

	if err := s.createProjectsTable(); err != nil {
		return nil, err
	}

	if err := s.createTasksTable(); err != nil {
		return nil, err
	}

	return s.db, nil
}

func (s *MySQLStorage) createTable(tableName, query string) error {
	_, err := s.db.Exec(query)
	if err != nil {
		log.Printf("[ERROR] Failed to create table %s: %v", tableName, err)
		return err
	}

	log.Printf("[INFO] Table %s created successfully", tableName)
	return nil
}

func (s *MySQLStorage) createUsersTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			user_id INT UNSIGNED NOT NULL AUTO_INCREMENT,
			email VARCHAR(255) NOT NULL,
			first_name VARCHAR(255) NOT NULL,
			last_name VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

			PRIMARY KEY (user_id),
			UNIQUE KEY (email)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	`
	return s.createTable("users", query)
}

func (s *MySQLStorage) createProjectsTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS projects (
			project_id INT UNSIGNED NOT NULL AUTO_INCREMENT,
			name VARCHAR(255) NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

			PRIMARY KEY (project_id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	`
	return s.createTable("projects", query)
}

func (s *MySQLStorage) createTasksTable() error {
	query := `
		CREATE TABLE IF NOT EXISTS tasks (
			task_id INT UNSIGNED NOT NULL AUTO_INCREMENT,
			name VARCHAR(255) NOT NULL,
			status ENUM('TODO', 'IN_PROGRESS', 'IN_TESTING', 'DONE') NOT NULL DEFAULT 'TODO',
			project_id INT UNSIGNED NOT NULL,
			assigned_to INT UNSIGNED NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

			PRIMARY KEY (task_id),
			FOREIGN KEY (assigned_to) REFERENCES users(user_id),
			FOREIGN KEY (project_id) REFERENCES projects(project_id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8;
	`
	return s.createTable("tasks", query)
}
