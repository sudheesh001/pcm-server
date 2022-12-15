package main

import (
	"github.com/gin-gonic/gin"
	"net"
	"os"
	"strings"
)

const (
	Host = "HOST"
	Port = "PORT"

	DefaultHost = "localhost"
	DefaultPort = "9998"
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
	server.GET("/trigger", triggerHandler)
	server.POST("/.well-known/private-click-measurement/report-attribution/", reportHandler)

	if err := server.Run(net.JoinHostPort(hostname, port)); err != nil {
		panic("failed to run a server.")
	}
}
