package http

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/open-falcon/falcon-plus/modules/gateway/g"
	"github.com/toolkits/file"
)

func configCommonRoutes() {
	routes.GET("/health", func(c *gin.Context) {
		c.Writer.Write([]byte("ok\n"))
	})

	routes.GET("/version", func(c *gin.Context) {
		c.Writer.Write([]byte(fmt.Sprintf("%s\n", g.VERSION)))
	})

	routes.GET("/workdir", func(c *gin.Context) {
		c.Writer.Write([]byte(fmt.Sprintf("%s\n", file.SelfDir())))
	})

	routes.GET("/config", func(c *gin.Context) {
		RenderDataJson(c.Writer, g.Config())
	})

	routes.GET("/config/reload", func(c *gin.Context) {
		if strings.HasPrefix(c.Request.RemoteAddr, "127.0.0.1") {
			g.ParseConfig(g.ConfigFile)
			RenderDataJson(c.Writer, "ok")
		} else {
			RenderDataJson(c.Writer, "no privilege")
		}
	})
}
