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
	apiBasePath = "api"
)

// Generates an API base path with the version appended (ex: /api/v1)
func ApiBasePath(version int) string {
	return fmt.Sprintf("/%s/v%d", apiBasePath, version)
}

func Setup(configFile string) {
	SetupConfigFile(configFile)
	SetupDatabase(Cfg)
}