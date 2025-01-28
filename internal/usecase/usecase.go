package usecase

import (
	"github.com/gofrs/uuid"
	"github.com/mtvy/confirm/internal/models"
	"github.com/mtvy/confirm/internal/repository"
	"github.com/pkg/errors"
)

// MessageUsecase defines methods for managing messages.
type MessageUsecase struct {
	repo repository.Messenger
}

func NewMessageUsecase(repo repository.Messenger) *MessageUsecase {
	return &MessageUsecase{repo: repo}
}

// SendMessage saves a new message for approval.
func (uc *MessageUsecase) SendMessage(req *models.Request) (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", errors.New("uuid gen error")
	}
	msg := &models.Message{
		ID:      id.String(),
		Content: req.Content,
	}
	if err := uc.repo.Save(msg); err != nil {
		return "", errors.Wrap(err, "save message")
	}

	return id.String(), nil
}

// ApproveMessage approves a message by its ID.
func (uc *MessageUsecase) ApproveMessage(id string) (*models.Message, error) {
	msg, err := uc.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if msg.Rejected != nil {
		return msg, errors.New("message rejected")
	}
	if msg.Approved != nil {
		return msg, errors.New("message approved")
	}
	if err := uc.repo.Approved(id); err != nil {
		return nil, errors.Wrap(err, "update message")
	}
	msg, err = uc.repo.FindByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "find message")
	}
	return msg, nil
}

// RejectMessage deletes a message by its ID.
func (uc *MessageUsecase) RejectMessage(id string) (*models.Message, error) {
	msg, err := uc.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if msg.Approved != nil {
		return msg, errors.New("message approved")
	}
	if msg.Rejected != nil {
		return msg, errors.New("message rejected")
	}
	if err := uc.repo.Rejected(id); err != nil {
		return nil, errors.Wrap(err, "update message")
	}
	msg, err = uc.repo.FindByID(id)
	if err != nil {
		return nil, errors.Wrap(err, "find message")
	}
	return msg, nil
}

func (uc *MessageUsecase) GetMessage(id string) (*models.Message, error) {
	return uc.repo.FindByID(id)
}
