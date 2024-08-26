package main

import (
	"github.com/kchxng/go-oracle-api/infra"
	"github.com/kchxng/go-oracle-api/infra/db"
)

// @title DDSP core API
// @version 1.0
// @description DDSP API - core
// @contact.name API Support
// @contact.email fiber@swagger.io
// @host localhost:4000
// @BasePath /

// @securityDefinitions.apiKey bearerAuth
// @in header
// @name Authorization
// @description Enter token with `Bearer: ` prefix, e.g. "Bearer abcde1234".
func main() {

	infra.InitConfig()
	infra.InitTimezone()

	db := db.InitDatabase()
	_ = db
	// port, _ := strconv.ParseInt(infra.GetENV("app.port"), 10, 64)

	// _ = infrastructure_setup.AppRun(int(port), AppRoutes, db)

}
