// Code generated by ent, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"
	"trec/ent"
)

// The AttachmentFunc type is an adapter to allow the use of ordinary
// function as Attachment mutator.
type AttachmentFunc func(context.Context, *ent.AttachmentMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AttachmentFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AttachmentMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AttachmentMutation", m)
	}
	return f(ctx, mv)
}

// The AuditTrailFunc type is an adapter to allow the use of ordinary
// function as AuditTrail mutator.
type AuditTrailFunc func(context.Context, *ent.AuditTrailMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AuditTrailFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AuditTrailMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AuditTrailMutation", m)
	}
	return f(ctx, mv)
}

// The CandidateFunc type is an adapter to allow the use of ordinary
// function as Candidate mutator.
type CandidateFunc func(context.Context, *ent.CandidateMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CandidateFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CandidateMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CandidateMutation", m)
	}
	return f(ctx, mv)
}

// The CandidateAwardFunc type is an adapter to allow the use of ordinary
// function as CandidateAward mutator.
type CandidateAwardFunc func(context.Context, *ent.CandidateAwardMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CandidateAwardFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CandidateAwardMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CandidateAwardMutation", m)
	}
	return f(ctx, mv)
}

// The CandidateCertificateFunc type is an adapter to allow the use of ordinary
// function as CandidateCertificate mutator.
type CandidateCertificateFunc func(context.Context, *ent.CandidateCertificateMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CandidateCertificateFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CandidateCertificateMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CandidateCertificateMutation", m)
	}
	return f(ctx, mv)
}

// The CandidateEducateFunc type is an adapter to allow the use of ordinary
// function as CandidateEducate mutator.
type CandidateEducateFunc func(context.Context, *ent.CandidateEducateMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CandidateEducateFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CandidateEducateMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CandidateEducateMutation", m)
	}
	return f(ctx, mv)
}

// The CandidateExpFunc type is an adapter to allow the use of ordinary
// function as CandidateExp mutator.
type CandidateExpFunc func(context.Context, *ent.CandidateExpMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CandidateExpFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CandidateExpMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CandidateExpMutation", m)
	}
	return f(ctx, mv)
}

// The CandidateHistoryCallFunc type is an adapter to allow the use of ordinary
// function as CandidateHistoryCall mutator.
type CandidateHistoryCallFunc func(context.Context, *ent.CandidateHistoryCallMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CandidateHistoryCallFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CandidateHistoryCallMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CandidateHistoryCallMutation", m)
	}
	return f(ctx, mv)
}

// The CandidateInterviewFunc type is an adapter to allow the use of ordinary
// function as CandidateInterview mutator.
type CandidateInterviewFunc func(context.Context, *ent.CandidateInterviewMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CandidateInterviewFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CandidateInterviewMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CandidateInterviewMutation", m)
	}
	return f(ctx, mv)
}

// The CandidateInterviewerFunc type is an adapter to allow the use of ordinary
// function as CandidateInterviewer mutator.
type CandidateInterviewerFunc func(context.Context, *ent.CandidateInterviewerMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CandidateInterviewerFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CandidateInterviewerMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CandidateInterviewerMutation", m)
	}
	return f(ctx, mv)
}

// The CandidateJobFunc type is an adapter to allow the use of ordinary
// function as CandidateJob mutator.
type CandidateJobFunc func(context.Context, *ent.CandidateJobMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CandidateJobFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CandidateJobMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CandidateJobMutation", m)
	}
	return f(ctx, mv)
}

// The CandidateJobFeedbackFunc type is an adapter to allow the use of ordinary
// function as CandidateJobFeedback mutator.
type CandidateJobFeedbackFunc func(context.Context, *ent.CandidateJobFeedbackMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CandidateJobFeedbackFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CandidateJobFeedbackMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CandidateJobFeedbackMutation", m)
	}
	return f(ctx, mv)
}

// The CandidateJobStepFunc type is an adapter to allow the use of ordinary
// function as CandidateJobStep mutator.
type CandidateJobStepFunc func(context.Context, *ent.CandidateJobStepMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CandidateJobStepFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CandidateJobStepMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CandidateJobStepMutation", m)
	}
	return f(ctx, mv)
}

// The CandidateNoteFunc type is an adapter to allow the use of ordinary
// function as CandidateNote mutator.
type CandidateNoteFunc func(context.Context, *ent.CandidateNoteMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CandidateNoteFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CandidateNoteMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CandidateNoteMutation", m)
	}
	return f(ctx, mv)
}

// The EmailRoleAttributeFunc type is an adapter to allow the use of ordinary
// function as EmailRoleAttribute mutator.
type EmailRoleAttributeFunc func(context.Context, *ent.EmailRoleAttributeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f EmailRoleAttributeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.EmailRoleAttributeMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.EmailRoleAttributeMutation", m)
	}
	return f(ctx, mv)
}

// The EmailTemplateFunc type is an adapter to allow the use of ordinary
// function as EmailTemplate mutator.
type EmailTemplateFunc func(context.Context, *ent.EmailTemplateMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f EmailTemplateFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.EmailTemplateMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.EmailTemplateMutation", m)
	}
	return f(ctx, mv)
}

