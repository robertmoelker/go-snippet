package internal

import (
	"database/sql"

	"github.com/google/uuid"
)

type Task struct {
	id       uuid.UUID
	title    string
	priority string
	status   string
}

type TaskModel struct {
	DB *sql.DB
}

func (m *TaskModel) Insert(title, priority, status string) (uuid.UUID, error) {
	id := uuid.New()
	_, err := m.DB.Exec("INSERT INTO tasks (id, title, priority, status) VALUES($1, $2, $3, $4)", id, title, priority, status)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func (m *TaskModel) Get(id uuid.UUID) (*Task, error) {
	t := &Task{}
	err := m.DB.QueryRow("SELECT id, title, priority, status FROM tasks WHERE id = $1", id).Scan(&t.id, &t.title, &t.priority, &t.status)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (m *TaskModel) Latest() ([]*Task, error) {
	rows, err := m.DB.Query("SELECT id, title, priority, status FROM tasks ORDER BY created DESC")
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	tasks := []*Task{}

	for rows.Next() {
		t := &Task{}
		err := rows.Scan(&t.id, &t.title, &t.priority, &t.status)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
