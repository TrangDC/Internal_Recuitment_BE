package servicebus

import (
	"context"
	"encoding/json"
	"log"
	"time"
	"trec/config"
	"trec/ent"
	"trec/models"
	"trec/repository"

	servicebus "github.com/Azure/azure-service-bus-go"
)

const (
	EmailEventTriggerQueue         = "trec-email-event-trigger"
	EmailEventTriggerCallbackQueue = "trec-email-event-trigger-callback"
	InterviewScheduleQueue         = "trec-interview-schedule"
	InterviewScheduleCallbackQueue = "trec-interview-schedule-callback"
)

type ServiceBus interface {
	ListenToSubscription(messages chan<- models.Messages)
	SendEmailTriggerMessage(ctx context.Context, input models.MessageInput) error
	SendInterviewScheduleMessage(ctx context.Context, input models.MessageInput, schedule time.Time) error
	ProcessMessages(ctx context.Context, messages <-chan models.Messages)
}

type serviceBusImpl struct {
	repository                repository.Repository
	emailEventTrigger         *servicebus.Queue
	emailEventTriggerCallback *servicebus.Queue
	interviewSchedule         *servicebus.Queue
	interviewScheduleCallback *servicebus.Queue
}

func NewServiceBus(config config.ServiceBusConfig, entClient *ent.Client) (ServiceBus, error) {
	repoRegistry := repository.NewRepository(entClient)
	ns, err := servicebus.NewNamespace(servicebus.NamespaceWithConnectionString(config.ConnectionString))
	if err != nil {
		log.Printf("failed to create namespace: %s \n", err)
	}
	emailEventTrigger, err := ns.NewQueue(EmailEventTriggerQueue)
	if err != nil {
		log.Printf("failed to create event trigger queue client: %s \n", err)
	}
	emailEventTriggerCallback, err := ns.NewQueue(EmailEventTriggerCallbackQueue)
	if err != nil {
		log.Printf("failed to create event trigger callback queue client: %s \n", err)
	}
	interviewSchedule, err := ns.NewQueue(InterviewScheduleQueue)
	if err != nil {
		log.Printf("failed to create interview schedule queue client: %s \n", err)
	}
	interviewScheduleCallback, err := ns.NewQueue(InterviewScheduleCallbackQueue)
	if err != nil {
		log.Printf("failed to create interview schedule callback queue client: %s \n", err)
	}
	return &serviceBusImpl{
		repository:                repoRegistry,
		emailEventTrigger:         emailEventTrigger,
		emailEventTriggerCallback: emailEventTriggerCallback,
		interviewSchedule:         interviewSchedule,
		interviewScheduleCallback: interviewScheduleCallback,
	}, nil
}

func (s *serviceBusImpl) ListenToSubscription(messages chan<- models.Messages) {
	for {
		emailEventTriggerCtx := context.Background()
		err := s.emailEventTriggerCallback.Receive(emailEventTriggerCtx, servicebus.HandlerFunc(func(ctx context.Context, msg *servicebus.Message) error {
			messages <- models.Messages{
				Message:   *msg,
				QueueName: EmailEventTriggerCallbackQueue,
			}
			return msg.Complete(ctx)
		}))
		if err != nil {
			log.Printf("failed to receive email trigger callback messages: %s", err)
			time.Sleep(1 * time.Second)
		}
		interviewScheduleQueueCtx := context.Background()
		err = s.interviewScheduleCallback.Receive(interviewScheduleQueueCtx, servicebus.HandlerFunc(func(ctx context.Context, msg *servicebus.Message) error {
			messages <- models.Messages{
				Message:   *msg,
				QueueName: InterviewScheduleCallbackQueue,
			}
			return msg.Complete(ctx)
		}))
		if err != nil {
			log.Printf("failed to receive interview schedule callback messages: %s", err)
			time.Sleep(1 * time.Second)
		}
	}
}

func (s *serviceBusImpl) ProcessMessages(ctx context.Context, messages <-chan models.Messages) {
	for {
		select {
		case msg := <-messages:
			var input models.MessageOutput
			json.Unmarshal(msg.Message.Data, &input)
			switch msg.QueueName {
			case EmailEventTriggerCallbackQueue:
				s.repository.OutgoingEmail().CallbackOutgoingEmail(ctx, input)
			case InterviewScheduleCallbackQueue:
				// Do something with interview schedule callback
				log.Printf("Interview schedule callback: %v", input)
			}
		}
	}
}

func (s *serviceBusImpl) SendEmailTriggerMessage(ctx context.Context, input models.MessageInput) error {
	jsonBytes, err := json.Marshal(input)
	if err != nil {
		return err
	}
	msg := servicebus.NewMessage(jsonBytes)
	sender, err := s.emailEventTrigger.NewSender(ctx)
	if err != nil {
		log.Printf("failed to create sender: %s", err)
	}
	err = s.sentQueue(ctx, msg, sender)
	return err
}

func (s *serviceBusImpl) SendInterviewScheduleMessage(ctx context.Context, input models.MessageInput, schedule time.Time) error {
	jsonBytes, err := json.Marshal(input)
	if err != nil {
		return err
	}
	msg := servicebus.NewMessage(jsonBytes)
	if schedule != (time.Time{}) {
		msg.ScheduleAt(schedule)
	}
	sender, err := s.interviewSchedule.NewSender(ctx)
	if err != nil {
		log.Printf("failed to create sender: %s", err)
	}
	err = s.sentQueue(ctx, msg, sender)
	return err
}

func (s *serviceBusImpl) sentQueue(ctx context.Context, msg *servicebus.Message, sender *servicebus.Sender) error {
	err := sender.Send(ctx, msg)
	if err != nil {
		if err = sender.Recover(ctx); err != nil {
			log.Printf("azure.queue.sender: can't recover sender %T - %s", err, err.Error())
			return err
		}
		if err := sender.Send(ctx, msg); err != nil {
			log.Printf("failed to send message: %s", err)
		}
	}
	return nil
}

// Path: internal/service_bus/service_bus_test.go