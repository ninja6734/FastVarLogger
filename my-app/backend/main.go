package main

import (
	"net/http"

	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type ConnectRequest struct {
	IP string `json:"ip"`
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

		// Hier kannst du die Logik implementieren, um eine Verbindung zum Roboter herzustellen
		// Zum Beispiel: err := connectToRobot(req.IP)

		fmt.Println("Connecting to: ", req.IP)

		c.JSON(http.StatusOK, gin.H{"message": "Verbindung zum Roboter hergestellt"})
	})

	r.Run(":8080") // läuft auf http://localhost:8080
}
