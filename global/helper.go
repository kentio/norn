package global

// StringInSlice returns true if the string is in the slice.
func StringInSlice(s string, slice []string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}
