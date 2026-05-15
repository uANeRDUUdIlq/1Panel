package main

import (
	"fmt"
	"os"

	"github.com/1Panel-dev/1Panel/backend/app/api/v1"
	"github.com/1Panel-dev/1Panel/backend/global"
	"github.com/1Panel-dev/1Panel/backend/init/boot"
	"github.com/1Panel-dev/1Panel/backend/router"
	"github.com/gin-gonic/gin"
)

// @title 1Panel API
// @version 1.0
// @description 1Panel backend API server
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url https://github.com/1Panel-dev/1Panel

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:9999
// @BasePath /api/v1
func main() {
	// Initialize core components: config, database, logger, etc.
	boot.Init()

	if global.CONF.System.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// Initialize router
	routerManager := router.NewRouter()
	engine := routerManager.Init()

	// Register API v1 routes
	publicGroup := engine.Group(global.CONF.System.BaseDir)
	{
		v1.InitBaseRouter(publicGroup)
	}

	privateGroup := engine.Group(global.CONF.System.BaseDir)
	{
		v1.InitRouter(privateGroup)
	}
	_ = privateGroup

	serverAddr := fmt.Sprintf("%s:%d",
		global.CONF.System.BindAddress,
		global.CONF.System.Port,
	)

	global.LOG.Infof("1Panel server is starting on %s", serverAddr)

	if err := engine.Run(serverAddr); err != nil {
		global.LOG.Errorf("Failed to start server: %v", err)
		// Print to stderr as well so the error is visible even if the logger fails
		fmt.Fprintf(os.Stderr, "Fatal: failed to start server: %v\n", err)
		os.Exit(1)
	}
}
