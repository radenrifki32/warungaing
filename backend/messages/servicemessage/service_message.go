package servicemessage

import (
	"context"
	"database/sql"
	"time"

	"github.com/rifki321/warungku/helper"
	"github.com/rifki321/warungku/messages/entitymessage"
	"github.com/rifki321/warungku/messages/repositorymessage"
)

type RequestMessage struct {
	Message        string `json:"message"`
	SubjectMessage string `json:"subject_message"`
	SenderTo       string `json:"send_to"`
}

type ResponseMessage struct {
	TanggalMessage time.Time `json:"tanggal_message"`
	SendTo         string    `json:"terkirim_ke"`
}

type ResponseMessageByUsername struct {
	Message        string    `json:"message"`
	SubjectMessage string    `json:"subject_message"`
	SendTo         string    `json:"terkirim_ke"`
	Datemessage    time.Time `json:"terkirim_tanggal"`
}
type ResponseMessageByIdMessage struct {
	IdMessage      int       `json:"id_message"`
	Message        string    `json:"message"`
	SubjectMessage string    `json:"subject_message"`
	SendTo         string    `json:"terkirim_ke"`
	Datemessage    time.Time `json:"terkirim_tanggal"`
}
type ServiceMessage interface {
	SendMessage(ctx context.Context, request *RequestMessage, username string) (*ResponseMessage, error)
	GetMessageByUsername(ctx context.Context, username string) ([]ResponseMessageByUsername, int, error)
	GetMessageByIdMessage(ctx context.Context, idMessage int) (ResponseMessageByIdMessage, error)
}

type ServiceMessageImpl struct {
	sql         *sql.DB
	repoMessage repositorymessage.Message
}

func NewServiceMessage(sql *sql.DB, repoMessage repositorymessage.Message) *ServiceMessageImpl {
	return &ServiceMessageImpl{
		sql:         sql,
		repoMessage: repoMessage,
	}
}
func (serviceMessage *ServiceMessageImpl) SendMessage(ctx context.Context, request *RequestMessage, username string) (*ResponseMessage, error) {
	tx, err := serviceMessage.sql.Begin()
	if err != nil {
		return &ResponseMessage{}, err
	}
	defer helper.CommitOrRollback(tx)
	messageRepo := entitymessage.Message{
		Subject:    request.SubjectMessage,
		Message:    request.Message,
		SenderName: request.SenderTo,
	}

	message, err := serviceMessage.repoMessage.PostMesage(ctx, username, messageRepo, tx)
	if err != nil {
		return &ResponseMessage{}, err
	}

	response := &ResponseMessage{
		TanggalMessage: message.Datemessage,
		SendTo:         message.SenderName,
	}
	return response, nil

}
func (serviceMessage *ServiceMessageImpl) GetMessageByUsername(ctx context.Context, username string) ([]ResponseMessageByUsername, int, error) {
	db := serviceMessage.sql

	message := entitymessage.Message{}
	messages, totalRows, err := serviceMessage.repoMessage.GetMessageByUsername(ctx, username, message, db)
	if err != nil {
		return []ResponseMessageByUsername{}, 0, err
	}

	var responseMessages []ResponseMessageByUsername
	for _, msg := range messages {
		responseMsg := ResponseMessageByUsername{
			Message:        msg.Message,
			SubjectMessage: msg.Subject,
			SendTo:         msg.User.Username,
			Datemessage:    msg.Datemessage,
		}
		responseMessages = append(responseMessages, responseMsg)
	}

	return responseMessages, totalRows, nil
}

func (serviceMessage *ServiceMessageImpl) GetMessageByIdMessage(ctx context.Context, idMessage int) (ResponseMessageByIdMessage, error) {
	tx := serviceMessage.sql

	message, err := serviceMessage.repoMessage.GetMessageByIdMessage(ctx, idMessage, tx)
	if err != nil {
		return ResponseMessageByIdMessage{}, err
	}
	responseMsg := ResponseMessageByIdMessage{
		IdMessage:      message.Id,
		Message:        message.Message,
		SubjectMessage: message.Subject,
		SendTo:         message.User.Username,
		Datemessage:    message.Datemessage,
	}
	return responseMsg, nil
}
