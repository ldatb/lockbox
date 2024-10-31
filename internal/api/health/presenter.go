package health

import "time"

// HealthResponse represents the structure of the health check response.
type HealthResponse struct {
	// The status of the application (e.g., "ok", "degraded", "down")
	Status string `json:"status"`

	// Details of the application status. Not required
	Details string `json:"details"`

	// The timestamp when the health check was performed
	Timestamp time.Time `json:"timestamp"`
}
