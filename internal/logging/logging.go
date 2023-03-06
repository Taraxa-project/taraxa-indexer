package logging

import (
	"encoding/json"
	"io"
	"os"
	"path/filepath"
	"reflect"

	log "github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Formatter struct {
	log.TextFormatter
}

func MakeFormatter() (formatter *Formatter) {
	formatter = new(Formatter)
	formatter.TextFormatter = log.TextFormatter{
		ForceColors:      true,
		DisableTimestamp: false,
		FullTimestamp:    true,
		TimestampFormat:  "2006-01-02T15:04:05",
	}

	return
}

// Whole this thing is needed to convert fields that are structs to json format
func (f *Formatter) Format(entry *log.Entry) ([]byte, error) {
	for k, v := range entry.Data {
		// skip error field
		if k == "error" {
			continue
		}
		val := reflect.ValueOf(v)
		if val.Kind() == reflect.Ptr {
			val = val.Elem()
		}
		if val.Kind() == reflect.Struct {
			json, _ := json.Marshal(v)
			entry.Data[k] = string(json)
		}
	}
	return f.TextFormatter.Format(entry)
}

// PanicLevel Level = iota
// // FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
// // logging level is set to Panic.
// FatalLevel
func levelFromString(l string) log.Level {
	switch l {
	case "trace":
		return log.TraceLevel
	case "debug":
		return log.DebugLevel
	case "info":
		return log.InfoLevel
	case "warn":
		return log.WarnLevel
	case "error":
		return log.ErrorLevel
	case "fatal":
		return log.FatalLevel
	default:
		log.WithField("level", l).Fatal("Incorrect log level passed. Could be: [trace, debug, info, warn, error, fatal]")
	}
	return log.InfoLevel
}

func Config(path, level string) {
	mw := io.MultiWriter(os.Stdout, &lumberjack.Logger{
		Filename:   filepath.Join(path, "main.log"),
		MaxSize:    100, // megabytes
		MaxBackups: 5,
		MaxAge:     2, //days
	})
	log.SetOutput(mw)
	log.SetLevel(levelFromString(level))
	log.SetFormatter(MakeFormatter())
}
