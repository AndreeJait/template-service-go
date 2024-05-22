package constant

import "strings"

type AppMode string

const (
	Development AppMode = "development"
	Production  AppMode = "production"
	Staging     AppMode = "staging"
)

const (
	FlagEnvMode      = "env"
	FlagShortEnvMode = "E"
)

var mapAppMode = map[string]AppMode{
	"dev":         Development,
	"development": Development,
	"d":           Development,

	"stag":    Staging,
	"staging": Staging,
	"s":       Staging,

	"prod":       Production,
	"production": Production,
	"p":          Production,
}

func GetAppMode(modeStr string) AppMode {
	if val, ok := mapAppMode[strings.ToLower(modeStr)]; ok {
		return val
	}
	return Development
}
