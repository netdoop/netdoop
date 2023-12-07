package utils

import (
	"regexp"
	"sync"
)

var regexpMapMutex sync.RWMutex
var regexpMap map[string]*regexp.Regexp

func init() {
	regexpMap = make(map[string]*regexp.Regexp)
}

func GetRegexp(pattern string) *regexp.Regexp {
	regexpMapMutex.Lock()
	defer regexpMapMutex.Unlock()
	v, ok := regexpMap[pattern]
	if ok {
		return v
	}

	v2 := regexp.MustCompile(pattern)
	regexpMap[pattern] = v2
	return v2
}
