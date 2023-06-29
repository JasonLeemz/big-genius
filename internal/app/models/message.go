package models

import "time"

type Message struct {
	Id        int64     `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	User      string    `gorm:"column:user;NOT NULL"`
	Question  string    `gorm:"column:question;NOT NULL"`
	MessageID string    `gorm:"column:message_id;NOT NULL"`
	Answer    string    `gorm:"column:answer;NOT NULL"`
	Reply     string    `gorm:"column:reply;NOT NULL"`
	TraceId   int64     `gorm:"column:trace_id;default:0;NOT NULL"`
	Ctime     time.Time `gorm:"column:ctime;default:current_timestamp();NOT NULL"`
	Utime     time.Time `gorm:"column:utime;default:current_timestamp();NOT NULL"`
}
