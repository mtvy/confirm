package repository

import "github.com/mtvy/confirm/internal/models"

// Messenger defines methods for interacting with messages.
type Messenger interface {
	Save(message *models.Message) error
	FindByID(id string) (*models.Message, error)
	Rejected(id string) error
	Approved(id string) error
}
