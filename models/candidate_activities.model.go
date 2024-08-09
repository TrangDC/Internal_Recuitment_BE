package models

import (
	"time"

	"github.com/google/uuid"
)

type CandidateActivityReference struct {
	Id        uuid.UUID
	CreatedAt time.Time
}
