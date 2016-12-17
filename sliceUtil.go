package sliceUtil

func ContainsString(slice []string, target string) bool {
	for _, elem := range slice {
		if elem == target {
			return true
		}
	}
	return false
}
