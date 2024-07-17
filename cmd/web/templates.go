package main

import "github.com/robertmoelker/lets-go/internal/models"

type templateData struct {
	Task *models.Task
	Tasks []*models.Task
}
