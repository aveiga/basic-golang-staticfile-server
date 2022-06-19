package customlogger

import (
	"encoding/json"
	"fmt"
	"os"

	"go.uber.org/zap"
)

func NewCustomLogger() *zap.SugaredLogger {
	if os.Getenv("APP_NAME") == "" {
		panic("APP_NAME env var is missing")
	}

	rawJSON := []byte(fmt.Sprintf(`{
		"level": "debug",
		"encoding": "json",
		"outputPaths": ["stdout"],
	  	"errorOutputPaths": ["stderr"],
		"initialFields": {
			"application": "%s"
		},
		"encoderConfig": {
			"messageKey": "message",
			"levelKey": "level",
			"timeKey": "timestamp",
			"nameKey": "name",
			"callerKey": "caller",
			"functionKey": "function",
			"stacktraceKey": "stacktrace",
			"levelEncoder": "lowercase",
			"timeEncoder": "iso8601",
			"callerEncoder": "short"
		}
	}`, os.Getenv("APP_NAME")))

	var cfg zap.Config
	if err := json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	sugar := logger.Sugar()

	return sugar
}
