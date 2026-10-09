package main

import (
	"fmt"
	"github.com/wsxiaoys/terminal/color"
	"os"
)

type Logger struct {
	errors []interface{}
	infos  []interface{}
	warns  []interface{}
}

func formatMsg(colorTag string, msgs ...interface{}) []interface{} {
	if len(msgs) > 0 {
		if s, ok := msgs[0].(string); ok {
			msgs[0] = colorTag + s
		} else {
			msgs = append([]interface{}{colorTag}, msgs...)
		}
	} else {
		msgs = []interface{}{colorTag}
	}
	return msgs
}

func (logger *Logger) Error(msgs ...interface{}) {
	logger.Errors(1, msgs...)
}

func (logger *Logger) Errors(status int, msgs ...interface{}) {
	msgs = formatMsg("@r", msgs...)
	if isTestEnv() {
		logger.errors = append(logger.errors, msgs...)
		return
	}

	color.Print(msgs...)
	os.Exit(status)
}

func (logger *Logger) Warn(msgs ...interface{}) {
	msgs = formatMsg("@y", msgs...)
	if isTestEnv() {
		logger.warns = append(logger.warns, msgs...)
		return
	}

	color.Print(msgs...)
}

func (logger *Logger) Info(msgs ...interface{}) {
	msgs = formatMsg("@c", msgs...)
	if isTestEnv() {
		logger.infos = append(logger.infos, msgs...)
		return
	}

	color.Print(msgs...)
}

func (logger *Logger) Errorln(msgs ...interface{}) {
	logger.Errorsln(1, msgs...)
}

func (logger *Logger) Errorsln(status int, msgs ...interface{}) {
	logger.Errors(status, append(msgs, "\n")...)
}

func (logger *Logger) Warnln(msgs ...interface{}) {
	logger.Warn(append(msgs, "\n")...)
}

func (logger *Logger) Infoln(msgs ...interface{}) {
	logger.Info(append(msgs, "\n")...)
}

func (logger *Logger) clear() {
	logger.infos = logger.infos[:0]
	logger.warns = logger.warns[:0]
	logger.errors = logger.errors[:0]
}

func (logger *Logger) flush() {
	fmt.Println(logger.infos)
	fmt.Println(logger.warns)
	fmt.Println(logger.errors)
	logger.clear()
}

var logger = new(Logger)
