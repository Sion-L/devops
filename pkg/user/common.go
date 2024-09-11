package user

func GetAttributeValueOrEmpty(value string) string {
	if value == "" {
		return "11111111111"
	}
	return value
}
