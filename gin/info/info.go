package info

import (
	"net/http"
	"runtime"
	"time"

	"github.com/gin-gonic/gin"
)

type InfoStruct struct {
	CurrentTime time.Time
	Uptime      string
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
		Uptime:      time.Now().Sub(startTime).String(),
		GoVersion:   runtime.Version(),
	}
	c.JSON(http.StatusOK, info)
}
