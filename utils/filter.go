package utils

const total = 3

func GetFilterElems() [total]string {
	var x [total]string = [total]string{"NodeID", "HTMLURL", "FollowersURL"}
	return x
}

func IsInFilter(x string) bool {
	for _, i := range GetFilterElems() {
		if i == x {
			return true
		}
	}
	return false
}
