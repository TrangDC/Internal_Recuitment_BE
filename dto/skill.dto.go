package dto

func FormatSkillField(input string) string {
	switch input {
	case "Name":
		return "model.skills.name"
	case "Description":
		return "model.skills.description"
	}
	return ""
}
