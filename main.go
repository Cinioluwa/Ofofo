package main // 1. Entry Point: Every Go program starts in package "main"

import (
	"net/http" // Standard HTTP library
	"time"     // Time library

	"github.com/gin-gonic/gin" // The web framework we just installed
)

// 2. The Data Model (Struct)
// This matches the Schema we designed.
// The `json:"..."` tags tell Go how to read/write JSON keys.
type Entity struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type SignalData struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

type SignalPayload struct {
	EventID    string     `json:"event_id"`
	Actor      Entity     `json:"actor"`
	Subject    Entity     `json:"subject"`
	Signal     SignalData `json:"signal"`
	OccurredAt int64      `json:"occurred_at"` // Unix timestamp (integer)
}

// 3. The Main Function (The Engine Start)
func main() {
	// Create a default Gin router
	r := gin.Default()

	// 4. Define a Route (POST /signals)
	// "c" is the Context (contains the Request and the Response)
	r.POST("/signals", func(c *gin.Context) {
		
		var payload SignalPayload

		// 5. Validation (The "Bind")
		// Try to read the JSON body into our `payload` struct.
		// If it fails (bad JSON), send a 400 error.
		if err := c.ShouldBindJSON(&payload); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 6. Logic (The "Brain" - MVP version)
		// For now, we just print it to the terminal so we know it worked.
		// Later, this is where we will push to Redis.
		println("------------------------------------------------")
		println("📥 RECEIVED SIGNAL:")
		println("   Event ID:", payload.EventID)
		println("   Actor:", payload.Actor.ID)
		println("   Subject:", payload.Subject.ID)
		println("   Signal:", payload.Signal.Name, payload.Signal.Value)
		println("------------------------------------------------")

		// 7. Response
		// Send back a 202 Accepted status.
		c.JSON(http.StatusAccepted, gin.H{
			"status":   "queued",
			"event_id": payload.EventID,
			"server_time": time.Now().Unix(),
		})
	})

	// 8. Run the Server on port 8080
	r.Run(":8080") 
}