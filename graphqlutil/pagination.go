package graphqlutil

import (
	"github.com/pkg/errors"
	"istudybookgitlab.hdzuoye.com/istudybook/server/golang-util.git"
)

func ParsePaginationFirst(args map[string]interface{}, def, max uint64) (first uint64) {
	first, ok := args["first"].(uint64)
	if !ok || first == 0 {
		first = def
	} else if first > max {
		first = max
	}
	return
}

func ParsePaginationUint64After(args map[string]interface{}) (after uint64, err error) {
	base64After, ok := args["after"].(string)
	if !ok {
		base64After = ""
	}
	if base64After != "" {
		after, err = util.Base64ToUint64(base64After)
		if err != nil {
			err = errors.WithStack(err)
			return
		}
	} else {
		after = 0
	}
	return
}

func ParsePaginationInt64After(args map[string]interface{}) (after int64, err error) {
	uint64After, err := ParsePaginationUint64After(args)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	after = int64(uint64After)
	return
}
