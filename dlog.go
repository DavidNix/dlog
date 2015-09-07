// Dlog provides a simple wrapper around the standard log package to conditionally print debug log statements.
// Many other logging packages offer too many configuration options or log levels.
// Dlog just has one: SetDebug(true or false)
//
package dlog

import (
	"log"
	"sync"
)

var (
	Fatal     = log.Fatal
	Fatalf    = log.Fatalf
	Fatalln   = log.Fatalln
	Flags     = log.Flags
	Panic     = log.Panic
	Panicf    = log.Panicf
	Prefix    = log.Prefix
	Print     = log.Print
	Printf    = log.Printf
	Println   = log.Println
	SetFlags  = log.SetFlags
	SetOutput = log.SetOutput
)

var (
	Ldate         = log.Ldate
	Ltime         = log.Ltime
	Lmicroseconds = log.Lmicroseconds
	Llongfile     = log.Llongfile
	Lshortfile    = log.Lshortfile
	LstdFlags     = log.LstdFlags
)

var (
	debug = false
	mu    = &sync.Mutex{}
)

func init() {
	SetFlags(0)
}

// SetPrefix delegates to standard log package SetPrefix and adds "[ ]" around the prefix
func SetPrefix(s string) {
	if s != "" {
		log.SetPrefix("[" + s + "] ")
	} else {
		log.SetPrefix("")
	}
}

// PanicIf delegates to standard log Panic iff error is present
func PanicIf(err error) {
	if err != nil {
		log.Panic(err)
	}
}

// FatalIf delegates to standard log Fatal iff error is present
func FatalIf(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// IsDebug returns whether the underlying logger is set to debug mode
func IsDebug() bool {
	mu.Lock()
	defer mu.Unlock()
	return debug
}

// SetDebug turns debug on or off
func SetDebug(d bool) {
	mu.Lock()
	defer mu.Unlock()
	debug = d
}

// Debug prints string if debug is set
// Prepends "[DEBUG]" to output
func Debug(s string) {
	if IsDebug() {
		Print(debugPrefix(s))
	}
}

// Debugf prints format string if debug is set
// Prepends "[DEBUG]" to output
func Debugf(format string, v ...interface{}) {
	if IsDebug() {
		Printf(debugPrefix(format), v...)
	}
}

// Debugln prints lines if debug is set
// Prepends "[DEBUG]" to output
func Debugln(v ...interface{}) {
	if IsDebug() {
		v = append([]interface{}{"[DEBUG]"}, v...)
		log.Println(v...)
	}
}

func debugPrefix(s string) string {
	return "[DEBUG] " + s
}
