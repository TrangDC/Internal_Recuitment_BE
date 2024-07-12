package models

import (
	"trec/ent"

	servicebus "github.com/Azure/azure-service-bus-go"
	"github.com/google/uuid"
)

type MessageInput struct {
	ID              uuid.UUID `json:"id"`
	To              []string  `json:"to"`
	Cc              []string  `json:"cc"`
	Bcc             []string  `json:"bcc"`
	Subject         string    `json:"subject"`
	Content         string    `json:"content"`
	Signature       string    `json:"signature"`
	ApplicationType string    `json:"applicationType"`
	Application     string    `json:"application"`
	ApplicationName string    `json:"applicationName"`
	QueueName       string    `json:"queueName"`
}

type Messages struct {
	Message   servicebus.Message
	QueueName string
}

type GroupModule struct {
	Interview    *ent.CandidateInterview
	Candidate    *ent.Candidate
	Team         *ent.Team
	HiringJob    *ent.HiringJob
	CandidateJob *ent.CandidateJob
}

type MessageOutput struct {
	ID        string `json:"id"`
	IsSuccess bool   `json:"isSuccess"`
	QueueName string `json:"queueName"`
}
