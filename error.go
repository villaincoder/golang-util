package util

import (
	"context"
	"google.golang.org/appengine/log"
	"runtime/debug"
)

func PrintErr(err error) {
	log.Debugf(context.Background(), "%v", err)
	debug.PrintStack()
}
