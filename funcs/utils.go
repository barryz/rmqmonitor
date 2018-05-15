package funcs

import (
	"net/url"
)

func urlEncode(str string) string {
	return url.QueryEscape(str)
}
