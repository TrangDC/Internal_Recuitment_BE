// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (a *AttachmentQuery) CollectFields(ctx context.Context, satisfies ...string) (*AttachmentQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return a, nil
	}
	if err := a.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return a, nil
}

func (a *AttachmentQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "candidateJob":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &CandidateJobQuery{config: a.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			a.withCandidateJob = query
		case "candidateJobFeedback":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &CandidateJobFeedbackQuery{config: a.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			a.withCandidateJobFeedback = query
		case "candidateInterview":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &CandidateInterviewQuery{config: a.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			a.withCandidateInterview = query
		}
	}
	return nil
}

type attachmentPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []AttachmentPaginateOption
}

func newAttachmentPaginateArgs(rv map[string]interface{}) *attachmentPaginateArgs {
	args := &attachmentPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]interface{}:
			var (
				err1, err2 error
				order      = &AttachmentOrder{Field: &AttachmentOrderField{}}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithAttachmentOrder(order))
			}
		case *AttachmentOrder:
			if v != nil {
				args.opts = append(args.opts, WithAttachmentOrder(v))
			}
		}
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (at *AuditTrailQuery) CollectFields(ctx context.Context, satisfies ...string) (*AuditTrailQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return at, nil
	}
	if err := at.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return at, nil
}

func (at *AuditTrailQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "userEdge":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &UserQuery{config: at.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			at.withUserEdge = query
		}
	}
	return nil
}

type audittrailPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []AuditTrailPaginateOption
}

func newAuditTrailPaginateArgs(rv map[string]interface{}) *audittrailPaginateArgs {
	args := &audittrailPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]interface{}:
			var (
				err1, err2 error
				order      = &AuditTrailOrder{Field: &AuditTrailOrderField{}}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithAuditTrailOrder(order))
			}
		case *AuditTrailOrder:
			if v != nil {
				args.opts = append(args.opts, WithAuditTrailOrder(v))
			}
		}
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (c *CandidateQuery) CollectFields(ctx context.Context, satisfies ...string) (*CandidateQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return c, nil
	}
	if err := c.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *CandidateQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "candidateJobEdges":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &CandidateJobQuery{config: c.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			c.WithNamedCandidateJobEdges(alias, func(wq *CandidateJobQuery) {
				*wq = *query
			})
		}
	}
	return nil
}

type candidatePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []CandidatePaginateOption
}

func newCandidatePaginateArgs(rv map[string]interface{}) *candidatePaginateArgs {
	args := &candidatePaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]interface{}:
			var (
				err1, err2 error
				order      = &CandidateOrder{Field: &CandidateOrderField{}}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithCandidateOrder(order))
			}
		case *CandidateOrder:
			if v != nil {
				args.opts = append(args.opts, WithCandidateOrder(v))
			}
		}
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (ci *CandidateInterviewQuery) CollectFields(ctx context.Context, satisfies ...string) (*CandidateInterviewQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return ci, nil
	}
	if err := ci.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return ci, nil
}

func (ci *CandidateInterviewQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "candidateJobEdge":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &CandidateJobQuery{config: ci.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			ci.withCandidateJobEdge = query
		case "attachmentEdges":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &AttachmentQuery{config: ci.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			ci.WithNamedAttachmentEdges(alias, func(wq *AttachmentQuery) {
				*wq = *query
			})
		case "interviewerEdges":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &UserQuery{config: ci.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			ci.WithNamedInterviewerEdges(alias, func(wq *UserQuery) {
				*wq = *query
			})
		case "createdByEdge":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &UserQuery{config: ci.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			ci.withCreatedByEdge = query
		case "userInterviewers":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &CandidateInterviewerQuery{config: ci.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			ci.WithNamedUserInterviewers(alias, func(wq *CandidateInterviewerQuery) {
				*wq = *query
			})
		}
	}
	return nil
}

type candidateinterviewPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []CandidateInterviewPaginateOption
}

func newCandidateInterviewPaginateArgs(rv map[string]interface{}) *candidateinterviewPaginateArgs {
	args := &candidateinterviewPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]interface{}:
			var (
				err1, err2 error
				order      = &CandidateInterviewOrder{Field: &CandidateInterviewOrderField{}}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithCandidateInterviewOrder(order))
			}
		case *CandidateInterviewOrder:
			if v != nil {
				args.opts = append(args.opts, WithCandidateInterviewOrder(v))
			}
		}
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (ci *CandidateInterviewerQuery) CollectFields(ctx context.Context, satisfies ...string) (*CandidateInterviewerQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return ci, nil
	}
	if err := ci.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return ci, nil
}

