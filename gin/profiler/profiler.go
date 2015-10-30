package profiler

import (
	"github.com/gin-gonic/gin"
	"github.com/nazwa/profiler"
)

func AttachRoutes(group *gin.RouterGroup) {
	group.GET("/debug/profiler", gin.WrapF(profiler.MemStatsHTMLHandler))
	group.GET("/debug/profiler/info", gin.WrapF(profiler.ProfilingInfoJSONHandler))
	group.GET("/debug/profiler/start", gin.WrapF(profiler.StartProfilingHandler))
	group.GET("/debug/profiler/stop", gin.WrapF(profiler.StopProfilingHandler))
}
