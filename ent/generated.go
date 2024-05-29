// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package ent

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type AttachmentInput struct {
	ID       string           `json:"id"`
	Folder   AttachmentFolder `json:"folder"`
	FileName string           `json:"fileName"`
	Action   AttachmentAction `json:"action"`
}

type AttachmentResponse struct {
	FileName string           `json:"fileName"`
	URL      string           `json:"url"`
	Action   AttachmentAction `json:"action"`
	ID       string           `json:"id"`
}

type AuditTrailFilter struct {
	RecordID   *string           `json:"recordId"`
	Module     *ProjectModule    `json:"module"`
	ActionType *AuditTrailAction `json:"actionType"`
	FromDate   *time.Time        `json:"fromDate"`
	ToDate     *time.Time        `json:"toDate"`
}

type AuditTrailFreeWord struct {
	RecordChange *string `json:"recordChange"`
}

type AuditTrailResponse struct {
	Data *AuditTrail `json:"data"`
}

type AuditTrailResponseGetAll struct {
	Edges      []*AuditTrailEdge `json:"edges"`
	Pagination *Pagination       `json:"pagination"`
}

type AuthenticationToken struct {
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refreshToken"`
	TokenType    string    `json:"tokenType"`
	ExpiresAt    time.Time `json:"expiresAt"`
	Email        string    `json:"email"`
}

type Base64Response struct {
	Data string `json:"data"`
}

type CandidateFilter struct {
	Name              *string                    `json:"name"`
	Email             *string                    `json:"email"`
	Phone             *string                    `json:"phone"`
	DobFromDate       *time.Time                 `json:"dob_from_date"`
	DobToDate         *time.Time                 `json:"dob_to_date"`
	Status            *CandidateStatusEnum       `json:"status"`
	FromDate          *time.Time                 `json:"from_date"`
	ToDate            *time.Time                 `json:"to_date"`
	IsBlackList       *bool                      `json:"is_black_list"`
	JobID             *string                    `json:"job_id"`
	IsAbleToInterview *bool                      `json:"is_able_to_interview"`
	FailedReason      []CandidateJobFailedReason `json:"failed_reason"`
}

type CandidateFreeWord struct {
	Name  *string `json:"name"`
	Email *string `json:"email"`
	Phone *string `json:"phone"`
}

type CandidateInterviewCalendarFilter struct {
	InterviewDate     *time.Time `json:"interview_date"`
	StartFrom         *time.Time `json:"start_from"`
	EndAt             *time.Time `json:"end_at"`
	Interviewer       []string   `json:"interviewer"`
	FromDate          *time.Time `json:"from_date"`
	ToDate            *time.Time `json:"to_date"`
	TeamID            *string    `json:"team_id"`
	HiringJobID       *string    `json:"hiring_job_id"`
	InterviewDateFrom *time.Time `json:"interview_date_from"`
	InterviewDateTo   *time.Time `json:"interview_date_to"`
}

type CandidateInterviewFilter struct {
	CandidateJobID string     `json:"candidate_job_id"`
	InterviewDate  *time.Time `json:"interview_date"`
	StartFrom      *time.Time `json:"start_from"`
	EndAt          *time.Time `json:"end_at"`
	Interviewer    []string   `json:"interviewer"`
	FromDate       *time.Time `json:"from_date"`
	ToDate         *time.Time `json:"to_date"`
}

type CandidateInterviewFreeWord struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

type CandidateInterviewResponse struct {
	Data *CandidateInterview `json:"data"`
}

type CandidateInterviewResponseGetAll struct {
	Edges      []*CandidateInterviewEdge `json:"edges"`
	Pagination *Pagination               `json:"pagination"`
}

type CandidateJobFeedbackFilter struct {
	CandidateJobID string  `json:"candidate_job_id"`
	CreatedBy      *string `json:"created_by"`
}

type CandidateJobFeedbackFreeWord struct {
	Feedback *string `json:"feedback"`
	UserName *string `json:"user_name"`
}

type CandidateJobFeedbackResponse struct {
	Data *CandidateJobFeedback `json:"data"`
}

type CandidateJobFeedbackResponseGetAll struct {
	Edges      []*CandidateJobFeedbackEdge `json:"edges"`
	Pagination *Pagination                 `json:"pagination"`
}

type CandidateJobFilter struct {
	Status       *CandidateJobStatus        `json:"status"`
	FromDate     *time.Time                 `json:"from_date"`
	ToDate       *time.Time                 `json:"to_date"`
	TeamID       *string                    `json:"team_id"`
	HiringJobID  *string                    `json:"hiring_job_id"`
	CandidateID  string                     `json:"candidate_id"`
	FailedReason []CandidateJobFailedReason `json:"failed_reason"`
}

