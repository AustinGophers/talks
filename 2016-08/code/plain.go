package main

import (
	"os"
	"runtime/trace"
	"time"
)

func main() {
	trace.Start(os.Stderr)
	// create new channel of type int
	go func() {
		// send 42 to channel
		time.Sleep(10 * time.Millisecond)
	}()
	trace.Stop()
}
