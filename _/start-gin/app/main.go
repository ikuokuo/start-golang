package main

import (
	"github.com/ikuokuo/start-golang/_/start-gin/app/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
