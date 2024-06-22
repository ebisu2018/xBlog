package common

import "strconv"

const (
	API = "/api/xblog/v1"
)

func ParseInt(n int) string {
	return strconv.Itoa(n)
}