package converter

import "strconv"

func StrToInt(str string) (int, error) {
	return strconv.Atoi(str)
}
