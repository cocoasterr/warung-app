package models

import (
	"github.com/google/uuid"
)

type Users struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (p *Users) Model() interface{} {
	return &Users{}
}

func (p *Users) ModelCreate(payload map[string]interface{}) map[string]interface{} {
	payload["id"] = uuid.New().String()
	return payload
}
func (p *Users) ModelUpdate(payload map[string]interface{}) map[string]interface{} {
	return payload
}
func (p *Users) TbName() string {
	return "users"
}
