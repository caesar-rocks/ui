package toast

import (
	"encoding/json"

	caesar "github.com/caesar-rocks/core"
)

const (
	INFO    = "info"
	SUCCESS = "success"
	WARNING = "warning"
	DANGER  = "danger"
)

type Toast struct {
	Level   string `json:"level"`
	Message string `json:"message"`
}

func New(level string, message string) Toast {
	return Toast{level, message}
}

func Info(ctx *caesar.Context, message string) {
	New(INFO, message).SetHXTriggerHeader(ctx)
}

func Success(ctx *caesar.Context, message string) {
	New(SUCCESS, message).SetHXTriggerHeader(ctx)
}

func Warning(ctx *caesar.Context, message string) {
	New(WARNING, message).SetHXTriggerHeader(ctx)
}

func Danger(ctx *caesar.Context, message string) {
	New(DANGER, message).SetHXTriggerHeader(ctx)
}

func (t Toast) SetHXTriggerHeader(ctx *caesar.Context) {
	eventMap := map[string]Toast{}
	eventMap["toast"] = t
	jsonData, err := json.Marshal(eventMap)
	if err != nil {
		return
	}

	ctx.SetHeader("HX-Trigger", string(jsonData))
}
