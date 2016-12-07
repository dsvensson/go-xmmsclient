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

// Flatten a {key: {source: value}} to {key: value} based on source preferences.
// Source preferences are defined as a list of strings that may end with '*' to
// match any suffix.
func PropDictToDict(inputValue XmmsDict, sourcePrefs []string) (XmmsDict, error) {
	result := XmmsDict{}
	for key, innerValue := range inputValue {
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

// Flatten a PropDict using default source preferences:
//  ["server", "client/*", "plugin/playlist", "plugin/segment",
//   "plugin/nibbler", "plugin/id3v2", "plugin/*", "*"]
func PropDictToDictDefault(input XmmsDict) (XmmsDict, error) {
	return PropDictToDict(input, defaultSourcePrefs)
}
