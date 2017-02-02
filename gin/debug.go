package gin

import (
	"bitbucket.org/tidepayments/gohelpers/gin/info"
	"bitbucket.org/tidepayments/gohelpers/gin/pprof"
	"bitbucket.org/tidepayments/gohelpers/gin/profiler"
	"gopkg.in/gin-gonic/gin.v1"
)

func AssignDebugHandlers(group *gin.RouterGroup) {
	profiler.AttachRoutes(group)
	pprof.AttachRoutes(group)
	info.AttachRoutes(group)

}