// The EntityPermissionFunc type is an adapter to allow the use of ordinary
// function as EntityPermission mutator.
type EntityPermissionFunc func(context.Context, *ent.EntityPermissionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f EntityPermissionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.EntityPermissionMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.EntityPermissionMutation", m)
	}
	return f(ctx, mv)
}

// The EntitySkillFunc type is an adapter to allow the use of ordinary
// function as EntitySkill mutator.
type EntitySkillFunc func(context.Context, *ent.EntitySkillMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f EntitySkillFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.EntitySkillMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.EntitySkillMutation", m)
	}
	return f(ctx, mv)
}

// The HiringJobFunc type is an adapter to allow the use of ordinary
// function as HiringJob mutator.
type HiringJobFunc func(context.Context, *ent.HiringJobMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f HiringJobFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.HiringJobMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.HiringJobMutation", m)
	}
	return f(ctx, mv)
}

// The HiringJobStepFunc type is an adapter to allow the use of ordinary
// function as HiringJobStep mutator.
type HiringJobStepFunc func(context.Context, *ent.HiringJobStepMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f HiringJobStepFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.HiringJobStepMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.HiringJobStepMutation", m)
	}
	return f(ctx, mv)
}

// The HiringTeamFunc type is an adapter to allow the use of ordinary
// function as HiringTeam mutator.
type HiringTeamFunc func(context.Context, *ent.HiringTeamMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f HiringTeamFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.HiringTeamMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.HiringTeamMutation", m)
	}
	return f(ctx, mv)
}

// The HiringTeamApproverFunc type is an adapter to allow the use of ordinary
// function as HiringTeamApprover mutator.
type HiringTeamApproverFunc func(context.Context, *ent.HiringTeamApproverMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f HiringTeamApproverFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.HiringTeamApproverMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.HiringTeamApproverMutation", m)
	}
	return f(ctx, mv)
}

// The HiringTeamManagerFunc type is an adapter to allow the use of ordinary
// function as HiringTeamManager mutator.
type HiringTeamManagerFunc func(context.Context, *ent.HiringTeamManagerMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f HiringTeamManagerFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.HiringTeamManagerMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.HiringTeamManagerMutation", m)
	}
	return f(ctx, mv)
}

// The JobPositionFunc type is an adapter to allow the use of ordinary
// function as JobPosition mutator.
type JobPositionFunc func(context.Context, *ent.JobPositionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f JobPositionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.JobPositionMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.JobPositionMutation", m)
	}
	return f(ctx, mv)
}

// The OutgoingEmailFunc type is an adapter to allow the use of ordinary
// function as OutgoingEmail mutator.
type OutgoingEmailFunc func(context.Context, *ent.OutgoingEmailMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f OutgoingEmailFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.OutgoingEmailMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.OutgoingEmailMutation", m)
	}
	return f(ctx, mv)
}

// The PermissionFunc type is an adapter to allow the use of ordinary
// function as Permission mutator.
type PermissionFunc func(context.Context, *ent.PermissionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PermissionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PermissionMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PermissionMutation", m)
	}
	return f(ctx, mv)
}

// The PermissionGroupFunc type is an adapter to allow the use of ordinary
// function as PermissionGroup mutator.
type PermissionGroupFunc func(context.Context, *ent.PermissionGroupMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PermissionGroupFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PermissionGroupMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PermissionGroupMutation", m)
	}
	return f(ctx, mv)
}

// The RecTeamFunc type is an adapter to allow the use of ordinary
// function as RecTeam mutator.
type RecTeamFunc func(context.Context, *ent.RecTeamMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RecTeamFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.RecTeamMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RecTeamMutation", m)
	}
	return f(ctx, mv)
}

// The RoleFunc type is an adapter to allow the use of ordinary
// function as Role mutator.
type RoleFunc func(context.Context, *ent.RoleMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f RoleFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.RoleMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.RoleMutation", m)
	}
	return f(ctx, mv)
}

// The SkillFunc type is an adapter to allow the use of ordinary
// function as Skill mutator.
type SkillFunc func(context.Context, *ent.SkillMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SkillFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.SkillMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SkillMutation", m)
	}
	return f(ctx, mv)
}

// The SkillTypeFunc type is an adapter to allow the use of ordinary
// function as SkillType mutator.
type SkillTypeFunc func(context.Context, *ent.SkillTypeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f SkillTypeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.SkillTypeMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.SkillTypeMutation", m)
	}
	return f(ctx, mv)
}

// The UserFunc type is an adapter to allow the use of ordinary
// function as User mutator.
type UserFunc func(context.Context, *ent.UserMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.UserMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserMutation", m)
	}
	return f(ctx, mv)
}

// The UserRoleFunc type is an adapter to allow the use of ordinary
// function as UserRole mutator.
type UserRoleFunc func(context.Context, *ent.UserRoleMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UserRoleFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.UserRoleMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UserRoleMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// FixedError is a hook returning a fixed error.
func FixedError(err error) ent.Hook {
	return func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(context.Context, ent.Mutation) (ent.Value, error) {
			return nil, err
		})
	}
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
func Reject(op ent.Op) ent.Hook {
	hk := FixedError(fmt.Errorf("%s operation is not allowed", op))
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
