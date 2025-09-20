package logs

import (
	log "github.com/sirupsen/logrus"
)

func Debug(fields map[string]interface{}, msg string) {
	log.WithFields(fields).Debug(msg)
}

func Info(fields map[string]interface{}, msg string) {
	log.WithFields(fields).Info(msg)
}

func Warn(fields map[string]interface{}, msg string) {
	log.WithFields(fields).Warn(msg)
}

func Error(fields map[string]interface{}, msg string) {
	log.WithFields(fields).Error(msg)
}
