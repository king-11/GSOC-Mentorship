package extractor

import (
	"strings"
)

func getGithubHandle(s string) string {
	s = strings.TrimSpace(s)
	return s[strings.LastIndex(s, "/")+1:]
}
