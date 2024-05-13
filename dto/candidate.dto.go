package dto

func FormatCandidateField(input string) string {
	switch input {
	case "Name":
		return "model.candidates.name"
	case "Email":
		return "model.candidates.email"
	case "Phone":
		return "model.candidates.phone"
	case "Dob":
		return "model.candidates.dob"
	case "IsBlacklist":
		return "model.candidates.is_blacklist"
	}
	return ""
}

func IsBlacklistI18n(input bool) string {
	switch input {
	case true:
		return "model.candidates.is_blacklist_enum.yes"
	default:
		return "model.candidates.is_blacklist_enum.no"
	}
}