func (ci *CandidateInterviewerQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "userEdge":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &UserQuery{config: ci.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			ci.withUserEdge = query
		case "interviewEdge":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &CandidateInterviewQuery{config: ci.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			ci.withInterviewEdge = query
		}
	}
	return nil
}

type candidateinterviewerPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []CandidateInterviewerPaginateOption
}

func newCandidateInterviewerPaginateArgs(rv map[string]interface{}) *candidateinterviewerPaginateArgs {
	args := &candidateinterviewerPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]interface{}:
			var (
				err1, err2 error
				order      = &CandidateInterviewerOrder{Field: &CandidateInterviewerOrderField{}}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithCandidateInterviewerOrder(order))
			}
		case *CandidateInterviewerOrder:
			if v != nil {
				args.opts = append(args.opts, WithCandidateInterviewerOrder(v))
			}
		}
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (cj *CandidateJobQuery) CollectFields(ctx context.Context, satisfies ...string) (*CandidateJobQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return cj, nil
	}
	if err := cj.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return cj, nil
}

func (cj *CandidateJobQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "attachmentEdges":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &AttachmentQuery{config: cj.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			cj.WithNamedAttachmentEdges(alias, func(wq *AttachmentQuery) {
				*wq = *query
			})
		case "hiringJobEdge":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &HiringJobQuery{config: cj.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			cj.withHiringJobEdge = query
		case "candidateJobFeedback":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &CandidateJobFeedbackQuery{config: cj.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			cj.WithNamedCandidateJobFeedback(alias, func(wq *CandidateJobFeedbackQuery) {
				*wq = *query
			})
		case "candidateEdge":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &CandidateQuery{config: cj.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			cj.withCandidateEdge = query
		case "candidateJobInterview":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &CandidateInterviewQuery{config: cj.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			cj.WithNamedCandidateJobInterview(alias, func(wq *CandidateInterviewQuery) {
				*wq = *query
			})
		case "createdByEdge":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &UserQuery{config: cj.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			cj.withCreatedByEdge = query
		case "candidateJobStep":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &CandidateJobStepQuery{config: cj.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			cj.WithNamedCandidateJobStep(alias, func(wq *CandidateJobStepQuery) {
				*wq = *query
			})
		}
	}
	return nil
}

type candidatejobPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []CandidateJobPaginateOption
}

func newCandidateJobPaginateArgs(rv map[string]interface{}) *candidatejobPaginateArgs {
	args := &candidatejobPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]interface{}:
			var (
				err1, err2 error
				order      = &CandidateJobOrder{Field: &CandidateJobOrderField{}}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithCandidateJobOrder(order))
			}
		case *CandidateJobOrder:
			if v != nil {
				args.opts = append(args.opts, WithCandidateJobOrder(v))
			}
		}
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (cjf *CandidateJobFeedbackQuery) CollectFields(ctx context.Context, satisfies ...string) (*CandidateJobFeedbackQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return cjf, nil
	}
	if err := cjf.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return cjf, nil
}

func (cjf *CandidateJobFeedbackQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "createdByEdge":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &UserQuery{config: cjf.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			cjf.withCreatedByEdge = query
		case "candidateJobEdge":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &CandidateJobQuery{config: cjf.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			cjf.withCandidateJobEdge = query
		case "attachmentEdges":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &AttachmentQuery{config: cjf.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			cjf.WithNamedAttachmentEdges(alias, func(wq *AttachmentQuery) {
				*wq = *query
			})
		}
	}
	return nil
}

type candidatejobfeedbackPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []CandidateJobFeedbackPaginateOption
}

func newCandidateJobFeedbackPaginateArgs(rv map[string]interface{}) *candidatejobfeedbackPaginateArgs {
	args := &candidatejobfeedbackPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]interface{}:
			var (
				err1, err2 error
				order      = &CandidateJobFeedbackOrder{Field: &CandidateJobFeedbackOrderField{}}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithCandidateJobFeedbackOrder(order))
			}
		case *CandidateJobFeedbackOrder:
			if v != nil {
				args.opts = append(args.opts, WithCandidateJobFeedbackOrder(v))
			}
		}
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (cjs *CandidateJobStepQuery) CollectFields(ctx context.Context, satisfies ...string) (*CandidateJobStepQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return cjs, nil
	}
	if err := cjs.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return cjs, nil
}

