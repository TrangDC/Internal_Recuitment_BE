// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (a *Attachment) CandidateJobEdge(ctx context.Context) (*CandidateJob, error) {
	result, err := a.Edges.CandidateJobEdgeOrErr()
	if IsNotLoaded(err) {
		result, err = a.QueryCandidateJobEdge().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (a *Attachment) CandidateJobFeedbackEdge(ctx context.Context) (*CandidateJobFeedback, error) {
	result, err := a.Edges.CandidateJobFeedbackEdgeOrErr()
	if IsNotLoaded(err) {
		result, err = a.QueryCandidateJobFeedbackEdge().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (a *Attachment) CandidateInterviewEdge(ctx context.Context) (*CandidateInterview, error) {
	result, err := a.Edges.CandidateInterviewEdgeOrErr()
	if IsNotLoaded(err) {
		result, err = a.QueryCandidateInterviewEdge().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (a *Attachment) CandidateEdge(ctx context.Context) (*Candidate, error) {
	result, err := a.Edges.CandidateEdgeOrErr()
	if IsNotLoaded(err) {
		result, err = a.QueryCandidateEdge().Only(ctx)
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

func (c *Candidate) ReferenceUserEdge(ctx context.Context) (*User, error) {
	result, err := c.Edges.ReferenceUserEdgeOrErr()
	if IsNotLoaded(err) {
		result, err = c.QueryReferenceUserEdge().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (c *Candidate) AttachmentEdges(ctx context.Context) (result []*Attachment, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedAttachmentEdges(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.AttachmentEdgesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryAttachmentEdges().All(ctx)
	}
	return result, err
}

func (c *Candidate) CandidateSkillEdges(ctx context.Context) (result []*EntitySkill, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = c.NamedCandidateSkillEdges(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = c.Edges.CandidateSkillEdgesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = c.QueryCandidateSkillEdges().All(ctx)
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

func (era *EmailRoleAttribute) EmailTemplateEdge(ctx context.Context) (*EmailTemplate, error) {
	result, err := era.Edges.EmailTemplateEdgeOrErr()
	if IsNotLoaded(err) {
		result, err = era.QueryEmailTemplateEdge().Only(ctx)
	}
	return result, err
}

func (era *EmailRoleAttribute) RoleEdge(ctx context.Context) (*Role, error) {
	result, err := era.Edges.RoleEdgeOrErr()
	if IsNotLoaded(err) {
		result, err = era.QueryRoleEdge().Only(ctx)
	}
	return result, err
}

func (et *EmailTemplate) RoleEdges(ctx context.Context) (result []*Role, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = et.NamedRoleEdges(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = et.Edges.RoleEdgesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = et.QueryRoleEdges().All(ctx)
	}
	return result, err
}

func (et *EmailTemplate) RoleEmailTemplates(ctx context.Context) (result []*EmailRoleAttribute, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = et.NamedRoleEmailTemplates(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = et.Edges.RoleEmailTemplatesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = et.QueryRoleEmailTemplates().All(ctx)
	}
	return result, err
}

func (ep *EntityPermission) PermissionEdges(ctx context.Context) (*Permission, error) {
	result, err := ep.Edges.PermissionEdgesOrErr()
	if IsNotLoaded(err) {
		result, err = ep.QueryPermissionEdges().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (ep *EntityPermission) UserEdge(ctx context.Context) (*User, error) {
	result, err := ep.Edges.UserEdgeOrErr()
	if IsNotLoaded(err) {
		result, err = ep.QueryUserEdge().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (ep *EntityPermission) RoleEdge(ctx context.Context) (*Role, error) {
	result, err := ep.Edges.RoleEdgeOrErr()
	if IsNotLoaded(err) {
		result, err = ep.QueryRoleEdge().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (es *EntitySkill) SkillEdge(ctx context.Context) (*Skill, error) {
	result, err := es.Edges.SkillEdgeOrErr()
	if IsNotLoaded(err) {
		result, err = es.QuerySkillEdge().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (es *EntitySkill) HiringJobEdge(ctx context.Context) (*HiringJob, error) {
	result, err := es.Edges.HiringJobEdgeOrErr()
	if IsNotLoaded(err) {
		result, err = es.QueryHiringJobEdge().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (es *EntitySkill) CandidateEdge(ctx context.Context) (*Candidate, error) {
	result, err := es.Edges.CandidateEdgeOrErr()
	if IsNotLoaded(err) {
		result, err = es.QueryCandidateEdge().Only(ctx)
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

func (hj *HiringJob) HiringJobSkillEdges(ctx context.Context) (result []*EntitySkill, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = hj.NamedHiringJobSkillEdges(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = hj.Edges.HiringJobSkillEdgesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = hj.QueryHiringJobSkillEdges().All(ctx)
	}
	return result, err
}

func (hj *HiringJob) HiringTeamEdge(ctx context.Context) (*HiringTeam, error) {
	result, err := hj.Edges.HiringTeamEdgeOrErr()
	if IsNotLoaded(err) {
		result, err = hj.QueryHiringTeamEdge().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (ht *HiringTeam) UserEdges(ctx context.Context) (result []*User, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = ht.NamedUserEdges(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = ht.Edges.UserEdgesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = ht.QueryUserEdges().All(ctx)
	}
	return result, err
}

func (ht *HiringTeam) HiringTeamJobEdges(ctx context.Context) (result []*HiringJob, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = ht.NamedHiringTeamJobEdges(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = ht.Edges.HiringTeamJobEdgesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = ht.QueryHiringTeamJobEdges().All(ctx)
	}
	return result, err
}

func (ht *HiringTeam) HiringMemberEdges(ctx context.Context) (result []*User, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = ht.NamedHiringMemberEdges(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = ht.Edges.HiringMemberEdgesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = ht.QueryHiringMemberEdges().All(ctx)
	}
	return result, err
}

func (ht *HiringTeam) UserHiringTeams(ctx context.Context) (result []*HiringTeamManager, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = ht.NamedUserHiringTeams(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = ht.Edges.UserHiringTeamsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = ht.QueryUserHiringTeams().All(ctx)
	}
	return result, err
}

func (htm *HiringTeamManager) UserEdge(ctx context.Context) (*User, error) {
	result, err := htm.Edges.UserEdgeOrErr()
	if IsNotLoaded(err) {
		result, err = htm.QueryUserEdge().Only(ctx)
	}
	return result, err
}

func (htm *HiringTeamManager) HiringTeamEdge(ctx context.Context) (*HiringTeam, error) {
	result, err := htm.Edges.HiringTeamEdgeOrErr()
	if IsNotLoaded(err) {
		result, err = htm.QueryHiringTeamEdge().Only(ctx)
	}
	return result, err
}

func (pe *Permission) GroupPermissionEdge(ctx context.Context) (*PermissionGroup, error) {
	result, err := pe.Edges.GroupPermissionEdgeOrErr()
	if IsNotLoaded(err) {
		result, err = pe.QueryGroupPermissionEdge().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (pe *Permission) UserPermissionEdge(ctx context.Context) (result []*EntityPermission, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = pe.NamedUserPermissionEdge(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = pe.Edges.UserPermissionEdgeOrErr()
	}
	if IsNotLoaded(err) {
		result, err = pe.QueryUserPermissionEdge().All(ctx)
	}
	return result, err
}

func (pg *PermissionGroup) GroupPermissionParent(ctx context.Context) (*PermissionGroup, error) {
	result, err := pg.Edges.GroupPermissionParentOrErr()
	if IsNotLoaded(err) {
		result, err = pg.QueryGroupPermissionParent().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (pg *PermissionGroup) GroupPermissionChildren(ctx context.Context) (result []*PermissionGroup, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = pg.NamedGroupPermissionChildren(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = pg.Edges.GroupPermissionChildrenOrErr()
	}
	if IsNotLoaded(err) {
		result, err = pg.QueryGroupPermissionChildren().All(ctx)
	}
	return result, err
}

func (pg *PermissionGroup) PermissionEdges(ctx context.Context) (result []*Permission, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = pg.NamedPermissionEdges(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = pg.Edges.PermissionEdgesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = pg.QueryPermissionEdges().All(ctx)
	}
	return result, err
}

func (rt *RecTeam) RecLeaderEdge(ctx context.Context) (*User, error) {
	result, err := rt.Edges.RecLeaderEdgeOrErr()
	if IsNotLoaded(err) {
		result, err = rt.QueryRecLeaderEdge().Only(ctx)
	}
	return result, err
}

func (rt *RecTeam) RecMemberEdges(ctx context.Context) (result []*User, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = rt.NamedRecMemberEdges(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = rt.Edges.RecMemberEdgesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = rt.QueryRecMemberEdges().All(ctx)
	}
	return result, err
}

func (r *Role) RolePermissionEdges(ctx context.Context) (result []*EntityPermission, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = r.NamedRolePermissionEdges(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = r.Edges.RolePermissionEdgesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = r.QueryRolePermissionEdges().All(ctx)
	}
	return result, err
}

func (r *Role) UserEdges(ctx context.Context) (result []*User, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = r.NamedUserEdges(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = r.Edges.UserEdgesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = r.QueryUserEdges().All(ctx)
	}
	return result, err
}

func (r *Role) EmailTemplateEdges(ctx context.Context) (result []*EmailTemplate, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = r.NamedEmailTemplateEdges(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = r.Edges.EmailTemplateEdgesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = r.QueryEmailTemplateEdges().All(ctx)
	}
	return result, err
}

func (r *Role) UserRoles(ctx context.Context) (result []*UserRole, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = r.NamedUserRoles(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = r.Edges.UserRolesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = r.QueryUserRoles().All(ctx)
	}
	return result, err
}

func (r *Role) EmailTemplateRoles(ctx context.Context) (result []*EmailRoleAttribute, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = r.NamedEmailTemplateRoles(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = r.Edges.EmailTemplateRolesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = r.QueryEmailTemplateRoles().All(ctx)
	}
	return result, err
}

func (s *Skill) SkillTypeEdge(ctx context.Context) (*SkillType, error) {
	result, err := s.Edges.SkillTypeEdgeOrErr()
	if IsNotLoaded(err) {
		result, err = s.QuerySkillTypeEdge().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (s *Skill) EntitySkillEdges(ctx context.Context) (result []*EntitySkill, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = s.NamedEntitySkillEdges(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = s.Edges.EntitySkillEdgesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = s.QueryEntitySkillEdges().All(ctx)
	}
	return result, err
}

func (st *SkillType) SkillEdges(ctx context.Context) (result []*Skill, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = st.NamedSkillEdges(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = st.Edges.SkillEdgesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = st.QuerySkillEdges().All(ctx)
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

func (u *User) CandidateReferenceEdges(ctx context.Context) (result []*Candidate, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedCandidateReferenceEdges(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.CandidateReferenceEdgesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryCandidateReferenceEdges().All(ctx)
	}
	return result, err
}

func (u *User) UserPermissionEdges(ctx context.Context) (result []*EntityPermission, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedUserPermissionEdges(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.UserPermissionEdgesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryUserPermissionEdges().All(ctx)
	}
	return result, err
}

func (u *User) RoleEdges(ctx context.Context) (result []*Role, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedRoleEdges(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.RoleEdgesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryRoleEdges().All(ctx)
	}
	return result, err
}

func (u *User) HiringTeamEdges(ctx context.Context) (result []*HiringTeam, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedHiringTeamEdges(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.HiringTeamEdgesOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryHiringTeamEdges().All(ctx)
	}
	return result, err
}

func (u *User) LedRecTeams(ctx context.Context) (result []*RecTeam, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedLedRecTeams(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.LedRecTeamsOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryLedRecTeams().All(ctx)
	}
	return result, err
}

func (u *User) RecTeams(ctx context.Context) (*RecTeam, error) {
	result, err := u.Edges.RecTeamsOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryRecTeams().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (u *User) MemberOfHiringTeamEdges(ctx context.Context) (*HiringTeam, error) {
	result, err := u.Edges.MemberOfHiringTeamEdgesOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryMemberOfHiringTeamEdges().Only(ctx)
	}
	return result, MaskNotFound(err)
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

func (u *User) RoleUsers(ctx context.Context) (result []*UserRole, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedRoleUsers(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.RoleUsersOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryRoleUsers().All(ctx)
	}
	return result, err
}

func (u *User) HiringTeamUsers(ctx context.Context) (result []*HiringTeamManager, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedHiringTeamUsers(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.HiringTeamUsersOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryHiringTeamUsers().All(ctx)
	}
	return result, err
}

func (ur *UserRole) UserEdge(ctx context.Context) (*User, error) {
	result, err := ur.Edges.UserEdgeOrErr()
	if IsNotLoaded(err) {
		result, err = ur.QueryUserEdge().Only(ctx)
	}
	return result, err
}

func (ur *UserRole) RoleEdge(ctx context.Context) (*Role, error) {
	result, err := ur.Edges.RoleEdgeOrErr()
	if IsNotLoaded(err) {
		result, err = ur.QueryRoleEdge().Only(ctx)
	}
	return result, err
}
