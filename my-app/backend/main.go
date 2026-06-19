package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/ninja6734/FastVarLogger/modules/UDPConnectionCodec"
	"github.com/ninja6734/FastVarLogger/modules/networker"
)

type ConnectRequest struct {
	TestData string `json:"testData"`
	Port     int    `json:"port"`
}

func main() {
	r := gin.Default()

	// CORS für React-Frontend erlauben
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	}))

	// Routen
	r.GET("/api/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hallo vom Go-Backend!",
		})
	})

	r.POST("/api/connect", func(c *gin.Context) {
		var req ConnectRequest
		fmt.Println("Received connection request")
		fmt.Println("Request body: ", c.Request.Body)
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Ungültige Anfrage"})
			return
		}

		fmt.Println("Received connection request for IP:", req.TestData, "port:", req.Port)

		port := req.Port
		if port == 0 {
			port = 9999
		}
		testData := req.TestData

		go networker.ReadBuffer(port, UDPConnectionCodec.ReadUDPData, strings.Split(testData, ","))

		c.JSON(http.StatusOK, gin.H{"message": "Verbindung zum Roboter hergestellt"})
	})

	r.Run(":8080") // läuft auf http://localhost:8080
}
