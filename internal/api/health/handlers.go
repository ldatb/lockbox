package health

import (
	"net/http"
	"time"

	"gitlab.com/xrs-cloud/lockbox/core/internal/global"
	"gitlab.com/xrs-cloud/lockbox/core/internal/utils"
)

// basicHealthCheck responds with a status for Kubernetes liveness and readiness probes
// This is a common endpoint used by Kubernetes for liveness and readiness probes.
// It ensures the service is running properly and ready to handle requests.
func basicHealthCheck(w http.ResponseWriter, r *http.Request) {
	// Create a simple response with a status of "ok"
	response := HealthResponse{
		Status:    "ok",
		Details:   "none",
		Timestamp: time.Now(),
	}

	// Encode the response and return
	utils.WriteJSONResponse(w, http.StatusOK, response)
}

// detailedHealthCheck handles the detailed health check request.
// This endpoint is used for a more detailed health check.
// It checks the connection to the database
func detailedHealthCheck(w http.ResponseWriter, r *http.Request) {
	// Status is ok by default and details is null
	httpStatus := http.StatusOK
	status := "ok"
	details := "none"

	// Perform database connectivity test
	sqlDB, err := global.Database.DB()
	if err != nil {
		httpStatus = http.StatusInternalServerError
		status = "degraded"
		details = "Failed to get database object"
	}

	err = sqlDB.Ping()
	if err != nil {
		httpStatus = http.StatusInternalServerError
		status = "degraded"
		details = "Database connection failed"
	}

	// Create response
	response := HealthResponse{
		Status:    status,
		Details:   details,
		Timestamp: time.Now(),
	}

	// Encode the response and return
	utils.WriteJSONResponse(w, httpStatus, response)
}
