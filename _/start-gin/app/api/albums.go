package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ikuokuo/start-golang/_/start-gin/app/entity"
	"github.com/ikuokuo/start-golang/_/start-gin/app/service"
)

// Get all albums
func GetAlbums(c *gin.Context) {
	albums := make([]entity.Album, 0, len(service.DB))
	for _, a := range service.DB {
		albums = append(albums, a.(entity.Album))
	}
	if len(albums) > 0 {
		c.JSON(http.StatusOK, albums)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "albums are empty"})
	}
}

// Create a new album
func PostAlbums(c *gin.Context) {
	var album entity.Album
	if err := c.ShouldBind(&album); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	album.ID = uuid.New().String()
	album.CreatedAt = time.Now()
	album.UpdatedAt = album.CreatedAt

	service.DB[album.ID] = album

	c.JSON(http.StatusCreated, album)
}

// Get an album by its id
func GetAlbumByID(c *gin.Context) {
	id := c.Param("id")
	if a, ok := service.DB[id]; ok {
		c.JSON(http.StatusOK, a)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "album not found"})
	}
}

// Update an album by its id
func UpdateAlbumByID(c *gin.Context) {
	id := c.Param("id")
	if a, ok := service.DB[id]; ok {
		var album entity.Album
		if err := c.ShouldBind(&album); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		albumNew := a.(entity.Album)
		albumNew.Title = album.Title
		albumNew.Artist = album.Artist
		albumNew.Price = album.Price
		albumNew.UpdatedAt = time.Now()
		service.DB[id] = albumNew
		c.JSON(http.StatusOK, albumNew)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "album not found"})
	}
}

// Delete an album by its id
func DeleteAlbumByID(c *gin.Context) {
	id := c.Param("id")
	if a, ok := service.DB[id]; ok {
		delete(service.DB, id)
		c.JSON(http.StatusNoContent, a)
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "album not found"})
	}
}
