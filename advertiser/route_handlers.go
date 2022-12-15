package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	helpers "github.com/sudheesh001/pcm-server-helpers"
	"net/http"
)

func indexHandler(context *gin.Context) {
	referer := context.Request.Referer()
	fmt.Printf("Index Handler: Referer: %v\n", referer)
	four := helpers.GenerateRandom4BitValue()
	six := helpers.GenerateRandom6BitValue()
	eight := helpers.GenerateRandom8BitValue()
	sixteen := helpers.GenerateRandom16BitValue()
	context.HTML(http.StatusOK, "index.html", gin.H{
		"four":    four,
		"six":     six,
		"eight":   eight,
		"sixteen": sixteen,
		"referer": referer,
	})
}

func triggerHandler(context *gin.Context) {
	referer := "https://pcm-target.gameautomators.com/"
	context.HTML(http.StatusOK, "trigger.html", gin.H{
		"referer": referer,
	})
}

func reportHandler(context *gin.Context) {
	fmt.Printf("Recieved report at Advertiser ...\n")
}