func (cjs *CandidateJobStepQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "candidateJobEdge":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &CandidateJobQuery{config: cjs.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			cjs.withCandidateJobEdge = query
		}
	}
	return nil
}

type candidatejobstepPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []CandidateJobStepPaginateOption
}

func newCandidateJobStepPaginateArgs(rv map[string]interface{}) *candidatejobstepPaginateArgs {
	args := &candidatejobstepPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]interface{}:
			var (
				err1, err2 error
				order      = &CandidateJobStepOrder{Field: &CandidateJobStepOrderField{}}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithCandidateJobStepOrder(order))
			}
		case *CandidateJobStepOrder:
			if v != nil {
				args.opts = append(args.opts, WithCandidateJobStepOrder(v))
			}
		}
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (hj *HiringJobQuery) CollectFields(ctx context.Context, satisfies ...string) (*HiringJobQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return hj, nil
	}
	if err := hj.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return hj, nil
}

func (hj *HiringJobQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "ownerEdge":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &UserQuery{config: hj.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			hj.withOwnerEdge = query
		case "teamEdge":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &TeamQuery{config: hj.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			hj.withTeamEdge = query
		case "candidateJobEdges":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &CandidateJobQuery{config: hj.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			hj.WithNamedCandidateJobEdges(alias, func(wq *CandidateJobQuery) {
				*wq = *query
			})
		}
	}
	return nil
}

type hiringjobPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []HiringJobPaginateOption
}

func newHiringJobPaginateArgs(rv map[string]interface{}) *hiringjobPaginateArgs {
	args := &hiringjobPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]interface{}:
			var (
				err1, err2 error
				order      = &HiringJobOrder{Field: &HiringJobOrderField{}}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithHiringJobOrder(order))
			}
		case *HiringJobOrder:
			if v != nil {
				args.opts = append(args.opts, WithHiringJobOrder(v))
			}
		}
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (t *TeamQuery) CollectFields(ctx context.Context, satisfies ...string) (*TeamQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return t, nil
	}
	if err := t.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return t, nil
}

func (t *TeamQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "userEdges":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &UserQuery{config: t.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			t.WithNamedUserEdges(alias, func(wq *UserQuery) {
				*wq = *query
			})
		case "teamJobEdges":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &HiringJobQuery{config: t.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			t.WithNamedTeamJobEdges(alias, func(wq *HiringJobQuery) {
				*wq = *query
			})
		case "userTeams":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &TeamManagerQuery{config: t.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			t.WithNamedUserTeams(alias, func(wq *TeamManagerQuery) {
				*wq = *query
			})
		}
	}
	return nil
}

type teamPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []TeamPaginateOption
}

func newTeamPaginateArgs(rv map[string]interface{}) *teamPaginateArgs {
	args := &teamPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]interface{}:
			var (
				err1, err2 error
				order      = &TeamOrder{Field: &TeamOrderField{}}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithTeamOrder(order))
			}
		case *TeamOrder:
			if v != nil {
				args.opts = append(args.opts, WithTeamOrder(v))
			}
		}
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (tm *TeamManagerQuery) CollectFields(ctx context.Context, satisfies ...string) (*TeamManagerQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return tm, nil
	}
	if err := tm.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return tm, nil
}

func (tm *TeamManagerQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "userEdge":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &UserQuery{config: tm.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			tm.withUserEdge = query
		case "teamEdge":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &TeamQuery{config: tm.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			tm.withTeamEdge = query
		}
	}
	return nil
}

type teammanagerPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []TeamManagerPaginateOption
}

