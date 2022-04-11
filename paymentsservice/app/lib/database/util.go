package database

func RepeatIntArgs(args ...int) string {
	var res string
	for i := range args {
		if i != 0 {
			res += ","
		}
		res += "?"
	}
	return res
}
