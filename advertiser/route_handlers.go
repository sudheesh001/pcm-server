package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	helpers "github.com/sudheesh001/pcm-server-helpers"
	"log"
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
	})
}

func trackingPixelHandler(context *gin.Context) {
	triggerAttribution := helpers.GenerateRandom4BitValue()
	priority := helpers.GenerateRandom6BitValue()
	log.Printf("Trigger Data: %v\n Priority: %v\n", triggerAttribution, priority)
	referer := context.Request.Referer()
	fmt.Printf("Referer: %v\n", referer)
	context.Redirect(
		http.StatusMovedPermanently,
		fmt.Sprintf("http://localhost:9999/.well-known/private-click-measurement/trigger-attribution/%v/%v", triggerAttribution, priority),
	)
}

func reportHandler(context *gin.Context) {
	fmt.Printf("Recieved report at Advertiser ...\n")
}
