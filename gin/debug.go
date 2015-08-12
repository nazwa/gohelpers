package gin

import (
	"bitbucket.org/tidepayments/gohelpers/gin/pprof"
	"bitbucket.org/tidepayments/gohelpers/gin/profiler"
	"github.com/gin-gonic/gin"
)

func AssignDebugHandlers(group *gin.RouterGroup) {
	profiler.AttachRoutes(group)
	pprof.AttachRoutes(group)
}
