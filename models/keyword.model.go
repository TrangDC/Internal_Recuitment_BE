package models

import "trec/ent"

var GeneralEmailTpKeyword = []string{
	"gl:receiver_name",
}

var HiringTeamEmailTpKeyword = []string{
	"hrtm:name",
	"hrtm:manager_name",
}

var HiringJobEmailTpKeyword = []string{
	"hrjb:name",
	"hrjb:skill_name",
	"hrjb:location",
	"hrjb:requester",
	"hrjb:staff_required",
	"hrjb:status",
	"hrjb:priority",
	"hrjb:salary",
	"hrjb:description",
}

var CandidateEmailTpKeyword = []string{
	"cd:name",
	"cd:email",
	"cd:phone",
	"cd:recruiter",
	"cd:recruit_date",
	"cd:dob",
	"cd:source",
	"cd:skill_name",
}

var CandidateAppEmailTpKeyword = []string{
	"cdjb:status",
	"cdjb:applied_date",
	"cdjb:onboard_date",
	"cdjb:offer_expiration_date",
}

var InterviewEmailTpKeyword = []string{
	"intv:title",
	"intv:interviewer_name",
	"intv:date",
	"intv:time",
	"intv:location",
}

var LinkEmailTpCandidateKeyword = []string{
	"lk:hiring_team",
	"lk:job",
	"lk:candidate",
	"lk:candidate_job_application",
}
var LinkEmailTpInterviewKeyword = []string{
	"lk:interview",
}

var AllEmailTPKeyword = map[string]string{
	"{{ gl:receiver_name }}":             "",
	"{{ hrtm:name }}":                    "",
	"{{ hrtm:manager_name }}":            "",
	"{{ hrjb:name }}":                    "",
	"{{ hrjb:skill_name }}":              "",
	"{{ hrjb:location }}":                "",
	"{{ hrjb:requester }}":               "",
	"{{ hrjb:staff_required }}":          "",
	"{{ hrjb:status }}":                  "",
	"{{ hrjb:priority }}":                "",
	"{{ hrjb:salary }}":                  "",
	"{{ hrjb:description }}":             "",
	"{{ cd:name }}":                      "",
	"{{ cd:email }}":                     "",
	"{{ cd:phone }}":                     "",
	"{{ cd:recruiter }}":                 "",
	"{{ cd:recruit_date }}":              "",
	"{{ cd:dob }}":                       "",
	"{{ cd:source }}":                    "",
	"{{ cd:skill_name }}":                "",
	"{{ cdjb:status }}":                  "",
	"{{ cdjb:applied_date }}":            "",
	"{{ cdjb:onboard_date }}":            "",
	"{{ cdjb:offer_expiration_date }}":   "",
	"{{ intv:title }}":                   "",
	"{{ intv:interviewer_name }}":        "",
	"{{ intv:date }}":                    "",
	"{{ intv:time }}":                    "",
	"{{ intv:location }}":                "",
	"{{ lk:hiring_team }}":               "",
	"{{ lk:job }}":                       "",
	"{{ lk:candidate }}":                 "",
	"{{ lk:candidate_job_application }}": "",
	"{{ lk:interview }}":                 "",
}

var EmailTpApplicationSubjectKeyword = append(
	append(
		append(
			append(GeneralEmailTpKeyword, HiringTeamEmailTpKeyword...),
			HiringJobEmailTpKeyword...),
		CandidateEmailTpKeyword...),
	CandidateAppEmailTpKeyword...)

var EmailTpApplicationContentKeyword = append(EmailTpApplicationSubjectKeyword, LinkEmailTpCandidateKeyword...)

var EmailTpInterviewSubjectKeyword = append(EmailTpApplicationSubjectKeyword, InterviewEmailTpKeyword...)

var EmailTpInterviewContentKeyword = append(
	append(
		EmailTpInterviewSubjectKeyword, LinkEmailTpCandidateKeyword...),
	LinkEmailTpInterviewKeyword...)

var GeneralEmailTpKeywordJson = []*ent.JSONFormat{
	{Key: "gl:receiver_name", Value: "Receiver name"},
}

var TeamEmailTpKeywordJson = []*ent.JSONFormat{
	{Key: "hrtm:name", Value: "Hiring Team name"},
	{Key: "hrtm:manager_name", Value: "Hiring Team Manager"},
}

var HiringJobEmailTpKeywordJson = []*ent.JSONFormat{
	{Key: "hrjb:name", Value: "Job name"},
	{Key: "hrjb:skill_name", Value: "Job's Skill Required"},
	{Key: "hrjb:location", Value: "Job Location"},
	{Key: "hrjb:requester", Value: "Job Requester"},
	{Key: "hrjb:staff_required", Value: "Job Staff required"},
	{Key: "hrjb:status", Value: "Job Status"},
	{Key: "hrjb:priority", Value: "Job Priority"},
	{Key: "hrjb:salary", Value: "Job Salary"},
	{Key: "hrjb:description", Value: "Job Description"},
}

var CandidateEmailTpKeywordJson = []*ent.JSONFormat{
	{Key: "cd:name", Value: "Candidate name"},
	{Key: "cd:email", Value: "Candidate Email"},
	{Key: "cd:phone", Value: "Candidate Phone"},
	{Key: "cd:recruiter", Value: "Candidate Recruiter"},
	{Key: "cd:recruit_date", Value: "Candidate Recruit date"},
	{Key: "cd:dob", Value: "Candidate Date of birth"},
	{Key: "cd:source", Value: "Candidate Source"},
	{Key: "cd:skill_name", Value: "Candidate Skill name"},
}

var CandidateAppEmailTpKeywordJson = []*ent.JSONFormat{
	{Key: "cdjb:status", Value: "Application Status"},
	{Key: "cdjb:applied_date", Value: "Application Applied date"},
	{Key: "cdjb:onboard_date", Value: "Candidate planned onboard date"},
	{Key: "cdjb:offer_expiration_date", Value: "Offer expiration date"},
}

var InterviewEmailTpKeywordJson = []*ent.JSONFormat{
	{Key: "intv:title", Value: "Interview Title"},
	{Key: "intv:interviewer_name", Value: "Interviewer name"},
	{Key: "intv:date", Value: "Interview Date"},
	{Key: "intv:time", Value: "Interview Time"},
	{Key: "intv:location", Value: "Interview Location"},
}

var LinkEmailTpKeywordCandidateJson = []*ent.JSONFormat{
	{Key: "lk:hiring_team", Value: "Link To Hiring Team"},
	{Key: "lk:job", Value: "Link To Job"},
	{Key: "lk:candidate", Value: "Link To Candidate"},
	{Key: "lk:candidate_job_application", Value: "Link To Application"},
}
var LinkEmailTpKeywordInterviewJson = []*ent.JSONFormat{
	{Key: "lk:interview", Value: "Link To Interview"},
}

var EmailTpErrorString = map[string]string{
	"gl":   "model.email_template.validation.gl.keyword_invalid",
	"hrtm": "model.email_template.validation.hrtm.keyword_invalid",
	"hrjb": "model.email_template.validation.hrjb.keyword_invalid",
	"cdjb": "model.email_template.validation.cdjb.keyword_invalid",
	"intv": "model.email_template.validation.intv.keyword_invalid",
	"lk":   "model.email_template.validation.lk.keyword_invalid",
	"cd":   "model.email_template.validation.cd.keyword_invalid",
}
