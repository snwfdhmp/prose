package prose

import (
	"regexp"
	"strings"
)

func RegexpWords(words ...string) *regexp.Regexp {
	var arr []string
	for _, w := range words {
		arr = append(arr, "\\b(?i:"+w+")")
	}
	regex := strings.Join(arr, "|")
	return regexp.MustCompile(regex)
}
