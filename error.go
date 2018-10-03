package util

import (
	"log"
	"runtime/debug"
)

func PrintError(err error) {
	log.Printf("[Error]:%v", err)
	debug.PrintStack()
}
