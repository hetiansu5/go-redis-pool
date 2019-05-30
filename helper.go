package go_redis_pool

import "strings"

func isOKString(str string, err error) (bool, error) {
	if strings.ToUpper(str) == "OK" {
		return true, err
	}
	return false, err
}
