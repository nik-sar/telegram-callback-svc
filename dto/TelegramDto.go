package dto

type UpdateTelegram struct {
	UpdateId int64 `json:"update_id"`
	Message MessageModel `json:"message"`
}

type MessageModel struct {
	Id int64 `json:"message_id"`
	User UserModel `json:"from"`
	Chat ChatModel `json:"chat"`
	Text string `json:"text"`
	Entity EntityModel `json:"entities"`
}

type UserModel struct {
	Id int64 `json:"id"`
	IsBot bool `json:"is_bot"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Username string `json:"username"`
	LanguageCode string `json:"language_code"`
}

type ChatModel struct {
	Id int64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Username string `json:"username"`
	Type string `json:"type"`
}

type EntityModel struct {
	Offset int `json:"offset"`
	Length int `json:"length"`
	Type string `json:"type"`
}

func (dto *UpdateTelegram) GetMessage() *MessageModel {
	if (nil == dto) {
		return nil
	}
	return &dto.Message
}

func (msg *MessageModel) GetEntity() *EntityModel {
	if (nil == msg) {
		return nil
	}
	return &msg.Entity
}