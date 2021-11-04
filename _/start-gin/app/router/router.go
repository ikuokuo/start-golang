package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ikuokuo/start-golang/_/start-gin/app/api"
	"github.com/ikuokuo/start-golang/_/start-gin/app/router/middleware"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	r.Use(middleware.CorsMiddleware())
	r.Use(middleware.VersionMiddleware())

	r.GET("/albums", api.GetAlbums)
	r.POST("/albums", api.PostAlbums)

	r.GET("/albums/:id", api.GetAlbumByID)
	r.PUT("/albums/:id", api.UpdateAlbumByID)
	r.DELETE("/albums/:id", api.DeleteAlbumByID)

	r.POST("/files/upload", api.PostFiles)

	return r
}
