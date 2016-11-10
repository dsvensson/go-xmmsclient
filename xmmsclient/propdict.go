package xmmsclient

import (
	"errors"
	"math"
	"strings"
)

var defaultSourcePrefs = []string{
	"server",
	"client/*",
	"plugin/playlist",
	"plugin/segment",
	"plugin/nibbler",
	"plugin/id3v2",
	"plugin/*",
	"*",
}

const (
	NoMatch int = math.MaxInt32
)

func match(sourcePrefs []string, source string) int {
	for index, pattern := range sourcePrefs {
		if strings.HasSuffix(pattern, "*") {
			if strings.HasPrefix(source, pattern[:len(pattern)-1]) {
				return index
			}
		} else if source == pattern {
			return index
		}
	}
	return NoMatch
}

func PropDictToDict(sourcePrefs []string, inputValue XmmsValue) (XmmsDict, error) {
	inputDict, ok := inputValue.(XmmsDict)
	if !ok {
		return nil, errors.New("Input not a XmmsDict->XmmsDict->XmmsValue")
	}

	result := XmmsDict{}
	for key, innerValue := range inputDict {
		innerDict, ok := innerValue.(XmmsDict)
		if !ok {
			return nil, errors.New("Input not a XmmsDict->XmmsDict->XmmsValue")
		}

		bestScore := NoMatch
		for source, value := range innerDict {
			score := match(sourcePrefs, source)
			if score < bestScore {
				result[key] = value
				bestScore = score
			}
		}
	}

	return result, nil
}

func PropDictToDictDefault(input XmmsValue) (XmmsDict, error) {
	return PropDictToDict(defaultSourcePrefs, input)
}
