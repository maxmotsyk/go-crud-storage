package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func SetLogger(appEnv string) {

	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{
		TimestampFormat: "02-01-2006 15:04:05",
	})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	if appEnv == "" {
		appEnv = "prod"
	}

	switch appEnv {
	case "dev":
		log.SetLevel(log.DebugLevel)
	case "prod":
		log.SetLevel(log.InfoLevel)
	default:
		log.SetLevel(log.WarnLevel)
	}

	// Only log the warning severity or above.

}
