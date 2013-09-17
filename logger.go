/*
	Simple logger package for go services supervised by DJB deamontools
*/

package logger

import (
	"fmt"
	"strings"
	"sync"
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

type logWritter struct {
	sync.Mutex
}

func (lw *logWritter) WriteLog(msg string) {
	lw.Lock()
	fmt.Println(msg)
	lw.Unlock()
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

var lw *logWritter

//var fout *bufio.Writer

func init() {
	lw = new(logWritter)
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
	msg := ""
	if level >= l.level {
		if l.addTimestamp {
			msg = fmt.Sprintf("%d - ", time.Now().UnixNano())
		}
		msg = fmt.Sprintf("%s%s - ", msg, strings.ToTitle(rlevels[level]))
		for i := range v {
			msg = fmt.Sprintf("%s%v", msg, v[i])
		}
		lw.WriteLog(msg)
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
