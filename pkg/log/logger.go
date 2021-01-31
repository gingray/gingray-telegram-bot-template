package log

import (
	"encoding/json"
	"go.uber.org/zap"
)

func Info(msg string, a ...interface{}) {
	sugar := zap.NewExample().Sugar()
	defer sugar.Sync()
	sugar.Infow(msg, a)
}

func InfoJson(msg string, jsonStr string) {
	sugar := zap.NewExample().Sugar()
	raw := json.RawMessage(jsonStr)
	defer sugar.Sync()
	sugar.Infow(msg,"payload", &raw)
}


