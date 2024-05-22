package config

import "github.com/AndreeJait/go-utility/configw"

const (
	Development configw.ConfigMode = "development"
	Staging     configw.ConfigMode = "staging"
	Production  configw.ConfigMode = "production"
)

var MapFiles = configw.LocationMap{
	Development: "files/config/development/config.yaml",
	Production:  "files/config/production/config.yaml",
	Staging:     "files/config/staging/config.yaml",
}

var ConfigMode = map[string]configw.ConfigMode{
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
