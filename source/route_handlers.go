package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	helpers "github.com/sudheesh001/pcm-server-helpers"
	"net/http"
)

func indexHandler(context *gin.Context) {
	four := helpers.GenerateRandom4BitValue()
	six := helpers.GenerateRandom6BitValue()
	eight := helpers.GenerateRandom8BitValue()
	sixteen := helpers.GenerateRandom16BitValue()
	context.HTML(http.StatusOK, "index.html", gin.H{
		"four":    four,
		"six":     six,
		"eight":   eight,
		"sixteen": sixteen,
	})
}

func reportHandler(context *gin.Context) {
	fmt.Printf("Received report at Source ...\n")
}
