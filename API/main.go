package main

import (
	"go-bake/route"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(cors.Default())

	route.UserRouter(r)
	route.ProductRouter(r)

	r.Run(":1204")

}
