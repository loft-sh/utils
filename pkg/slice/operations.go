package slice

func Contains(haystack []string, needle string) bool {
	for _, s := range haystack {
		if s == "*" || s == needle {
			return true
		}
	}
	return false
}
