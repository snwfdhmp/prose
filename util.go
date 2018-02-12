package prose

import (
	"regexp"
	"strings"
)

//RegexpWords returns a regexp matching if one word in params is present in string (case insensitive)
func RegexpWords(words ...string) *regexp.Regexp {
	var arr []string
	for _, w := range words {
		arr = append(arr, "\\b(?i:"+w+")")
	}
	regex := strings.Join(arr, "|")
	return regexp.MustCompile(regex)
}

//newTrigger returns an empty 2D slice of pointer to Entity
func newTrigger() [][]*Entity {
	ts := make([][]*Entity, 0)
	for i := range ts {
		ts[i] = make([]*Entity, 0)
	}
	return ts
}
