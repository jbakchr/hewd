package rules

// first returns the first element of a slice or an empty string.
func first(list []string) string {
	if len(list) == 0 {
		return ""
	}
	return list[0]
}
