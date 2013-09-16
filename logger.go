/*
	Simple logger package for go services supervised by DJB deamontools
*/

package logger

import (
	"fmt"
	"strings"
	"time"
)

const (
	CRITICAL = 4
	ERROR    = 3
	WARNING  = 2
	NOTICE   = 1
	DEBUG    = 0
)

type Logger struct {
	level        int
	addTimestamp bool
}

var levels = map[string]int{
	"critical": CRITICAL,
	"error":    ERROR,
	"warning":  WARNING,
	"notice":   NOTICE,
	"debug":    DEBUG,
}

var rlevels = [5]string{
	"debug",
	"notice",
	"warning",
	"error",
	"critical",
}

func init() {

}

func (l *Logger) SetLevel(level string) {
	level = strings.ToLower(level)
	l.level = levels[level]
}

func (l *Logger) SetTimeStamp(ts bool) {
	l.addTimestamp = ts
}

// Base logger
func (l *Logger) log(level int, v ...interface{}) {
	if level >= l.level {
		if l.addTimestamp {
			fmt.Print(fmt.Sprintf("%d - ", time.Now().UnixNano()))
		}
		fmt.Print(fmt.Sprintf("%s - ", strings.ToTitle(rlevels[level])))
		fmt.Println(v...)
	}
}

func (l *Logger) Debug(v ...interface{}) {
	l.log(DEBUG, v...)
}

func (l *Logger) Notice(v ...interface{}) {
	l.log(NOTICE, v...)
}

func (l *Logger) Warning(v ...interface{}) {
	l.log(WARNING, v...)
}

func (l *Logger) Error(v ...interface{}) {
	l.log(ERROR, v...)
}

func (l *Logger) Critical(v ...interface{}) {
	l.log(CRITICAL, v...)
}
