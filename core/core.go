package core

import (
	"fmt"
	"time"
)

var (
	Version = "v0.0.0" // this will get set by -idflags at build time
	EST, _ = time.LoadLocation("America/New_York")
	CST, _ = time.LoadLocation("America/Chicago")
	UTC, _ = time.LoadLocation("UTC")
	YYYY_MM_DD = "2006-01-02"
	apiVersion = 1
	ApiBasePath = fmt.Sprintf("/api/v%d", apiVersion)
)

func Setup(configFile string) {
	SetupConfigFile(configFile)
	SetupDatabase(Cfg)
}