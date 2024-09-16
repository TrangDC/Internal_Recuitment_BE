package models

type I18nObject struct {
	Model struct {
		Candidates struct {
			ModelName   string `json:"model_name"`
			Name        string `json:"name"`
			Email       string `json:"email"`
			Phone       string `json:"phone"`
			Dob         string `json:"dob"`
			IsBlacklist string `json:"is_blacklist"`
		} `json:"candidates"`
	} `json:"model"`
	Excel struct {
		Id string `json:"id"`
	}
}

type I18n struct {
	En I18nObject
	Vi I18nObject
}

type I18nFormat struct {
	AuditTrail, Email string
}
