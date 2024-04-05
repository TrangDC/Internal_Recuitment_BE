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

type AuthenticationToken struct {
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refreshToken"`
	TokenType    string    `json:"tokenType"`
	ExpiresAt    time.Time `json:"expiresAt"`
	Email        string    `json:"email"`
}

type Base64Response struct {
	Data *string `json:"data"`
}

type NewPreInput struct {
	StringInput string `json:"stringInput"`
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

type Pre struct {
	ID           string     `json:"id"`
	StringOutput string     `json:"stringOutput"`
	CreatedAt    time.Time  `json:"createdAt"`
	UpdatedAt    *time.Time `json:"updatedAt"`
	DeletedAt    *time.Time `json:"deletedAt"`
}

type PreEdge struct {
	Node   *Pre   `json:"node"`
	Cursor Cursor `json:"cursor"`
}

type PreFilter struct {
	StringInput *string    `json:"stringInput"`
	FromDate    *time.Time `json:"fromDate"`
	ToDate      *time.Time `json:"toDate"`
}

type PreFreeWord struct {
	StringInput *string `json:"stringInput"`
}

type PreGetAllResponse struct {
	Edges      []*PreEdge  `json:"edges"`
	Pagination *Pagination `json:"pagination"`
}

type PreOrder struct {
	Direction OrderDirection `json:"direction"`
	Field     PreOrderField  `json:"field"`
}

type PreResponse struct {
	Data *Pre `json:"data"`
}

type UpdatePreInput struct {
	StringInput string     `json:"stringInput"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
	DeletedAt   *time.Time `json:"deletedAt"`
}

type JSONFormat struct {
	Key   string `json:"key"`
	Value string `json:"value"`
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
	AttachmentFolderEmployee AttachmentFolder = "employee"
	AttachmentFolderContract AttachmentFolder = "contract"
)

var AllAttachmentFolder = []AttachmentFolder{
	AttachmentFolderEmployee,
	AttachmentFolderContract,
}

func (e AttachmentFolder) IsValid() bool {
	switch e {
	case AttachmentFolderEmployee, AttachmentFolderContract:
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

type AttendanceStatus string

const (
	AttendanceStatusConfirmed   AttendanceStatus = "confirmed"
	AttendanceStatusUnconfirmed AttendanceStatus = "unconfirmed"
	AttendanceStatusDiscarded   AttendanceStatus = "discarded"
)

var AllAttendanceStatus = []AttendanceStatus{
	AttendanceStatusConfirmed,
	AttendanceStatusUnconfirmed,
	AttendanceStatusDiscarded,
}

func (e AttendanceStatus) IsValid() bool {
	switch e {
	case AttendanceStatusConfirmed, AttendanceStatusUnconfirmed, AttendanceStatusDiscarded:
		return true
	}
	return false
}

func (e AttendanceStatus) String() string {
	return string(e)
}

func (e *AttendanceStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AttendanceStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid AttendanceStatus", str)
	}
	return nil
}

func (e AttendanceStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type CalculationType string

const (
	CalculationTypeAddition  CalculationType = "addition"
	CalculationTypeDeduction CalculationType = "deduction"
	CalculationTypeNone      CalculationType = "none"
)

var AllCalculationType = []CalculationType{
	CalculationTypeAddition,
	CalculationTypeDeduction,
	CalculationTypeNone,
}

func (e CalculationType) IsValid() bool {
	switch e {
	case CalculationTypeAddition, CalculationTypeDeduction, CalculationTypeNone:
		return true
	}
	return false
}

func (e CalculationType) String() string {
	return string(e)
}

func (e *CalculationType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CalculationType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CalculationType", str)
	}
	return nil
}

func (e CalculationType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type EmployeeDegree string

const (
	EmployeeDegreeVocational   EmployeeDegree = "vocational"
	EmployeeDegreeCollege      EmployeeDegree = "college"
	EmployeeDegreeUniversity   EmployeeDegree = "university"
	EmployeeDegreePostgraduate EmployeeDegree = "postgraduate"
	EmployeeDegreeMaster       EmployeeDegree = "master"
)

var AllEmployeeDegree = []EmployeeDegree{
	EmployeeDegreeVocational,
	EmployeeDegreeCollege,
	EmployeeDegreeUniversity,
	EmployeeDegreePostgraduate,
	EmployeeDegreeMaster,
}

func (e EmployeeDegree) IsValid() bool {
	switch e {
	case EmployeeDegreeVocational, EmployeeDegreeCollege, EmployeeDegreeUniversity, EmployeeDegreePostgraduate, EmployeeDegreeMaster:
		return true
	}
	return false
}

func (e EmployeeDegree) String() string {
	return string(e)
}

func (e *EmployeeDegree) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EmployeeDegree(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EmployeeDegree", str)
	}
	return nil
}

func (e EmployeeDegree) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type EntityType string

const (
	EntityTypeDepartment EntityType = "department"
	EntityTypeCompany    EntityType = "company"
)

var AllEntityType = []EntityType{
	EntityTypeDepartment,
	EntityTypeCompany,
}

func (e EntityType) IsValid() bool {
	switch e {
	case EntityTypeDepartment, EntityTypeCompany:
		return true
	}
	return false
}

func (e EntityType) String() string {
	return string(e)
}

func (e *EntityType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = EntityType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid EntityType", str)
	}
	return nil
}

func (e EntityType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type FilterOperator string

const (
	FilterOperatorEq          FilterOperator = "EQ"
	FilterOperatorNeq         FilterOperator = "NEQ"
	FilterOperatorGt          FilterOperator = "GT"
	FilterOperatorGte         FilterOperator = "GTE"
	FilterOperatorLt          FilterOperator = "LT"
	FilterOperatorLte         FilterOperator = "LTE"
	FilterOperatorIn          FilterOperator = "IN"
	FilterOperatorNin         FilterOperator = "NIN"
	FilterOperatorContains    FilterOperator = "CONTAINS"
	FilterOperatorNotContains FilterOperator = "NOT_CONTAINS"
	FilterOperatorStartWith   FilterOperator = "START_WITH"
	FilterOperatorEndWith     FilterOperator = "END_WITH"
)

var AllFilterOperator = []FilterOperator{
	FilterOperatorEq,
	FilterOperatorNeq,
	FilterOperatorGt,
	FilterOperatorGte,
	FilterOperatorLt,
	FilterOperatorLte,
	FilterOperatorIn,
	FilterOperatorNin,
	FilterOperatorContains,
	FilterOperatorNotContains,
	FilterOperatorStartWith,
	FilterOperatorEndWith,
}

func (e FilterOperator) IsValid() bool {
	switch e {
	case FilterOperatorEq, FilterOperatorNeq, FilterOperatorGt, FilterOperatorGte, FilterOperatorLt, FilterOperatorLte, FilterOperatorIn, FilterOperatorNin, FilterOperatorContains, FilterOperatorNotContains, FilterOperatorStartWith, FilterOperatorEndWith:
		return true
	}
	return false
}

func (e FilterOperator) String() string {
	return string(e)
}

func (e *FilterOperator) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FilterOperator(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FilterOperator", str)
	}
	return nil
}

func (e FilterOperator) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Gender string

const (
	GenderMale   Gender = "MALE"
	GenderFemale Gender = "FEMALE"
	GenderOther  Gender = "OTHER"
)

var AllGender = []Gender{
	GenderMale,
	GenderFemale,
	GenderOther,
}

func (e Gender) IsValid() bool {
	switch e {
	case GenderMale, GenderFemale, GenderOther:
		return true
	}
	return false
}

func (e Gender) String() string {
	return string(e)
}

func (e *Gender) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Gender(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid GENDER", str)
	}
	return nil
}

func (e Gender) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type IDNumberType string

const (
	IDNumberTypeIDNumber    IDNumberType = "ID_NUMBER"
	IDNumberTypeOldIDNumber IDNumberType = "OLD_ID_NUMBER"
	IDNumberTypePassportID  IDNumberType = "PASSPORT_ID"
)

var AllIDNumberType = []IDNumberType{
	IDNumberTypeIDNumber,
	IDNumberTypeOldIDNumber,
	IDNumberTypePassportID,
}

func (e IDNumberType) IsValid() bool {
	switch e {
	case IDNumberTypeIDNumber, IDNumberTypeOldIDNumber, IDNumberTypePassportID:
		return true
	}
	return false
}

func (e IDNumberType) String() string {
	return string(e)
}

func (e *IDNumberType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = IDNumberType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ID_NUMBER_TYPE", str)
	}
	return nil
}

func (e IDNumberType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Language string

const (
	LanguageVn Language = "vn"
	LanguageEn Language = "en"
)

var AllLanguage = []Language{
	LanguageVn,
	LanguageEn,
}

func (e Language) IsValid() bool {
	switch e {
	case LanguageVn, LanguageEn:
		return true
	}
	return false
}

func (e Language) String() string {
	return string(e)
}

func (e *Language) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Language(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Language", str)
	}
	return nil
}

func (e Language) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type MaritalStatus string

const (
	MaritalStatusSingle   MaritalStatus = "SINGLE"
	MaritalStatusMarried  MaritalStatus = "MARRIED"
	MaritalStatusDivorced MaritalStatus = "DIVORCED"
	MaritalStatusWidowed  MaritalStatus = "WIDOWED"
)

var AllMaritalStatus = []MaritalStatus{
	MaritalStatusSingle,
	MaritalStatusMarried,
	MaritalStatusDivorced,
	MaritalStatusWidowed,
}

func (e MaritalStatus) IsValid() bool {
	switch e {
	case MaritalStatusSingle, MaritalStatusMarried, MaritalStatusDivorced, MaritalStatusWidowed:
		return true
	}
	return false
}

func (e MaritalStatus) String() string {
	return string(e)
}

func (e *MaritalStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MaritalStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MARITAL_STATUS", str)
	}
	return nil
}

func (e MaritalStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PayrollTemplateFrequency string

const (
	PayrollTemplateFrequencyAnnually  PayrollTemplateFrequency = "annually"
	PayrollTemplateFrequencyMonthly   PayrollTemplateFrequency = "monthly"
	PayrollTemplateFrequencyQuarterly PayrollTemplateFrequency = "quarterly"
)

var AllPayrollTemplateFrequency = []PayrollTemplateFrequency{
	PayrollTemplateFrequencyAnnually,
	PayrollTemplateFrequencyMonthly,
	PayrollTemplateFrequencyQuarterly,
}

func (e PayrollTemplateFrequency) IsValid() bool {
	switch e {
	case PayrollTemplateFrequencyAnnually, PayrollTemplateFrequencyMonthly, PayrollTemplateFrequencyQuarterly:
		return true
	}
	return false
}

func (e PayrollTemplateFrequency) String() string {
	return string(e)
}

func (e *PayrollTemplateFrequency) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PayrollTemplateFrequency(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PayrollTemplateFrequency", str)
	}
	return nil
}

func (e PayrollTemplateFrequency) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type PreOrderField string

const (
	PreOrderFieldCreatedAt PreOrderField = "CREATED_AT"
)

var AllPreOrderField = []PreOrderField{
	PreOrderFieldCreatedAt,
}

func (e PreOrderField) IsValid() bool {
	switch e {
	case PreOrderFieldCreatedAt:
		return true
	}
	return false
}

func (e PreOrderField) String() string {
	return string(e)
}

func (e *PreOrderField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = PreOrderField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid PreOrderField", str)
	}
	return nil
}

func (e PreOrderField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Relationship string

const (
	RelationshipHusband     Relationship = "HUSBAND"
	RelationshipWife        Relationship = "WIFE"
	RelationshipBrother     Relationship = "BROTHER"
	RelationshipChild       Relationship = "CHILD"
	RelationshipGrandchild  Relationship = "GRANDCHILD"
	RelationshipFather      Relationship = "FATHER"
	RelationshipMother      Relationship = "MOTHER"
	RelationshipGrandfather Relationship = "GRANDFATHER"
	RelationshipGrandmother Relationship = "GRANDMOTHER"
	RelationshipOther       Relationship = "OTHER"
)

var AllRelationship = []Relationship{
	RelationshipHusband,
	RelationshipWife,
	RelationshipBrother,
	RelationshipChild,
	RelationshipGrandchild,
	RelationshipFather,
	RelationshipMother,
	RelationshipGrandfather,
	RelationshipGrandmother,
	RelationshipOther,
}

func (e Relationship) IsValid() bool {
	switch e {
	case RelationshipHusband, RelationshipWife, RelationshipBrother, RelationshipChild, RelationshipGrandchild, RelationshipFather, RelationshipMother, RelationshipGrandfather, RelationshipGrandmother, RelationshipOther:
		return true
	}
	return false
}

func (e Relationship) String() string {
	return string(e)
}

func (e *Relationship) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Relationship(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid RELATIONSHIP", str)
	}
	return nil
}

func (e Relationship) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type StatusEnum string

const (
	StatusEnumActive   StatusEnum = "active"
	StatusEnumInactive StatusEnum = "inactive"
)

var AllStatusEnum = []StatusEnum{
	StatusEnumActive,
	StatusEnumInactive,
}

func (e StatusEnum) IsValid() bool {
	switch e {
	case StatusEnumActive, StatusEnumInactive:
		return true
	}
	return false
}

func (e StatusEnum) String() string {
	return string(e)
}

func (e *StatusEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = StatusEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid STATUS_ENUM", str)
	}
	return nil
}

func (e StatusEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type SalaryComponentValueType string

const (
	SalaryComponentValueTypeFormula SalaryComponentValueType = "formula"
	SalaryComponentValueTypeNumber  SalaryComponentValueType = "number"
	SalaryComponentValueTypeText    SalaryComponentValueType = "text"
)

var AllSalaryComponentValueType = []SalaryComponentValueType{
	SalaryComponentValueTypeFormula,
	SalaryComponentValueTypeNumber,
	SalaryComponentValueTypeText,
}

func (e SalaryComponentValueType) IsValid() bool {
	switch e {
	case SalaryComponentValueTypeFormula, SalaryComponentValueTypeNumber, SalaryComponentValueTypeText:
		return true
	}
	return false
}

func (e SalaryComponentValueType) String() string {
	return string(e)
}

func (e *SalaryComponentValueType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = SalaryComponentValueType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid SalaryComponentValueType", str)
	}
	return nil
}

func (e SalaryComponentValueType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type StatusOfKeepSalary string

const (
	StatusOfKeepSalaryNoKeepSalary                    StatusOfKeepSalary = "no_keep_salary"
	StatusOfKeepSalaryMissingPersonalInformation      StatusOfKeepSalary = "missing_personal_information"
	StatusOfKeepSalaryNotSignedContract               StatusOfKeepSalary = "not_signed_contract"
	StatusOfKeepSalaryUncompletedResignationDocuments StatusOfKeepSalary = "uncompleted_resignation_documents"
)

var AllStatusOfKeepSalary = []StatusOfKeepSalary{
	StatusOfKeepSalaryNoKeepSalary,
	StatusOfKeepSalaryMissingPersonalInformation,
	StatusOfKeepSalaryNotSignedContract,
	StatusOfKeepSalaryUncompletedResignationDocuments,
}

func (e StatusOfKeepSalary) IsValid() bool {
	switch e {
	case StatusOfKeepSalaryNoKeepSalary, StatusOfKeepSalaryMissingPersonalInformation, StatusOfKeepSalaryNotSignedContract, StatusOfKeepSalaryUncompletedResignationDocuments:
		return true
	}
	return false
}

func (e StatusOfKeepSalary) String() string {
	return string(e)
}

func (e *StatusOfKeepSalary) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = StatusOfKeepSalary(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid StatusOfKeepSalary", str)
	}
	return nil
}

func (e StatusOfKeepSalary) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ValueType string

const (
	ValueTypePercentage ValueType = "PERCENTAGE"
	ValueTypeNumber     ValueType = "NUMBER"
)

var AllValueType = []ValueType{
	ValueTypePercentage,
	ValueTypeNumber,
}

func (e ValueType) IsValid() bool {
	switch e {
	case ValueTypePercentage, ValueTypeNumber:
		return true
	}
	return false
}

func (e ValueType) String() string {
	return string(e)
}

func (e *ValueType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ValueType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ValueType", str)
	}
	return nil
}

func (e ValueType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type AttachmentLabel string

const (
	AttachmentLabelAttachment AttachmentLabel = "attachment"
	AttachmentLabelContract   AttachmentLabel = "contract"
	AttachmentLabelOther      AttachmentLabel = "other"
)

var AllAttachmentLabel = []AttachmentLabel{
	AttachmentLabelAttachment,
	AttachmentLabelContract,
	AttachmentLabelOther,
}

func (e AttachmentLabel) IsValid() bool {
	switch e {
	case AttachmentLabelAttachment, AttachmentLabelContract, AttachmentLabelOther:
		return true
	}
	return false
}

func (e AttachmentLabel) String() string {
	return string(e)
}

func (e *AttachmentLabel) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = AttachmentLabel(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid attachmentLabel", str)
	}
	return nil
}

func (e AttachmentLabel) MarshalGQL(w io.Writer) {
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
	ProjectModuleEmployee            ProjectModule = "employee"
	ProjectModuleJobTitle            ProjectModule = "job_title"
	ProjectModuleEntity              ProjectModule = "entity"
	ProjectModuleEmployeeType        ProjectModule = "employee_type"
	ProjectModuleSkill               ProjectModule = "skill"
	ProjectModuleSkillType           ProjectModule = "skill_type"
	ProjectModuleContract            ProjectModule = "contract"
	ProjectModuleContractType        ProjectModule = "contract_type"
	ProjectModuleSalaryComponent     ProjectModule = "salary_component"
	ProjectModuleSalaryComponentType ProjectModule = "salary_component_type"
	ProjectModuleSalaryRank          ProjectModule = "salary_rank"
	ProjectModuleSalaryRankType      ProjectModule = "salary_rank_type"
	ProjectModuleSalaryTemplate      ProjectModule = "salary_template"
	ProjectModuleInsurance           ProjectModule = "insurance"
	ProjectModuleInsuranceType       ProjectModule = "insurance_type"
	ProjectModuleAttendance          ProjectModule = "attendance"
	ProjectModulePayrollTemplate     ProjectModule = "payroll_template"
	ProjectModuleCoefficientType     ProjectModule = "coefficient_type"
	ProjectModuleCoefficient         ProjectModule = "coefficient"
	ProjectModulePayroll             ProjectModule = "payroll"
	ProjectModuleEmployeePayslip     ProjectModule = "employee_payslip"
	ProjectModulePayslipTemplate     ProjectModule = "payslip_template"
)

var AllProjectModule = []ProjectModule{
	ProjectModuleEmployee,
	ProjectModuleJobTitle,
	ProjectModuleEntity,
	ProjectModuleEmployeeType,
	ProjectModuleSkill,
	ProjectModuleSkillType,
	ProjectModuleContract,
	ProjectModuleContractType,
	ProjectModuleSalaryComponent,
	ProjectModuleSalaryComponentType,
	ProjectModuleSalaryRank,
	ProjectModuleSalaryRankType,
	ProjectModuleSalaryTemplate,
	ProjectModuleInsurance,
	ProjectModuleInsuranceType,
	ProjectModuleAttendance,
	ProjectModulePayrollTemplate,
	ProjectModuleCoefficientType,
	ProjectModuleCoefficient,
	ProjectModulePayroll,
	ProjectModuleEmployeePayslip,
	ProjectModulePayslipTemplate,
}

func (e ProjectModule) IsValid() bool {
	switch e {
	case ProjectModuleEmployee, ProjectModuleJobTitle, ProjectModuleEntity, ProjectModuleEmployeeType, ProjectModuleSkill, ProjectModuleSkillType, ProjectModuleContract, ProjectModuleContractType, ProjectModuleSalaryComponent, ProjectModuleSalaryComponentType, ProjectModuleSalaryRank, ProjectModuleSalaryRankType, ProjectModuleSalaryTemplate, ProjectModuleInsurance, ProjectModuleInsuranceType, ProjectModuleAttendance, ProjectModulePayrollTemplate, ProjectModuleCoefficientType, ProjectModuleCoefficient, ProjectModulePayroll, ProjectModuleEmployeePayslip, ProjectModulePayslipTemplate:
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
