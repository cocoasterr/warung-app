package models

import (
	"time"

	"github.com/google/uuid"
)

type Product struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
}

func (p *Product) ModelCreate(payload map[string]interface{}) map[string]interface{} {
	if payload["name"] != "" || payload["stock"] != "" {
		payload["id"] = uuid.New().String()
		payload["created_at"] = time.Now()
		payload["created_by"] = "adminwarungapp"
		payload["updated_at"] = time.Now()
		payload["updated_by"] = "adminwarungapp"
		return payload
	}
	return nil
}
func (p *Product) ModelUpdate(payload map[string]interface{}) map[string]interface{} {
	payload["updatedby"] = "adminwarungapp"
	payload["updatedat"] = time.Now()
	return payload
}

func (p *Product) Model() interface{} {
	return &Product{}
}

func (p *Product) TbName() string {
	return "product"
}