type CandidateJobFreeWord struct {
	Team      *string `json:"team"`
	HiringJob *string `json:"hiring_job"`
}

type CandidateJobGroupByInterview struct {
	Hired        *CandidateJobGroupInterviewFeedback `json:"hired"`
	Kiv          *CandidateJobGroupInterviewFeedback `json:"kiv"`
	OfferLost    *CandidateJobGroupInterviewFeedback `json:"offer_lost"`
	ExStaff      *CandidateJobGroupInterviewFeedback `json:"ex_staff"`
	Applied      *CandidateJobGroupInterviewFeedback `json:"applied"`
	Interviewing *CandidateJobGroupInterviewFeedback `json:"interviewing"`
	Offering     *CandidateJobGroupInterviewFeedback `json:"offering"`
}

type CandidateJobGroupByInterviewResponse struct {
	Data *CandidateJobGroupByInterview `json:"data"`
}

type CandidateJobGroupByStatus struct {
	Hired        []*CandidateJob `json:"hired"`
	Kiv          []*CandidateJob `json:"kiv"`
	OfferLost    []*CandidateJob `json:"offer_lost"`
	ExStaff      []*CandidateJob `json:"ex_staff"`
	Applied      []*CandidateJob `json:"applied"`
	Interviewing []*CandidateJob `json:"interviewing"`
	Offering     []*CandidateJob `json:"offering"`
}

type CandidateJobGroupByStatusFilter struct {
	HiringJobID string  `json:"hiring_job_id"`
	CreatedBy   *string `json:"created_by"`
}

type CandidateJobGroupByStatusResponse struct {
	Data *CandidateJobGroupByStatus `json:"data"`
}

type CandidateJobGroupInterviewFeedback struct {
	Interview []*CandidateInterview   `json:"interview"`
	Feedback  []*CandidateJobFeedback `json:"feedback"`
}

type CandidateJobResponse struct {
	Data *CandidateJob `json:"data"`
}

type CandidateJobResponseGetAll struct {
	Edges      []*CandidateJobEdge `json:"edges"`
	Pagination *Pagination         `json:"pagination"`
}

type CandidateResponse struct {
	Data *Candidate `json:"data"`
}

type CandidateResponseGetAll struct {
	Edges      []*CandidateEdge `json:"edges"`
	Pagination *Pagination      `json:"pagination"`
}

type HiringJobFilter struct {
	Name    *string          `json:"name"`
	TeamIds []string         `json:"team_ids"`
	Status  *HiringJobStatus `json:"status"`
}

type HiringJobFreeWord struct {
	Name *string `json:"name"`
}

type HiringJobOrderBy struct {
	Direction OrderDirection        `json:"direction"`
	Field     HiringJobOrderByField `json:"field"`
}

type HiringJobResponse struct {
	Data *HiringJob `json:"data"`
}

type HiringJobResponseGetAll struct {
	Edges      []*HiringJobEdge `json:"edges"`
	Pagination *Pagination      `json:"pagination"`
}

type NewAttachmentInput struct {
	DocumentName string `json:"document_name"`
	DocumentID   string `json:"document_id"`
}

type NewCandidateInput struct {
	Name  string     `json:"name"`
	Email string     `json:"email"`
	Phone string     `json:"phone"`
	Dob   *time.Time `json:"dob"`
}

type NewCandidateInterview4CalendarInput struct {
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	CandidateID   []string  `json:"candidate_id"`
	JobID         string    `json:"job_id"`
	InterviewDate time.Time `json:"interview_date"`
	StartFrom     time.Time `json:"start_from"`
	EndAt         time.Time `json:"end_at"`
	Interviewer   []string  `json:"interviewer"`
}

type NewCandidateInterviewInput struct {
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	CandidateJobID string    `json:"candidate_job_id"`
	InterviewDate  time.Time `json:"interview_date"`
	StartFrom      time.Time `json:"start_from"`
	EndAt          time.Time `json:"end_at"`
	Interviewer    []string  `json:"interviewer"`
}

type NewCandidateJobFeedbackInput struct {
	CandidateJobID string                `json:"candidate_job_id"`
	Feedback       string                `json:"feedback"`
	Attachments    []*NewAttachmentInput `json:"attachments"`
}

type NewCandidateJobInput struct {
	CandidateID string                `json:"candidate_id"`
	HiringJobID string                `json:"hiring_job_id"`
	Status      CandidateJobStatus    `json:"status"`
	Attachments []*NewAttachmentInput `json:"attachments"`
}

