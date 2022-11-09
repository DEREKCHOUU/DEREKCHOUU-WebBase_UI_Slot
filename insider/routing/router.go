package routing

import (
	"github.com/gin-gonic/gin"

	"go/UI/insider/routing/middleware"
	_DB "go/UI/outsider/SQLDB/MYSQL"
	_UI "go/UI/outsider/frontend"
)

type ArticleHandler struct {
	AUsecase _DB.DeviceListSql
}

func Routersetup(router *gin.Engine) {
	middl := middleware.InitMiddleware()
	router.Use(middl.CORS())
	router.LoadHTMLGlob("./UI/html/*")
	router.Static("./UI/", "./")
	router.GET("/", _UI.UI)
	// if !key {
	// 	router.Use(MiddleDocker())
	// }
	setting := router.Group("/setting")
	setting.GET("/device")
	dock := setting.Group("/docker")
	dock.GET("/", docker)
	dock.GET("/ip", setip)
	dock.GET("/sql", setsql)
	dock.GET("/setDB", rebuildDB)
}
