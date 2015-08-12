package profiler

import (
	"github.com/gin-gonic/gin"
	"github.com/nazwa/profiler"
)

func AttachRoutes(group *gin.RouterGroup) {
	group.GET("/profiler", gin.WrapF(profiler.MemStatsHTMLHandler))
	group.GET("/profiler/info", gin.WrapF(profiler.ProfilingInfoJSONHandler))
	group.GET("/profiler/start", gin.WrapF(profiler.StartProfilingHandler))
	group.GET("/profiler/stop", gin.WrapF(profiler.StopProfilingHandler))
}
