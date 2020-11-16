package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
		FieldMap:      log.FieldMap{
			//FieldKeyTime:  "@timestamp",
			//FieldKeyLevel: "@level",
			//FieldKeyMsg:   "@message",
		},
	})
	log.WithFields(log.Fields{
		"animal": "walrus",
		"number": 1,
		"size":   10,
	}).Debugf("Debug - A walrus appears")
}
