package models

import (
	"github.com/google/uuid"
)

type CreateBulkAttachmentInput struct {
	DocumentID   uuid.UUID `json:"document_id"`
	DocumentName string    `json:"document_name"`
	OrderID      int       `json:"order_id"`
}
