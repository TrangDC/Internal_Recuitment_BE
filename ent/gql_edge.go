// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (a *Attachment) CandidateJob(ctx context.Context) (*CandidateJob, error) {
	result, err := a.Edges.CandidateJobOrErr()
	if IsNotLoaded(err) {
		result, err = a.QueryCandidateJob().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (a *Attachment) CandidateJobFeedback(ctx context.Context) (*CandidateJobFeedback, error) {
	result, err := a.Edges.CandidateJobFeedbackOrErr()
	if IsNotLoaded(err) {
		result, err = a.QueryCandidateJobFeedback().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (a *Attachment) CandidateInterview(ctx context.Context) (*CandidateInterview, error) {
	result, err := a.Edges.CandidateInterviewOrErr()
	if IsNotLoaded(err) {
		result, err = a.QueryCandidateInterview().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (at *AuditTrail) UserEdge(ctx context.Context) (*User, error) {
	result, err := at.Edges.UserEdgeOrErr()
	if IsNotLoaded(err) {
		result, err = at.QueryUserEdge().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (c *Candidate) CandidateJobEdges(ctx context.Context) (result []*CandidateJob, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedCandidateJobEdges(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.CandidateJobEdgesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryCandidateJobEdges().All(ctx)
	}
	return result, err
}

func (ci *CandidateInterview) CandidateJobEdge(ctx context.Context) (*CandidateJob, error) {
	result, err := ci.Edges.CandidateJobEdgeOrErr()
	if IsNotLoaded(err) {
		result, err = ci.QueryCandidateJobEdge().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (ci *CandidateInterview) AttachmentEdges(ctx context.Context) (result []*Attachment, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = ci.NamedAttachmentEdges(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = ci.Edges.AttachmentEdgesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = ci.QueryAttachmentEdges().All(ctx)
	}
	return result, err
}

func (ci *CandidateInterview) InterviewerEdges(ctx context.Context) (result []*User, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = ci.NamedInterviewerEdges(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = ci.Edges.InterviewerEdgesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = ci.QueryInterviewerEdges().All(ctx)
	}
	return result, err
}

func (ci *CandidateInterview) CreatedByEdge(ctx context.Context) (*User, error) {
	result, err := ci.Edges.CreatedByEdgeOrErr()
	if IsNotLoaded(err) {
		result, err = ci.QueryCreatedByEdge().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (ci *CandidateInterview) UserInterviewers(ctx context.Context) (result []*CandidateInterviewer, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = ci.NamedUserInterviewers(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = ci.Edges.UserInterviewersOrErr()
	}
	if IsNotLoaded(err) {
		result, err = ci.QueryUserInterviewers().All(ctx)
	}
	return result, err
}

func (ci *CandidateInterviewer) UserEdge(ctx context.Context) (*User, error) {
	result, err := ci.Edges.UserEdgeOrErr()
	if IsNotLoaded(err) {
		result, err = ci.QueryUserEdge().Only(ctx)
	}
	return result, err
}

func (ci *CandidateInterviewer) InterviewEdge(ctx context.Context) (*CandidateInterview, error) {
	result, err := ci.Edges.InterviewEdgeOrErr()
	if IsNotLoaded(err) {
		result, err = ci.QueryInterviewEdge().Only(ctx)
	}
	return result, err
}

func (cj *CandidateJob) AttachmentEdges(ctx context.Context) (result []*Attachment, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = cj.NamedAttachmentEdges(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = cj.Edges.AttachmentEdgesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = cj.QueryAttachmentEdges().All(ctx)
	}
	return result, err
}

func (cj *CandidateJob) HiringJobEdge(ctx context.Context) (*HiringJob, error) {
	result, err := cj.Edges.HiringJobEdgeOrErr()
	if IsNotLoaded(err) {
		result, err = cj.QueryHiringJobEdge().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (cj *CandidateJob) CandidateJobFeedback(ctx context.Context) (result []*CandidateJobFeedback, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = cj.NamedCandidateJobFeedback(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = cj.Edges.CandidateJobFeedbackOrErr()
	}
	if IsNotLoaded(err) {
		result, err = cj.QueryCandidateJobFeedback().All(ctx)
	}
	return result, err
}

func (cj *CandidateJob) CandidateEdge(ctx context.Context) (*Candidate, error) {
	result, err := cj.Edges.CandidateEdgeOrErr()
	if IsNotLoaded(err) {
		result, err = cj.QueryCandidateEdge().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (cj *CandidateJob) CandidateJobInterview(ctx context.Context) (result []*CandidateInterview, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = cj.NamedCandidateJobInterview(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = cj.Edges.CandidateJobInterviewOrErr()
	}
	if IsNotLoaded(err) {
		result, err = cj.QueryCandidateJobInterview().All(ctx)
	}
	return result, err
}

func (cj *CandidateJob) CreatedByEdge(ctx context.Context) (*User, error) {
	result, err := cj.Edges.CreatedByEdgeOrErr()
	if IsNotLoaded(err) {
		result, err = cj.QueryCreatedByEdge().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (cj *CandidateJob) CandidateJobStep(ctx context.Context) (result []*CandidateJobStep, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = cj.NamedCandidateJobStep(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = cj.Edges.CandidateJobStepOrErr()
	}
	if IsNotLoaded(err) {
		result, err = cj.QueryCandidateJobStep().All(ctx)
	}
	return result, err
}

func (cjf *CandidateJobFeedback) CreatedByEdge(ctx context.Context) (*User, error) {
	result, err := cjf.Edges.CreatedByEdgeOrErr()
	if IsNotLoaded(err) {
		result, err = cjf.QueryCreatedByEdge().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (cjf *CandidateJobFeedback) CandidateJobEdge(ctx context.Context) (*CandidateJob, error) {
	result, err := cjf.Edges.CandidateJobEdgeOrErr()
	if IsNotLoaded(err) {
		result, err = cjf.QueryCandidateJobEdge().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (cjf *CandidateJobFeedback) AttachmentEdges(ctx context.Context) (result []*Attachment, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = cjf.NamedAttachmentEdges(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = cjf.Edges.AttachmentEdgesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = cjf.QueryAttachmentEdges().All(ctx)
	}
	return result, err
}

func (cjs *CandidateJobStep) CandidateJobEdge(ctx context.Context) (*CandidateJob, error) {
	result, err := cjs.Edges.CandidateJobEdgeOrErr()
	if IsNotLoaded(err) {
		result, err = cjs.QueryCandidateJobEdge().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (hj *HiringJob) OwnerEdge(ctx context.Context) (*User, error) {
	result, err := hj.Edges.OwnerEdgeOrErr()
	if IsNotLoaded(err) {
		result, err = hj.QueryOwnerEdge().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (hj *HiringJob) TeamEdge(ctx context.Context) (*Team, error) {
	result, err := hj.Edges.TeamEdgeOrErr()
	if IsNotLoaded(err) {
		result, err = hj.QueryTeamEdge().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (hj *HiringJob) CandidateJobEdges(ctx context.Context) (result []*CandidateJob, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = hj.NamedCandidateJobEdges(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = hj.Edges.CandidateJobEdgesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = hj.QueryCandidateJobEdges().All(ctx)
	}
	return result, err
}

func (t *Team) UserEdges(ctx context.Context) (result []*User, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = t.NamedUserEdges(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = t.Edges.UserEdgesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = t.QueryUserEdges().All(ctx)
	}
	return result, err
}

func (t *Team) TeamJobEdges(ctx context.Context) (result []*HiringJob, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = t.NamedTeamJobEdges(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = t.Edges.TeamJobEdgesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = t.QueryTeamJobEdges().All(ctx)
	}
	return result, err
}

func (t *Team) UserTeams(ctx context.Context) (result []*TeamManager, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = t.NamedUserTeams(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = t.Edges.UserTeamsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = t.QueryUserTeams().All(ctx)
	}
	return result, err
}

func (tm *TeamManager) UserEdge(ctx context.Context) (*User, error) {
	result, err := tm.Edges.UserEdgeOrErr()
	if IsNotLoaded(err) {
		result, err = tm.QueryUserEdge().Only(ctx)
	}
	return result, err
}

func (tm *TeamManager) TeamEdge(ctx context.Context) (*Team, error) {
	result, err := tm.Edges.TeamEdgeOrErr()
	if IsNotLoaded(err) {
		result, err = tm.QueryTeamEdge().Only(ctx)
	}
	return result, err
}

func (u *User) AuditEdge(ctx context.Context) (result []*AuditTrail, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedAuditEdge(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.AuditEdgeOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryAuditEdge().All(ctx)
	}
	return result, err
}

func (u *User) HiringOwner(ctx context.Context) (result []*HiringJob, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedHiringOwner(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.HiringOwnerOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryHiringOwner().All(ctx)
	}
	return result, err
}

func (u *User) TeamEdges(ctx context.Context) (result []*Team, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedTeamEdges(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.TeamEdgesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryTeamEdges().All(ctx)
	}
	return result, err
}

func (u *User) CandidateJobFeedback(ctx context.Context) (result []*CandidateJobFeedback, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedCandidateJobFeedback(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.CandidateJobFeedbackOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryCandidateJobFeedback().All(ctx)
	}
	return result, err
}

func (u *User) InterviewEdges(ctx context.Context) (result []*CandidateInterview, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedInterviewEdges(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.InterviewEdgesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryInterviewEdges().All(ctx)
	}
	return result, err
}

func (u *User) CandidateJobEdges(ctx context.Context) (result []*CandidateJob, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedCandidateJobEdges(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.CandidateJobEdgesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryCandidateJobEdges().All(ctx)
	}
	return result, err
}

func (u *User) CandidateInterviewEdges(ctx context.Context) (result []*CandidateInterview, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedCandidateInterviewEdges(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.CandidateInterviewEdgesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryCandidateInterviewEdges().All(ctx)
	}
	return result, err
}

func (u *User) TeamUsers(ctx context.Context) (result []*TeamManager, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedTeamUsers(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.TeamUsersOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryTeamUsers().All(ctx)
	}
	return result, err
}

func (u *User) InterviewUsers(ctx context.Context) (result []*CandidateInterviewer, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedInterviewUsers(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.InterviewUsersOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryInterviewUsers().All(ctx)
	}
	return result, err
}
