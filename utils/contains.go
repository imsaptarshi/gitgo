package utils

func Contains(x string, v []string) bool {
	for _, i := range v {
		if x == i {
			return true
		}
	}
	return false
}
