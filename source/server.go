package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"os"
	"strings"
)

const (
	Host = "HOST"
	Port = "PORT"

	DefaultHost = "localhost"
	DefaultPort = "9999"
)

func ParseOSEnvData() (string, string) {
	host := os.Getenv(Host)
	port := os.Getenv(Port)

	// sanitize
	host = strings.TrimSpace(host)
	port = strings.TrimSpace(port)

	// Check and assign defaults
	if host == "" || len(host) == 0 {
		host = DefaultHost
	}
	if port == "" || len(port) == 0 {
		port = DefaultPort
	}

	return host, port
}

func main() {
	hostname, port := ParseOSEnvData()

	server := gin.Default()
	server.Static("/static/", "assets/")
	server.LoadHTMLGlob("templates/*")

	server.GET("/", indexHandler)
	server.GET("/.well-known/private-click-measurement/trigger-attribution/:triggerData/:priority", attributionHandler)
	server.GET("/.well-known/private-click-measurement/trigger-attribution/:triggerData", attributionHandler)
	server.POST("/.well-known/private-click-measurement/report-attribution/", reportHandler)

	if err := server.Run(net.JoinHostPort(hostname, port)); err != nil {
		panic("failed to run a server.")
	}
}

func attributionHandler(context *gin.Context) {
	advertiser := context.Request.Header.Get("Referer")
	triggerData := context.Param("triggerData")
	priority := context.Param("priority")
	fmt.Printf("Received trigger data from advertiser: %v with trigger data: %v and priority %v\n", advertiser, triggerData, priority)
	context.Data(http.StatusOK, "image/png", []byte{})
}
