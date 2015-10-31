package info

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
	"time"
)

type InfoStruct struct {
	CurrentTime time.Time
	Uptime      time.Duration
	GoVersion   string
}

var (
	startTime time.Time
)

func init() {
	startTime = time.Now()
}

func AttachRoutes(group *gin.RouterGroup) {
	group.GET("/debug/info", InfoHandler)
}

func InfoHandler(c *gin.Context) {
	info := InfoStruct{
		CurrentTime: time.Now(),
		Uptime:      time.Now().Sub(startTime),
		GoVersion:   runtime.Version(),
	}
	c.JSON(http.StatusOK, info)
}
