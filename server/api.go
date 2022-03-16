package server

import (
	"FluxNodeStats/config"
	"FluxNodeStats/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func engine() *gin.Engine {
	gin.ForceConsoleColor()
	server := gin.New()
	server.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {

		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	server.Use(gin.Recovery())
	server.GET("/health", routes.Health)
	server.GET("/flux_nodes_data", routes.HarvestNodesInfo)
	server.GET("/flux_blocs_data", routes.HarvestBlocksInfo)
	server.GET("/calcul_nodes_rentability", routes.CalculNodesRentability)
	return server
}

func GoGinServer() {
	server := engine()
	server.Use(gin.Logger())
	serverUrl := fmt.Sprintf("%s:%d", config.Cfg.ApiAdress, config.Cfg.ApiPort)
	if err := engine().Run(serverUrl); err != nil {
		log.Fatal("Unable to start:", err)
	}
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
}
