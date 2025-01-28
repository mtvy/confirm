package repository

import (
	"database/sql"

	"github.com/jackc/pgx"
	_ "github.com/lib/pq"
	"github.com/mtvy/confirm/internal/models"
	"github.com/pkg/errors"
)

type MessageRepo struct {
	conn *pgx.Conn
}

func NewMessageRepo(conn *pgx.Conn) *MessageRepo {
	return &MessageRepo{conn: conn}
}

// Save saves a new message or updates an existing one.
func (repo *MessageRepo) Save(message *models.Message) error {
	_, err := repo.conn.Exec(`INSERT INTO messages (id, content) VALUES ($1, $2)`,
		message.ID, message.Content)

	return err
}

// FindByID retrieves a message by its ID.
func (repo *MessageRepo) FindByID(id string) (*models.Message, error) {
	msg := &models.Message{}
	row := repo.conn.QueryRow("SELECT id, content, approved_at, rejected_at FROM messages WHERE id = $1", id)
	if err := row.Scan(&msg.ID, &msg.Content, &msg.Approved, &msg.Rejected); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("message not found")
		}
		return nil, errors.Wrap(err, "scan")
	}

	return msg, nil
}

// Delete removes a message by its ID.
func (repo *MessageRepo) Rejected(id string) error {
	_, err := repo.conn.Exec("UPDATE messages SET rejected_at = NOW() WHERE id = $1", id)
	return err
}

// Delete removes a message by its ID.
func (repo *MessageRepo) Approved(id string) error {
	_, err := repo.conn.Exec("UPDATE messages SET approved_at = NOW() WHERE id = $1", id)
	return err
}
