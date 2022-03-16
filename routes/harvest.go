package routes

import (
	"FluxNodeStats/data"
	"FluxNodeStats/thirdapp"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func Health(c *gin.Context) {
	c.String(200, "{\"status\": \"We are Alive\"}")
}

func HarvestNodesInfo(c *gin.Context) {
	Nodes := data.Nodes{}
	Code, Json := thirdapp.GetNodesStats()
	if Code != 200 {
		c.String(Code, string(Json))
		return
	}
	err := json.Unmarshal(Json, &Nodes)
	if err != nil {
		c.String(500, fmt.Sprintf("%s error on GetNodesStats", err))
		return
	}
	log.Println(string(Json))
	c.String(200, string(Json))
}

func HarvestBlocksInfo(c *gin.Context) {
	FluxBlocsStats := data.FluxBlocsStats{}
	Code, Json := thirdapp.GetBlocStats()
	if Code != 200 {
		c.String(Code, string(Json))
		return
	}
	err := json.Unmarshal(Json, &FluxBlocsStats)
	if err != nil {
		c.String(500, fmt.Sprintf("%s error on HarvestBlocksInfo", err))
		return
	}
	log.Println(string(Json))
	c.String(200, string(Json))
}
