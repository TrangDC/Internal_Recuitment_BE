package service

import (
	"context"
	"net/http"
	"strings"
	"trec/ent"
	"trec/ent/candidate"
	"trec/ent/outgoingemail"
	"trec/ent/predicate"
	"trec/internal/util"
	"trec/models"
	"trec/repository"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqljson"
	"github.com/google/uuid"
	"github.com/samber/lo"

	"go.uber.org/zap"
)

type OutgoingEmailService interface {
	CreateBulkOutgoingEmail(ctx context.Context, input []models.MessageInput, candidateId uuid.UUID) ([]*ent.OutgoingEmail, error)
	CallbackOutgoingEmail(ctx context.Context, input models.MessageOutput) (*ent.OutgoingEmail, error)
	GetAllOutgoingEmails(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.OutgoingEmailFreeWord, filter ent.OutgoingEmailFilter, orderBy *ent.OutgoingEmailOrder) (*ent.OutgoingEmailResponseGetAll, error)
}

type outgoingEmailSvcImpl struct {
	repoRegistry repository.Repository
	logger       *zap.Logger
}

func NewOutgoingEmailService(repoRegistry repository.Repository, logger *zap.Logger) OutgoingEmailService {
	return &outgoingEmailSvcImpl{
		repoRegistry: repoRegistry,
		logger:       logger,
	}
}

func (svc outgoingEmailSvcImpl) CreateBulkOutgoingEmail(ctx context.Context, input []models.MessageInput, candidateId uuid.UUID) ([]*ent.OutgoingEmail, error) {
	return svc.repoRegistry.OutgoingEmail().CreateBulkOutgoingEmail(ctx, input, candidateId)
}

func (svc outgoingEmailSvcImpl) CallbackOutgoingEmail(ctx context.Context, input models.MessageOutput) (*ent.OutgoingEmail, error) {
	return svc.repoRegistry.OutgoingEmail().CallbackOutgoingEmail(ctx, input)
}

func (svc outgoingEmailSvcImpl) GetAllOutgoingEmails(ctx context.Context, pagination *ent.PaginationInput, freeWord *ent.OutgoingEmailFreeWord, filter ent.OutgoingEmailFilter, orderBy *ent.OutgoingEmailOrder) (*ent.OutgoingEmailResponseGetAll, error) {
	var result *ent.OutgoingEmailResponseGetAll
	var edges []*ent.OutgoingEmailEdge
	var page int
	var perPage int
	query := svc.repoRegistry.OutgoingEmail().BuildQuery()
	svc.filter(ctx, query, filter)
	svc.freeWord(query, freeWord)
	count, err := svc.repoRegistry.OutgoingEmail().BuildCount(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	order := ent.Desc(outgoingemail.FieldCreatedAt)
	if orderBy != nil {
		order = ent.Desc(strings.ToLower(orderBy.Field.String()))
		if orderBy.Direction == ent.OrderDirectionAsc {
			order = ent.Asc(strings.ToLower(orderBy.Field.String()))
		}
	}
	query = query.Order(order)
	if pagination != nil {
		page = *pagination.Page
		perPage = *pagination.PerPage
		query = query.Limit(perPage).Offset((page - 1) * perPage)
	}
	outgoingEmails, err := svc.repoRegistry.OutgoingEmail().BuildList(ctx, query)
	if err != nil {
		svc.logger.Error(err.Error(), zap.Error(err))
		return nil, util.WrapGQLError(ctx, err.Error(), http.StatusInternalServerError, util.ErrorFlagInternalError)
	}
	edges = lo.Map(outgoingEmails, func(v *ent.OutgoingEmail, index int) *ent.OutgoingEmailEdge {
		return &ent.OutgoingEmailEdge{
			Node: v,
			Cursor: ent.Cursor{
				Value: v.ID.String(),
			},
		}
	})
	result = &ent.OutgoingEmailResponseGetAll{
		Edges: edges,
		Pagination: &ent.Pagination{
			Total:   count,
			Page:    page,
			PerPage: perPage,
		},
	}
	return result, nil
}

// common function
func (svc *outgoingEmailSvcImpl) freeWord(outgoingEmailQuery *ent.OutgoingEmailQuery, input *ent.OutgoingEmailFreeWord) {
	var predicates []predicate.OutgoingEmail
	if input != nil {
		if input.Subject != nil {
			predicates = append(predicates, outgoingemail.SubjectContainsFold(strings.TrimSpace(*input.Subject)))
		}
		if input.Content != nil {
			predicates = append(predicates, outgoingemail.ContentContainsFold(strings.TrimSpace(*input.Content)))
		}
	}
	if len(predicates) > 0 {
		outgoingEmailQuery.Where(outgoingemail.Or(predicates...))
	}
}

func (svc *outgoingEmailSvcImpl) filter(ctx context.Context, outgoingEmailQuery *ent.OutgoingEmailQuery, input ent.OutgoingEmailFilter) {
	if input.CandidateID != nil {
		candidateRec, _ := svc.repoRegistry.Candidate().BuildGet(ctx, svc.repoRegistry.Candidate().BuildBaseQuery().
			Where(candidate.ID(uuid.MustParse(*input.CandidateID))))
		outgoingEmailQuery.Where(outgoingemail.CandidateIDEQ(uuid.MustParse(*input.CandidateID)))
		if candidateRec != nil {
			outgoingEmailQuery.Where(outgoingemail.Or(
				func(s *sql.Selector) { s.Where(sqljson.ValueContains(outgoingemail.FieldTo, candidateRec.Email)) },
				func(s *sql.Selector) { s.Where(sqljson.ValueContains(outgoingemail.FieldCc, candidateRec.Email)) },
			))
		}
	}
	if input.Status != nil {
		status := lo.Map(input.Status, func(v ent.OutgoingEmailStatus, index int) outgoingemail.Status {
			return outgoingemail.Status(v.String())
		})
		outgoingEmailQuery.Where(outgoingemail.StatusIn(status...))
	}
	if input.RecipientType != nil {
		recipientType := lo.Map(input.RecipientType, func(v ent.OutgoingEmailRecipientType, index int) outgoingemail.RecipientType {
			return outgoingemail.RecipientType(v.String())
		})
		outgoingEmailQuery.Where(outgoingemail.RecipientTypeIn(recipientType...))
	}
	if (input.FromDate != nil && input.ToDate != nil) && (!input.FromDate.IsZero() && !input.ToDate.IsZero()) {
		outgoingEmailQuery.Where(outgoingemail.CreatedAtGTE(*input.FromDate), outgoingemail.CreatedAtLTE(*input.ToDate))
	}
	if input.EventID != nil {
		outgoingEmailQuery.Where(outgoingemail.EventID(uuid.MustParse(*input.EventID)))
	}
}
