package common

import "strconv"

const (
	API = "/api/xblog"
)

func ParseInt(n int) string {
	return strconv.Itoa(n)
}