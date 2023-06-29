package services

import (
	"big-genius/core/utils/wechat"
	"big-genius/internal/app/dao"
	"big-genius/internal/app/dto"
)

type MessageService struct {
	MessageDAO *dao.MessageDAO
}

var theMessageService = new(MessageService)

func NewMessageService() *MessageService {
	if theMessageService.MessageDAO == nil {
		theMessageService.MessageDAO = dao.NewMessageDAO()
	}
	return theMessageService
}

func (s *MessageService) GetMessageDetailByUser(user string) ([]dto.MessageDTO, error) {

	return nil, nil
}

func (s *MessageService) GetMessageDetailByDate(date string) ([]dto.MessageDTO, error) {

	return nil, nil
}

func (s *MessageService) GetMessageDetailByPage(pageNo int64) ([]dto.MessageDTO, error) {

	return nil, nil
}

func (s *MessageService) Ask(msg wechat.MsgContent) (int64, error) {
	data := dto.MessageDTO{
		User:      msg.ToUsername, // 在controller出转成了入参的fromuser
		Question:  msg.Content,
		MessageID: msg.Msgid,
	}
	return s.MessageDAO.Add(data)

}

func (s *MessageService) Answer(msg wechat.MsgContent) (int64, error) {
	data := dto.MessageDTO{
		MessageID: msg.Msgid,
		Answer:    msg.Content,
	}
	return s.MessageDAO.Update(data)

}
