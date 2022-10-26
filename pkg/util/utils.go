package util

func GetTargetType(name string, isdir bool) string {
	var res string
	if isdir {
		res = "dir"
		return res
	}
	for i := len(name) - 1; i >= 0; i-- {
		if name[i] == '.' {
			break
		}
		res = string(name[i]) + res
	}
	return res
}
