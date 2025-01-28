package models

import "time"

const (
	statusApproved = "approved"
	statusRejected = "rejected"
	statusPending  = "pending"
)

// Message represents a message that needs approval.
type Message struct {
	ID       string
	Content  string
	Approved *time.Time
	Rejected *time.Time
}

type Response struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	Status  string `json:"status"`
}

func GetResponse(msg *Message) *Response {
	if msg == nil {
		return nil
	}

	status := statusPending
	if msg.Rejected != nil {
		status = statusRejected
	}
	if msg.Approved != nil {
		status = statusApproved
	}

	return &Response{
		ID:      msg.ID,
		Content: msg.Content,
		Status:  status,
	}
}

type Request struct {
	Content string `json:"content"`
}
