package dao

import (
	"big-genius/internal/app/dto"
	"big-genius/internal/app/models"
	"big-genius/internal/app/models/database"
	"gorm.io/gorm"
)

type MessageDAO struct {
	db *gorm.DB
}

func (dao *MessageDAO) Add(data dto.MessageDTO) (int64, error) {
	m := models.Message{
		User:      data.User,
		Question:  data.Question,
		MessageID: data.MessageID,
		TraceId:   0,
	}
	tx := dao.db.Table(tblMessage).Create(m)

	return tx.RowsAffected, tx.Error
}

func (dao *MessageDAO) Update(data dto.MessageDTO) (int64, error) {

	tx := dao.db.Table(tblMessage).Where("message_id = ?", data.MessageID).Updates(data)
	return tx.RowsAffected, tx.Error
}

func (dao *MessageDAO) Select() []models.Message {
	//TODO implement me
	panic("implement me")
}

func (dao *MessageDAO) Find(m map[string]interface{}) []models.Message {
	//TODO implement me
	panic("implement me")
}

var tblMessage = "message"

func NewMessageDAO() *MessageDAO {
	return &MessageDAO{
		db: database.DB,
	}
}

type Message interface {
	Add(data dto.MessageDTO) (int64, error)
	Update() (int64, error)
	Select() []models.Message
	Find(map[string]interface{}) []models.Message
}
