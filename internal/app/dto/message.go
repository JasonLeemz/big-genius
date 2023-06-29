package dto

type MessageDTO struct {
	ID        int64  `json:"id,omitempty"`
	User      string `json:"user,omitempty"`
	Question  string `json:"question,omitempty"`
	MessageID string `json:"message_id,omitempty"`
	Answer    string `json:"answer,omitempty"`
	Reply     string `json:"reply,omitempty"`
	TraceId   int64  `json:"trace_id,omitempty"`
	Ctime     string `json:"ctime,omitempty"`
	Utime     string `json:"utime,omitempty"`
}