func newTeamManagerPaginateArgs(rv map[string]interface{}) *teammanagerPaginateArgs {
	args := &teammanagerPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]interface{}:
			var (
				err1, err2 error
				order      = &TeamManagerOrder{Field: &TeamManagerOrderField{}}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithTeamManagerOrder(order))
			}
		case *TeamManagerOrder:
			if v != nil {
				args.opts = append(args.opts, WithTeamManagerOrder(v))
			}
		}
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (u *UserQuery) CollectFields(ctx context.Context, satisfies ...string) (*UserQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return u, nil
	}
	if err := u.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return u, nil
}

func (u *UserQuery) collectField(ctx context.Context, op *graphql.OperationContext, field graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	for _, field := range graphql.CollectFields(op, field.Selections, satisfies) {
		switch field.Name {
		case "auditEdge":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &AuditTrailQuery{config: u.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			u.WithNamedAuditEdge(alias, func(wq *AuditTrailQuery) {
				*wq = *query
			})
		case "hiringOwner":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &HiringJobQuery{config: u.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			u.WithNamedHiringOwner(alias, func(wq *HiringJobQuery) {
				*wq = *query
			})
		case "teamEdges":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &TeamQuery{config: u.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			u.WithNamedTeamEdges(alias, func(wq *TeamQuery) {
				*wq = *query
			})
		case "candidateJobFeedback":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &CandidateJobFeedbackQuery{config: u.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			u.WithNamedCandidateJobFeedback(alias, func(wq *CandidateJobFeedbackQuery) {
				*wq = *query
			})
		case "interviewEdges":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &CandidateInterviewQuery{config: u.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			u.WithNamedInterviewEdges(alias, func(wq *CandidateInterviewQuery) {
				*wq = *query
			})
		case "candidateJobEdges":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &CandidateJobQuery{config: u.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			u.WithNamedCandidateJobEdges(alias, func(wq *CandidateJobQuery) {
				*wq = *query
			})
		case "candidateInterviewEdges":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &CandidateInterviewQuery{config: u.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			u.WithNamedCandidateInterviewEdges(alias, func(wq *CandidateInterviewQuery) {
				*wq = *query
			})
		case "teamUsers":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &TeamManagerQuery{config: u.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			u.WithNamedTeamUsers(alias, func(wq *TeamManagerQuery) {
				*wq = *query
			})
		case "interviewUsers":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = &CandidateInterviewerQuery{config: u.config}
			)
			if err := query.collectField(ctx, op, field, path, satisfies...); err != nil {
				return err
			}
			u.WithNamedInterviewUsers(alias, func(wq *CandidateInterviewerQuery) {
				*wq = *query
			})
		}
	}
	return nil
}

type userPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []UserPaginateOption
}

func newUserPaginateArgs(rv map[string]interface{}) *userPaginateArgs {
	args := &userPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]interface{}:
			var (
				err1, err2 error
				order      = &UserOrder{Field: &UserOrderField{}}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithUserOrder(order))
			}
		case *UserOrder:
			if v != nil {
				args.opts = append(args.opts, WithUserOrder(v))
			}
		}
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput interface{}, path ...string) map[string]interface{} {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	for _, name := range path {
		var field *graphql.CollectedField
		for _, f := range graphql.CollectFields(oc, fc.Field.Selections, nil) {
			if f.Alias == name {
				field = &f
				break
			}
		}
		if field == nil {
			return nil
		}
		cf, err := fc.Child(ctx, *field)
		if err != nil {
			args := field.ArgumentMap(oc.Variables)
			return unmarshalArgs(ctx, whereInput, args)
		}
		fc = cf
	}
	return fc.Args
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput interface{}, args map[string]interface{}) map[string]interface{} {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

func limitRows(partitionBy string, limit int, orderBy ...sql.Querier) func(s *sql.Selector) {
	return func(s *sql.Selector) {
		d := sql.Dialect(s.Dialect())
		s.SetDistinct(false)
		with := d.With("src_query").
			As(s.Clone()).
			With("limited_query").
			As(
				d.Select("*").
					AppendSelectExprAs(
						sql.RowNumber().PartitionBy(partitionBy).OrderExpr(orderBy...),
						"row_number",
					).
					From(d.Table("src_query")),
			)
		t := d.Table("limited_query").As(s.TableName())
		*s = *d.Select(s.UnqualifiedColumns()...).
			From(t).
			Where(sql.LTE(t.C("row_number"), limit)).
			Prefix(with)
	}
}
