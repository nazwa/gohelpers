package pprof

import (
	"github.com/gin-gonic/gin"
	"net/http/pprof"
)

func AttachRoutes(group *gin.RouterGroup) {
	group.GET("/pprof/", gin.WrapF(pprof.Index))
	group.GET("/pprof/heap", gin.WrapF(pprof.Index))
	group.GET("/pprof/block", gin.WrapF(pprof.Index))
	group.GET("/pprof/goroutine", gin.WrapF(pprof.Index))
	group.GET("/pprof/threadcreate", gin.WrapF(pprof.Index))
	group.GET("/pprof/cmdline", gin.WrapF(pprof.Cmdline))
	group.GET("/pprof/profile", gin.WrapF(pprof.Profile))
	group.GET("/pprof/symbol", gin.WrapF(pprof.Symbol))
}