type NewHiringJobInput struct {
	Status      HiringJobStatus `json:"status"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Amount      int             `json:"amount"`
	Location    LocationEnum    `json:"location"`
	SalaryType  SalaryTypeEnum  `json:"salary_type"`
	SalaryFrom  int             `json:"salary_from"`
	SalaryTo    int             `json:"salary_to"`
	Currency    CurrencyEnum    `json:"currency"`
	TeamID      string          `json:"team_id"`
	CreatedBy   string          `json:"created_by"`
	Priority    int             `json:"priority"`
}

type NewTeamInput struct {
	Name    string   `json:"name"`
	Members []string `json:"members"`
}

type NewUserInput struct {
	Name      string `json:"name"`
	WorkEmail string `json:"work_email"`
}

type Pagination struct {
	Page    int `json:"page"`
	PerPage int `json:"perPage"`
	Total   int `json:"total"`
}

type PaginationInput struct {
	Page    *int `json:"page"`
	PerPage *int `json:"perPage"`
}

type TeamFilter struct {
	Name *string `json:"name"`
}

type TeamFreeWord struct {
	Name *string `json:"name"`
}

type TeamOrderBy struct {
	Direction OrderDirection   `json:"direction"`
	Field     TeamOrderByField `json:"field"`
}

type TeamResponse struct {
	Data *Team `json:"data"`
}

type TeamResponseGetAll struct {
	Edges      []*TeamEdge `json:"edges"`
	Pagination *Pagination `json:"pagination"`
}

type UpdateCandidateAttachment struct {
	Attachments []*NewAttachmentInput `json:"attachments"`
}

type UpdateCandidateInput struct {
	Name  string     `json:"name"`
	Email string     `json:"email"`
	Phone string     `json:"phone"`
	Dob   *time.Time `json:"dob"`
}

type UpdateCandidateInterviewInput struct {
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	CandidateJobID string    `json:"candidate_job_id"`
	InterviewDate  time.Time `json:"interview_date"`
	StartFrom      time.Time `json:"start_from"`
	EndAt          time.Time `json:"end_at"`
	Interviewer    []string  `json:"interviewer"`
}

type UpdateCandidateInterviewScheduleInput struct {
	InterviewDate time.Time `json:"interview_date"`
	StartFrom     time.Time `json:"start_from"`
	EndAt         time.Time `json:"end_at"`
	Interviewer   []string  `json:"interviewer"`
}

type UpdateCandidateJobFeedbackInput struct {
	Feedback    string                `json:"feedback"`
	Attachments []*NewAttachmentInput `json:"attachments"`
}

type UpdateCandidateJobStatus struct {
	Status       CandidateJobStatus         `json:"status"`
	FailedReason []CandidateJobFailedReason `json:"failed_reason"`
}

type UpdateHiringJobInput struct {
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Amount      int            `json:"amount"`
	Location    LocationEnum   `json:"location"`
	SalaryType  SalaryTypeEnum `json:"salary_type"`
	SalaryFrom  int            `json:"salary_from"`
	SalaryTo    int            `json:"salary_to"`
	Currency    CurrencyEnum   `json:"currency"`
	TeamID      string         `json:"team_id"`
	CreatedBy   string         `json:"created_by"`
	Priority    int            `json:"priority"`
}

type UpdateTeamInput struct {
	Name    string   `json:"name"`
	Members []string `json:"members"`
}

type UpdateUserInput struct {
	Name      string     `json:"name"`
	WorkEmail string     `json:"work_email"`
	Status    UserStatus `json:"status"`
}

type UpdateUserStatusInput struct {
	Status UserStatus `json:"status"`
}

type UserFilter struct {
	Name      *string     `json:"name"`
	Ids       []string    `json:"ids"`
	IgnoreIds []string    `json:"ignore_ids"`
	NotInTeam *bool       `json:"not_in_team"`
	Status    *UserStatus `json:"status"`
}

type UserFreeWord struct {
	Name *string `json:"name"`
}

type UserResponse struct {
	Data *User `json:"data"`
}

type UserResponseGetAll struct {
	Edges      []*UserEdge `json:"edges"`
	Pagination *Pagination `json:"pagination"`
}

type UserSelection struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	WorkEmail string `json:"work_email"`
}

type UserSelectionEdge struct {
	Node   *UserSelection `json:"node"`
	Cursor Cursor         `json:"cursor"`
}

type UserSelectionResponseGetAll struct {
	Edges      []*UserSelectionEdge `json:"edges"`
	Pagination *Pagination          `json:"pagination"`
}

type AttachmentAction string

const (
	AttachmentActionUpload   AttachmentAction = "UPLOAD"
	AttachmentActionDownload AttachmentAction = "DOWNLOAD"
)

var AllAttachmentAction = []AttachmentAction{
	AttachmentActionUpload,
	AttachmentActionDownload,
}

func (e AttachmentAction) IsValid() bool {
	switch e {
	case AttachmentActionUpload, AttachmentActionDownload:
		return true
	}
	return false
}

func (e AttachmentAction) String() string {
	return string(e)
}

func (e *AttachmentAction) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AttachmentAction(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AttachmentAction", str)
	}
	return nil
}

func (e AttachmentAction) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type AttachmentFolder string

const (
	AttachmentFolderCandidate         AttachmentFolder = "candidate"
	AttachmentFolderCandidateFeedback AttachmentFolder = "candidate_feedback"
)

var AllAttachmentFolder = []AttachmentFolder{
	AttachmentFolderCandidate,
	AttachmentFolderCandidateFeedback,
}

func (e AttachmentFolder) IsValid() bool {
	switch e {
	case AttachmentFolderCandidate, AttachmentFolderCandidateFeedback:
		return true
	}
	return false
}

func (e AttachmentFolder) String() string {
	return string(e)
}

func (e *AttachmentFolder) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AttachmentFolder(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AttachmentFolder", str)
	}
	return nil
}

func (e AttachmentFolder) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type AttachmentRelationType string

const (
	AttachmentRelationTypeCandidateJobs         AttachmentRelationType = "candidate_jobs"
	AttachmentRelationTypeCandidateJobFeedbacks AttachmentRelationType = "candidate_job_feedbacks"
)

var AllAttachmentRelationType = []AttachmentRelationType{
	AttachmentRelationTypeCandidateJobs,
	AttachmentRelationTypeCandidateJobFeedbacks,
}

func (e AttachmentRelationType) IsValid() bool {
	switch e {
	case AttachmentRelationTypeCandidateJobs, AttachmentRelationTypeCandidateJobFeedbacks:
		return true
	}
	return false
}

func (e AttachmentRelationType) String() string {
	return string(e)
}

func (e *AttachmentRelationType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AttachmentRelationType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AttachmentRelationType", str)
	}
	return nil
}

func (e AttachmentRelationType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type CandidateInterviewStatusEditable string

const (
	CandidateInterviewStatusEditableApplied      CandidateInterviewStatusEditable = "applied"
	CandidateInterviewStatusEditableInterviewing CandidateInterviewStatusEditable = "interviewing"
)

var AllCandidateInterviewStatusEditable = []CandidateInterviewStatusEditable{
	CandidateInterviewStatusEditableApplied,
	CandidateInterviewStatusEditableInterviewing,
}

func (e CandidateInterviewStatusEditable) IsValid() bool {
	switch e {
	case CandidateInterviewStatusEditableApplied, CandidateInterviewStatusEditableInterviewing:
		return true
	}
	return false
}

func (e CandidateInterviewStatusEditable) String() string {
	return string(e)
}

func (e *CandidateInterviewStatusEditable) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CandidateInterviewStatusEditable(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CandidateInterviewStatusEditable", str)
	}
	return nil
}

func (e CandidateInterviewStatusEditable) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type CandidateJobFailedReason string

const (
	CandidateJobFailedReasonPoorProfessionalism      CandidateJobFailedReason = "poor_professionalism"
	CandidateJobFailedReasonPoorFitAndEngagement     CandidateJobFailedReason = "poor_fit_and_engagement"
	CandidateJobFailedReasonOverExpectations         CandidateJobFailedReason = "over_expectations"
	CandidateJobFailedReasonOverQualification        CandidateJobFailedReason = "over_qualification"
	CandidateJobFailedReasonLanguageDeficiency       CandidateJobFailedReason = "language_deficiency"
	CandidateJobFailedReasonWeakTechnicalSkills      CandidateJobFailedReason = "weak_technical_skills"
	CandidateJobFailedReasonPoorInterpersonalSkills  CandidateJobFailedReason = "poor_interpersonal_skills"
	CandidateJobFailedReasonPoorProblemSolvingSkills CandidateJobFailedReason = "poor_problem_solving_skills"
	CandidateJobFailedReasonPoorManagementSkills     CandidateJobFailedReason = "poor_management_skills"
	CandidateJobFailedReasonCandidateWithdrawal      CandidateJobFailedReason = "candidate_withdrawal"
	CandidateJobFailedReasonOthers                   CandidateJobFailedReason = "others"
)

var AllCandidateJobFailedReason = []CandidateJobFailedReason{
	CandidateJobFailedReasonPoorProfessionalism,
	CandidateJobFailedReasonPoorFitAndEngagement,
	CandidateJobFailedReasonOverExpectations,
	CandidateJobFailedReasonOverQualification,
	CandidateJobFailedReasonLanguageDeficiency,
	CandidateJobFailedReasonWeakTechnicalSkills,
	CandidateJobFailedReasonPoorInterpersonalSkills,
	CandidateJobFailedReasonPoorProblemSolvingSkills,
	CandidateJobFailedReasonPoorManagementSkills,
	CandidateJobFailedReasonCandidateWithdrawal,
	CandidateJobFailedReasonOthers,
}

func (e CandidateJobFailedReason) IsValid() bool {
	switch e {
	case CandidateJobFailedReasonPoorProfessionalism, CandidateJobFailedReasonPoorFitAndEngagement, CandidateJobFailedReasonOverExpectations, CandidateJobFailedReasonOverQualification, CandidateJobFailedReasonLanguageDeficiency, CandidateJobFailedReasonWeakTechnicalSkills, CandidateJobFailedReasonPoorInterpersonalSkills, CandidateJobFailedReasonPoorProblemSolvingSkills, CandidateJobFailedReasonPoorManagementSkills, CandidateJobFailedReasonCandidateWithdrawal, CandidateJobFailedReasonOthers:
		return true
	}
	return false
}

func (e CandidateJobFailedReason) String() string {
	return string(e)
}

func (e *CandidateJobFailedReason) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CandidateJobFailedReason(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CandidateJobFailedReason", str)
	}
	return nil
}

func (e CandidateJobFailedReason) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type CandidateJobStatus string

const (
	CandidateJobStatusHired        CandidateJobStatus = "hired"
	CandidateJobStatusKiv          CandidateJobStatus = "kiv"
	CandidateJobStatusOfferLost    CandidateJobStatus = "offer_lost"
	CandidateJobStatusExStaff      CandidateJobStatus = "ex_staff"
	CandidateJobStatusApplied      CandidateJobStatus = "applied"
	CandidateJobStatusInterviewing CandidateJobStatus = "interviewing"
	CandidateJobStatusOffering     CandidateJobStatus = "offering"
	CandidateJobStatusNew          CandidateJobStatus = "new"
)

var AllCandidateJobStatus = []CandidateJobStatus{
	CandidateJobStatusHired,
	CandidateJobStatusKiv,
	CandidateJobStatusOfferLost,
	CandidateJobStatusExStaff,
	CandidateJobStatusApplied,
	CandidateJobStatusInterviewing,
	CandidateJobStatusOffering,
	CandidateJobStatusNew,
}

func (e CandidateJobStatus) IsValid() bool {
	switch e {
	case CandidateJobStatusHired, CandidateJobStatusKiv, CandidateJobStatusOfferLost, CandidateJobStatusExStaff, CandidateJobStatusApplied, CandidateJobStatusInterviewing, CandidateJobStatusOffering, CandidateJobStatusNew:
		return true
	}
	return false
}

func (e CandidateJobStatus) String() string {
	return string(e)
}

func (e *CandidateJobStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CandidateJobStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CandidateJobStatus", str)
	}
	return nil
}

func (e CandidateJobStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type CandidateJobStatusEnded string

const (
	CandidateJobStatusEndedHired     CandidateJobStatusEnded = "hired"
	CandidateJobStatusEndedKiv       CandidateJobStatusEnded = "kiv"
	CandidateJobStatusEndedOfferLost CandidateJobStatusEnded = "offer_lost"
	CandidateJobStatusEndedExStaff   CandidateJobStatusEnded = "ex_staff"
)

var AllCandidateJobStatusEnded = []CandidateJobStatusEnded{
	CandidateJobStatusEndedHired,
	CandidateJobStatusEndedKiv,
	CandidateJobStatusEndedOfferLost,
	CandidateJobStatusEndedExStaff,
}

func (e CandidateJobStatusEnded) IsValid() bool {
	switch e {
	case CandidateJobStatusEndedHired, CandidateJobStatusEndedKiv, CandidateJobStatusEndedOfferLost, CandidateJobStatusEndedExStaff:
		return true
	}
	return false
}

func (e CandidateJobStatusEnded) String() string {
	return string(e)
}

func (e *CandidateJobStatusEnded) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CandidateJobStatusEnded(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CandidateJobStatusEnded", str)
	}
	return nil
}

func (e CandidateJobStatusEnded) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type CandidateJobStatusFailed string

const (
	CandidateJobStatusFailedOfferLost CandidateJobStatusFailed = "offer_lost"
	CandidateJobStatusFailedKiv       CandidateJobStatusFailed = "kiv"
)

var AllCandidateJobStatusFailed = []CandidateJobStatusFailed{
	CandidateJobStatusFailedOfferLost,
	CandidateJobStatusFailedKiv,
}

func (e CandidateJobStatusFailed) IsValid() bool {
	switch e {
	case CandidateJobStatusFailedOfferLost, CandidateJobStatusFailedKiv:
		return true
	}
	return false
}

func (e CandidateJobStatusFailed) String() string {
	return string(e)
}

func (e *CandidateJobStatusFailed) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CandidateJobStatusFailed(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CandidateJobStatusFailed", str)
	}
	return nil
}

func (e CandidateJobStatusFailed) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type CandidateJobStatusOpen string

const (
	CandidateJobStatusOpenApplied      CandidateJobStatusOpen = "applied"
	CandidateJobStatusOpenInterviewing CandidateJobStatusOpen = "interviewing"
	CandidateJobStatusOpenOffering     CandidateJobStatusOpen = "offering"
)

var AllCandidateJobStatusOpen = []CandidateJobStatusOpen{
	CandidateJobStatusOpenApplied,
	CandidateJobStatusOpenInterviewing,
	CandidateJobStatusOpenOffering,
}

func (e CandidateJobStatusOpen) IsValid() bool {
	switch e {
	case CandidateJobStatusOpenApplied, CandidateJobStatusOpenInterviewing, CandidateJobStatusOpenOffering:
		return true
	}
	return false
}

func (e CandidateJobStatusOpen) String() string {
	return string(e)
}

func (e *CandidateJobStatusOpen) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CandidateJobStatusOpen(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CandidateJobStatusOpen", str)
	}
	return nil
}

func (e CandidateJobStatusOpen) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type CandidateStatusEnum string

const (
	CandidateStatusEnumApplied      CandidateStatusEnum = "applied"
	CandidateStatusEnumInterviewing CandidateStatusEnum = "interviewing"
	CandidateStatusEnumOffering     CandidateStatusEnum = "offering"
	CandidateStatusEnumHired        CandidateStatusEnum = "hired"
	CandidateStatusEnumKiv          CandidateStatusEnum = "kiv"
	CandidateStatusEnumOfferLost    CandidateStatusEnum = "offer_lost"
	CandidateStatusEnumExStaff      CandidateStatusEnum = "ex_staff"
	CandidateStatusEnumNew          CandidateStatusEnum = "new"
)

var AllCandidateStatusEnum = []CandidateStatusEnum{
	CandidateStatusEnumApplied,
	CandidateStatusEnumInterviewing,
	CandidateStatusEnumOffering,
	CandidateStatusEnumHired,
	CandidateStatusEnumKiv,
	CandidateStatusEnumOfferLost,
	CandidateStatusEnumExStaff,
	CandidateStatusEnumNew,
}

func (e CandidateStatusEnum) IsValid() bool {
	switch e {
	case CandidateStatusEnumApplied, CandidateStatusEnumInterviewing, CandidateStatusEnumOffering, CandidateStatusEnumHired, CandidateStatusEnumKiv, CandidateStatusEnumOfferLost, CandidateStatusEnumExStaff, CandidateStatusEnumNew:
		return true
	}
	return false
}

func (e CandidateStatusEnum) String() string {
	return string(e)
}

func (e *CandidateStatusEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CandidateStatusEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CandidateStatusEnum", str)
	}
	return nil
}

func (e CandidateStatusEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type CurrencyEnum string

const (
	CurrencyEnumVnd CurrencyEnum = "vnd"
	CurrencyEnumUsd CurrencyEnum = "usd"
	CurrencyEnumJpy CurrencyEnum = "jpy"
)

var AllCurrencyEnum = []CurrencyEnum{
	CurrencyEnumVnd,
	CurrencyEnumUsd,
	CurrencyEnumJpy,
}

func (e CurrencyEnum) IsValid() bool {
	switch e {
	case CurrencyEnumVnd, CurrencyEnumUsd, CurrencyEnumJpy:
		return true
	}
	return false
}

func (e CurrencyEnum) String() string {
	return string(e)
}

func (e *CurrencyEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CurrencyEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CurrencyEnum", str)
	}
	return nil
}

func (e CurrencyEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type HiringJobOrderByAdditionalField string

const (
	HiringJobOrderByAdditionalFieldTotalCandidatesRecruited HiringJobOrderByAdditionalField = "total_candidates_recruited"
)

var AllHiringJobOrderByAdditionalField = []HiringJobOrderByAdditionalField{
	HiringJobOrderByAdditionalFieldTotalCandidatesRecruited,
}

func (e HiringJobOrderByAdditionalField) IsValid() bool {
	switch e {
	case HiringJobOrderByAdditionalFieldTotalCandidatesRecruited:
		return true
	}
	return false
}

func (e HiringJobOrderByAdditionalField) String() string {
	return string(e)
}

func (e *HiringJobOrderByAdditionalField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = HiringJobOrderByAdditionalField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid HiringJobOrderByAdditionalField", str)
	}
	return nil
}

func (e HiringJobOrderByAdditionalField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type HiringJobOrderByField string

const (
	HiringJobOrderByFieldName                     HiringJobOrderByField = "name"
	HiringJobOrderByFieldCreatedAt                HiringJobOrderByField = "created_at"
	HiringJobOrderByFieldAmount                   HiringJobOrderByField = "amount"
	HiringJobOrderByFieldSalaryFrom               HiringJobOrderByField = "salary_from"
	HiringJobOrderByFieldSalaryTo                 HiringJobOrderByField = "salary_to"
	HiringJobOrderByFieldLastApplyDate            HiringJobOrderByField = "last_apply_date"
	HiringJobOrderByFieldTotalCandidatesRecruited HiringJobOrderByField = "total_candidates_recruited"
	HiringJobOrderByFieldPriority                 HiringJobOrderByField = "priority"
)

var AllHiringJobOrderByField = []HiringJobOrderByField{
	HiringJobOrderByFieldName,
	HiringJobOrderByFieldCreatedAt,
	HiringJobOrderByFieldAmount,
	HiringJobOrderByFieldSalaryFrom,
	HiringJobOrderByFieldSalaryTo,
	HiringJobOrderByFieldLastApplyDate,
	HiringJobOrderByFieldTotalCandidatesRecruited,
	HiringJobOrderByFieldPriority,
}

func (e HiringJobOrderByField) IsValid() bool {
	switch e {
	case HiringJobOrderByFieldName, HiringJobOrderByFieldCreatedAt, HiringJobOrderByFieldAmount, HiringJobOrderByFieldSalaryFrom, HiringJobOrderByFieldSalaryTo, HiringJobOrderByFieldLastApplyDate, HiringJobOrderByFieldTotalCandidatesRecruited, HiringJobOrderByFieldPriority:
		return true
	}
	return false
}

func (e HiringJobOrderByField) String() string {
	return string(e)
}

func (e *HiringJobOrderByField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = HiringJobOrderByField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid HiringJobOrderByField", str)
	}
	return nil
}

func (e HiringJobOrderByField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type HiringJobStatus string

const (
	HiringJobStatusDraft  HiringJobStatus = "draft"
	HiringJobStatusOpened HiringJobStatus = "opened"
	HiringJobStatusClosed HiringJobStatus = "closed"
)

var AllHiringJobStatus = []HiringJobStatus{
	HiringJobStatusDraft,
	HiringJobStatusOpened,
	HiringJobStatusClosed,
}

func (e HiringJobStatus) IsValid() bool {
	switch e {
	case HiringJobStatusDraft, HiringJobStatusOpened, HiringJobStatusClosed:
		return true
	}
	return false
}

func (e HiringJobStatus) String() string {
	return string(e)
}

func (e *HiringJobStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = HiringJobStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid HiringJobStatus", str)
	}
	return nil
}

func (e HiringJobStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type I18nLanguage string

const (
	I18nLanguageEn I18nLanguage = "en"
	I18nLanguageVi I18nLanguage = "vi"
)

var AllI18nLanguage = []I18nLanguage{
	I18nLanguageEn,
	I18nLanguageVi,
}

func (e I18nLanguage) IsValid() bool {
	switch e {
	case I18nLanguageEn, I18nLanguageVi:
		return true
	}
	return false
}

func (e I18nLanguage) String() string {
	return string(e)
}

func (e *I18nLanguage) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = I18nLanguage(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid I18nLanguage", str)
	}
	return nil
}

func (e I18nLanguage) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type LocationEnum string

const (
	LocationEnumHaNoi     LocationEnum = "ha_noi"
	LocationEnumHoChiMinh LocationEnum = "ho_chi_minh"
	LocationEnumDaNang    LocationEnum = "da_nang"
	LocationEnumJapan     LocationEnum = "japan"
)

var AllLocationEnum = []LocationEnum{
	LocationEnumHaNoi,
	LocationEnumHoChiMinh,
	LocationEnumDaNang,
	LocationEnumJapan,
}

func (e LocationEnum) IsValid() bool {
	switch e {
	case LocationEnumHaNoi, LocationEnumHoChiMinh, LocationEnumDaNang, LocationEnumJapan:
		return true
	}
	return false
}

func (e LocationEnum) String() string {
	return string(e)
}

func (e *LocationEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = LocationEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid LocationEnum", str)
	}
	return nil
}

func (e LocationEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SalaryTypeEnum string

const (
	SalaryTypeEnumRange     SalaryTypeEnum = "range"
	SalaryTypeEnumUpTo      SalaryTypeEnum = "up_to"
	SalaryTypeEnumNegotiate SalaryTypeEnum = "negotiate"
	SalaryTypeEnumMinimum   SalaryTypeEnum = "minimum"
)

var AllSalaryTypeEnum = []SalaryTypeEnum{
	SalaryTypeEnumRange,
	SalaryTypeEnumUpTo,
	SalaryTypeEnumNegotiate,
	SalaryTypeEnumMinimum,
}

func (e SalaryTypeEnum) IsValid() bool {
	switch e {
	case SalaryTypeEnumRange, SalaryTypeEnumUpTo, SalaryTypeEnumNegotiate, SalaryTypeEnumMinimum:
		return true
	}
	return false
}

func (e SalaryTypeEnum) String() string {
	return string(e)
}

func (e *SalaryTypeEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SalaryTypeEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SalaryTypeEnum", str)
	}
	return nil
}

func (e SalaryTypeEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TeamOrderByAdditionalField string

const (
	TeamOrderByAdditionalFieldOpeningRequests TeamOrderByAdditionalField = "opening_requests"
	TeamOrderByAdditionalFieldNewestApplied   TeamOrderByAdditionalField = "newest_applied"
)

var AllTeamOrderByAdditionalField = []TeamOrderByAdditionalField{
	TeamOrderByAdditionalFieldOpeningRequests,
	TeamOrderByAdditionalFieldNewestApplied,
}

func (e TeamOrderByAdditionalField) IsValid() bool {
	switch e {
	case TeamOrderByAdditionalFieldOpeningRequests, TeamOrderByAdditionalFieldNewestApplied:
		return true
	}
	return false
}

func (e TeamOrderByAdditionalField) String() string {
	return string(e)
}

func (e *TeamOrderByAdditionalField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TeamOrderByAdditionalField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TeamOrderByAdditionalField", str)
	}
	return nil
}

func (e TeamOrderByAdditionalField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TeamOrderByField string

const (
	TeamOrderByFieldName            TeamOrderByField = "name"
	TeamOrderByFieldCreatedAt       TeamOrderByField = "created_at"
	TeamOrderByFieldOpeningRequests TeamOrderByField = "opening_requests"
	TeamOrderByFieldNewestApplied   TeamOrderByField = "newest_applied"
)

var AllTeamOrderByField = []TeamOrderByField{
	TeamOrderByFieldName,
	TeamOrderByFieldCreatedAt,
	TeamOrderByFieldOpeningRequests,
	TeamOrderByFieldNewestApplied,
}

func (e TeamOrderByField) IsValid() bool {
	switch e {
	case TeamOrderByFieldName, TeamOrderByFieldCreatedAt, TeamOrderByFieldOpeningRequests, TeamOrderByFieldNewestApplied:
		return true
	}
	return false
}

func (e TeamOrderByField) String() string {
	return string(e)
}

func (e *TeamOrderByField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TeamOrderByField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TeamOrderByField", str)
	}
	return nil
}

func (e TeamOrderByField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type UserStatus string

const (
	UserStatusActive   UserStatus = "active"
	UserStatusInactive UserStatus = "inactive"
)

var AllUserStatus = []UserStatus{
	UserStatusActive,
	UserStatusInactive,
}

func (e UserStatus) IsValid() bool {
	switch e {
	case UserStatusActive, UserStatusInactive:
		return true
	}
	return false
}

func (e UserStatus) String() string {
	return string(e)
}

func (e *UserStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserStatus", str)
	}
	return nil
}

func (e UserStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type AuditTrailAction string

const (
	AuditTrailActionCreate AuditTrailAction = "create"
	AuditTrailActionUpdate AuditTrailAction = "update"
	AuditTrailActionDelete AuditTrailAction = "delete"
)

var AllAuditTrailAction = []AuditTrailAction{
	AuditTrailActionCreate,
	AuditTrailActionUpdate,
	AuditTrailActionDelete,
}

func (e AuditTrailAction) IsValid() bool {
	switch e {
	case AuditTrailActionCreate, AuditTrailActionUpdate, AuditTrailActionDelete:
		return true
	}
	return false
}

func (e AuditTrailAction) String() string {
	return string(e)
}

func (e *AuditTrailAction) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AuditTrailAction(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid auditTrailAction", str)
	}
	return nil
}

func (e AuditTrailAction) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ProjectModule string

const (
	ProjectModuleTeams      ProjectModule = "teams"
	ProjectModuleHiringJobs ProjectModule = "hiring_jobs"
	ProjectModuleCandidates ProjectModule = "candidates"
)

var AllProjectModule = []ProjectModule{
	ProjectModuleTeams,
	ProjectModuleHiringJobs,
	ProjectModuleCandidates,
}

func (e ProjectModule) IsValid() bool {
	switch e {
	case ProjectModuleTeams, ProjectModuleHiringJobs, ProjectModuleCandidates:
		return true
	}
	return false
}

func (e ProjectModule) String() string {
	return string(e)
}

func (e *ProjectModule) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ProjectModule(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid projectModule", str)
	}
	return nil
}

func (e ProjectModule) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
