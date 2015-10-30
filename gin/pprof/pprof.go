package pprof

import (
	"github.com/gin-gonic/gin"
	"net/http/pprof"
)

func AttachRoutes(group *gin.RouterGroup) {
	group.GET("/debug/pprof/", gin.WrapF(pprof.Index))
	group.GET("/debug/pprof/heap", gin.WrapF(pprof.Index))
	group.GET("/debug/pprof/block", gin.WrapF(pprof.Index))
	group.GET("/debug/pprof/goroutine", gin.WrapF(pprof.Index))
	group.GET("/debug/pprof/threadcreate", gin.WrapF(pprof.Index))
	group.GET("/debug/pprof/cmdline", gin.WrapF(pprof.Cmdline))
	group.GET("/debug/pprof/profile", gin.WrapF(pprof.Profile))
	group.GET("/debug/pprof/symbol", gin.WrapF(pprof.Symbol))
}
