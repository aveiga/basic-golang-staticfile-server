package customlogger

import (
	"encoding/json"
	"fmt"

	"go.uber.org/zap"
)

func NewCustomLogger(appName string) *zap.SugaredLogger {
	rawJSON := []byte(fmt.Sprintf(`{
		"level": "debug",
		"encoding": "json",
		"initialFields": {
			"application": "%s"
		}
	}`, appName))

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
