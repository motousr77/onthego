package main

import (
	"net/http"
	"onthego/common"

	"github.com/gin-gonic/gin"
)

func main() {
	var node = common.SomeNode{Node: nil, Value: "Adam"}
	node.Print()

	var rt = gin.Default()
	rt.GET("/check", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "OK")
	})
	rt.Run("localhost:8000")
}


/*
curl https://checkip.amazonaws.com/
*/