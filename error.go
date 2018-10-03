package util

import (
	"context"
	"google.golang.org/appengine/log"
	"runtime/debug"
)

func PrintError(err error) {
	log.Debugf(context.Background(), "%v", err)
	debug.PrintStack()
}
