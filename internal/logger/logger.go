package logger

import "log"

func New(prefix string) *log.Logger {
    return log.New(log.Writer(), prefix, log.LstdFlags)
}
