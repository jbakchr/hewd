package rules

func first(list []string) string {
	if len(list) == 0 {
		return ""
	}
	return list[0]
}
