package sidekick

// InStrings checks if a value exists in an string slice
func InStrings(s string, ss []string) bool {
	for _, v := range ss {
		if s == v {
			return true
		}
	}
	return false
}
