package user

func GetAttributeValueOrEmpty(value string) string {
	if value == "" {
		return "null"
	}
	return value
}
