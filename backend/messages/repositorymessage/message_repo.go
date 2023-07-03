package repositorymessage

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/rifki321/warungku/messages/entitymessage"
)

type Message interface {
	PostMesage(ctx context.Context, username string, Message entitymessage.Message, tx *sql.Tx) (entitymessage.Message, error)
	GetMessageByUsername(ctx context.Context, username string, Message entitymessage.Message, db *sql.DB) ([]entitymessage.Message, int, error)
	GetMessageByIdMessage(ctx context.Context, idMessage int, tx *sql.DB) (entitymessage.Message, error)
}

type MessageImpl struct {
}

func NewMessageRepository() *MessageImpl {
	return &MessageImpl{}
}

func (messageRepo *MessageImpl) PostMesage(ctx context.Context, username string, Message entitymessage.Message, tx *sql.Tx) (entitymessage.Message, error) {
	fmt.Println(username)
	fmt.Println(Message.SenderName)
	sqlInsert := "insert into messages(user_name,subject_message,message,sender_name) values(?,?,?,?)"
	result, err := tx.ExecContext(ctx, sqlInsert, Message.SenderName, Message.Subject, Message.Message, username)
	if err != nil {
		return entitymessage.Message{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return entitymessage.Message{}, err
	}
	Message.Id = int(id)
	return Message, nil
}

func (messageRepo *MessageImpl) GetMessageByUsername(ctx context.Context, username string, Message entitymessage.Message, db *sql.DB) ([]entitymessage.Message, int, error) {
	sqlQuery := `SELECT m.message, m.tanggal_message, m.subject_message, u.username
	FROM messages AS m
	JOIN users AS u ON m.user_name = u.username
	WHERE m.sender_name = ?`

	sqlForTotalRow := "SELECT COUNT(*) FROM messages WHERE sender_name = ?"
	var totalRows int
	row, err := db.QueryContext(ctx, sqlForTotalRow, username)
	if err != nil {
		return []entitymessage.Message{}, 0, err
	}
	defer row.Close()

	if row.Next() {
		err := row.Scan(&totalRows)
		if err != nil {
			return []entitymessage.Message{}, 0, err
		}
	}

	rows, err := db.QueryContext(ctx, sqlQuery, username)
	if err != nil {
		return []entitymessage.Message{}, 0, err
	}
	defer rows.Close()

	var messages []entitymessage.Message
	for rows.Next() {
		message := entitymessage.Message{}
		if err := rows.Scan(&message.Message, &message.Datemessage, &message.Subject, &message.User.Username); err != nil {
			return []entitymessage.Message{}, 0, err
		}
		messages = append(messages, message)
	}

	return messages, totalRows, nil
}

func (messageRepo *MessageImpl) GetMessageByIdMessage(ctx context.Context, idMessage int, tx *sql.DB) (entitymessage.Message, error) {
	sqlQuery := "select m.id, m.message,m.subject_message,m.sender_name,m.tanggal_message,u.username from messages as m join users as u on m.user_name = u.username where m.id = ?"
	row, err := tx.QueryContext(ctx, sqlQuery, idMessage)
	if err != nil {
		return entitymessage.Message{}, err
	}
	defer row.Close()
	message := entitymessage.Message{}

	if row.Next() {
		if err := row.Scan(&message.Id, &message.Message, &message.Subject, &message.SenderName, &message.Datemessage, &message.User.Username); err != nil {
			return entitymessage.Message{}, err
		}

	}
	return message, nil

}
