package models

import (
	"fmt"
	"trec/ent"

	"github.com/samber/lo"
)

var (
	GeneralEmailTpKeyword = []string{
		"gl:receiver_name",
	}

	RecTeamEmailTpKeyword = []string{
		"rec:name",
		"rec:leader",
	}

	HiringTeamEmailTpKeyword = []string{
		"hrtm:name",
		"hrtm:manager_name",
		"hrtm:approvers",
	}

	JobPositionEmailTpKeyword = []string{
		"jbpos:name",
	}

	HiringJobEmailTpKeyword = []string{
		"hrjb:name",
		"hrjb:rec_in_charge",
		"hrjb:skill_name",
		"hrjb:level",
		"hrjb:location",
		"hrjb:requester",
		"hrjb:staff_required",
		"hrjb:status",
		"hrjb:priority",
		"hrjb:salary",
		"hrjb:description",
		"hrjb:resolution_time",
		"hrjb:audit_trail",
	}

	CandidateEmailTpKeyword = []string{
		"cd:name",
		"cd:email",
		"cd:phone",
		"cd:address",
		"cd:recruiter",
		"cd:recruit_date",
		"cd:dob",
		"cd:source",
		"cd:skill_name",
	}

	CandidateAppEmailTpKeyword = []string{
		"cdjb:status",
		"cdjb:applied_date",
		"cdjb:onboard_date",
		"cdjb:offer_expiration_date",
		"cdjb:rec_in_charge",
	}

	InterviewEmailTpKeyword = []string{
		"intv:title",
		"intv:interviewer_name",
		"intv:date",
		"intv:time",
		"intv:location",
		"intv:meeting_link",
	}

	LinkEmailTpJobRequestKeyword = []string{
		"lk:hiring_team",
		"lk:rec_team",
		"lk:job",
	}

	LinkEmailTpCandidateKeyword = append(
		LinkEmailTpJobRequestKeyword,
		"lk:candidate",
		"lk:candidate_job_application",
	)

	LinkEmailTpInterviewKeyword = []string{
		"lk:interview",
	}

	LinkEmailTpActionKeyword = []string{
		"lk:approve",
		"lk:reject",
	}

	AllEmailTPKeyword = lo.SliceToMap(
		lo.Union(
			GeneralEmailTpKeyword, RecTeamEmailTpKeyword, HiringTeamEmailTpKeyword,
			JobPositionEmailTpKeyword, HiringJobEmailTpKeyword, CandidateEmailTpKeyword,
			CandidateAppEmailTpKeyword, InterviewEmailTpKeyword, LinkEmailTpCandidateKeyword, LinkEmailTpInterviewKeyword,
			LinkEmailTpActionKeyword,
		),
		func(keyword string) (string, string) {
			wrappedKeyword := fmt.Sprintf("{{ %s }}", keyword)
			return wrappedKeyword, ""
		},
	)

	EmailTpJobRequestSubjectKeyword = lo.Union(
		GeneralEmailTpKeyword, RecTeamEmailTpKeyword, HiringTeamEmailTpKeyword,
		JobPositionEmailTpKeyword, HiringJobEmailTpKeyword,
	)

	EmailTpJobRequestContentKeyword = lo.Union(
		EmailTpJobRequestSubjectKeyword, LinkEmailTpJobRequestKeyword, LinkEmailTpActionKeyword,
	)

	EmailTpApplicationSubjectKeyword = lo.Union(
		EmailTpJobRequestSubjectKeyword, CandidateEmailTpKeyword, CandidateAppEmailTpKeyword,
	)

	EmailTpApplicationContentKeyword = append(EmailTpApplicationSubjectKeyword, LinkEmailTpCandidateKeyword...)

	EmailTpInterviewSubjectKeyword = append(EmailTpApplicationSubjectKeyword, InterviewEmailTpKeyword...)

	EmailTpInterviewContentKeyword = lo.Union(
		EmailTpInterviewSubjectKeyword, LinkEmailTpCandidateKeyword, LinkEmailTpInterviewKeyword,
	)

	GeneralEmailTpKeywordJson = []*ent.JSONFormat{
		{Key: "gl:receiver_name", Value: "Receiver name"},
	}

	RecTeamEmailTpKeywordJson = []*ent.JSONFormat{
		{Key: "rec:name", Value: "Recruiter name"},
		{Key: "rec:leader", Value: "Recruiter leader"},
	}

	HiringTeamEmailTpKeywordJson = []*ent.JSONFormat{
		{Key: "hrtm:name", Value: "Hiring Team name"},
		{Key: "hrtm:manager_name", Value: "Hiring Team Manager"},
		{Key: "hrtm:approvers", Value: "Hiring Team Approvers"},
	}

	JobPositionEmailTpKeywordJson = []*ent.JSONFormat{
		{Key: "jbpos:name", Value: "Job Position name"},
	}

	HiringJobEmailTpKeywordJson = []*ent.JSONFormat{
		{Key: "hrjb:name", Value: "Job name"},
		{Key: "hrjb:rec_in_charge", Value: "Job Recruiter In Charge"},
		{Key: "hrjb:skill_name", Value: "Job's Skill Required"},
		{Key: "hrjb:level", Value: "Job's Level"},
		{Key: "hrjb:location", Value: "Job Location"},
		{Key: "hrjb:requester", Value: "Job Requester"},
		{Key: "hrjb:staff_required", Value: "Job Staff required"},
		{Key: "hrjb:status", Value: "Job Status"},
		{Key: "hrjb:priority", Value: "Job Priority"},
		{Key: "hrjb:salary", Value: "Job Salary"},
		{Key: "hrjb:description", Value: "Job Description"},
		{Key: "hrjb:resolution_time", Value: "Job Resolution Time"},
		{Key: "hrjb:audit_trail", Value: "Job History Log"},
	}

	CandidateEmailTpKeywordJson = []*ent.JSONFormat{
		{Key: "cd:name", Value: "Candidate name"},
		{Key: "cd:email", Value: "Candidate Email"},
		{Key: "cd:phone", Value: "Candidate Phone"},
		{Key: "cd:address", Value: "Candidate Address"},
		{Key: "cd:recruiter", Value: "Candidate Recruiter"},
		{Key: "cd:recruit_date", Value: "Candidate Recruit date"},
		{Key: "cd:dob", Value: "Candidate Date of birth"},
		{Key: "cd:source", Value: "Candidate Source"},
		{Key: "cd:skill_name", Value: "Candidate Skill name"},
	}

	CandidateAppEmailTpKeywordJson = []*ent.JSONFormat{
		{Key: "cdjb:status", Value: "Application Status"},
		{Key: "cdjb:applied_date", Value: "Application Applied date"},
		{Key: "cdjb:onboard_date", Value: "Candidate planned onboard date"},
		{Key: "cdjb:offer_expiration_date", Value: "Offer expiration date"},
		{Key: "cdjb:rec_in_charge", Value: "Application Recruiter In Charge"},
	}

	InterviewEmailTpKeywordJson = []*ent.JSONFormat{
		{Key: "intv:title", Value: "Interview Title"},
		{Key: "intv:interviewer_name", Value: "Interviewer name"},
		{Key: "intv:date", Value: "Interview Date"},
		{Key: "intv:time", Value: "Interview Time"},
		{Key: "intv:location", Value: "Interview Location"},
		{Key: "intv:meeting_link", Value: "Interview Meeting Link"},
	}

	LinkEmailTpJobRequestKeywordJson = []*ent.JSONFormat{
		{Key: "lk:hiring_team", Value: "Link To Hiring Team"},
		{Key: "lk:rec_team", Value: "Link To Recruitment Team"},
		{Key: "lk:job", Value: "Link To Job"},
	}

	LinkEmailTpKeywordCandidateJson = []*ent.JSONFormat{
		{Key: "lk:candidate", Value: "Link To Candidate"},
		{Key: "lk:candidate_job_application", Value: "Link To Application"},
	}

	LinkEmailTpKeywordInterviewJson = []*ent.JSONFormat{
		{Key: "lk:interview", Value: "Link To Interview"},
	}

	LinkEmailTpKeywordActionJson = []*ent.JSONFormat{
		{Key: "lk:approve", Value: "Link To Approve"},
		{Key: "lk:reject", Value: "Link To Reject"},
	}

	EmailTpErrorString = map[string]string{
		"gl":    "model.email_template.validation.gl.keyword_invalid",
		"rec":   "model.email_template.validation.rec.keyword_invalid",
		"hrtm":  "model.email_template.validation.hrtm.keyword_invalid",
		"jbpos": "model.email_template.validation.jbpos.keyword_invalid",
		"hrjb":  "model.email_template.validation.hrjb.keyword_invalid",
		"cdjb":  "model.email_template.validation.cdjb.keyword_invalid",
		"intv":  "model.email_template.validation.intv.keyword_invalid",
		"lk":    "model.email_template.validation.lk.keyword_invalid",
		"cd":    "model.email_template.validation.cd.keyword_invalid",
	}
)
