package main

import (
	"github.com/gin-gonic/gin"
	"go-bake/route"
)

func main() {
	r := gin.Default()

	route.UserRouter(r)

	r.Run(":1204")

}
