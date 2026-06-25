package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestUDPCodec(t *testing.T) {
	r := gin.Default()

	// CORS für React-Frontend erlauben
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	}))

	r.POST("/api/test-udp-codec", func(c *gin.Context) {
		var req ConnectRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Ungültige Anfrage"})
			return
		}

		testData := []byte(req.TestData)
		packet := CreateUDPDataPacket(testData)

		fmt.Println("Created UDP packet:", packet)

		decodedData, err := DecodeUDPData(packet)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Fehler beim Dekodieren"})
			return
		}

		fmt.Println("Decoded data:", string(decodedData))

		c.JSON(http.StatusOK, gin.H{"message": "Test erfolgreich", "decodedData": string(decodedData)})
	})

	r.Run(":8081")
}
