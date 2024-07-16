package internal

import "github.com/google/uuid"

type Task struct {
	id       uuid.UUID
	title    string
	priority string
	status   string
}
