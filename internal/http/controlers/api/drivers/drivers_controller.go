package drivers

import (
	"encoding/json"
	"github.com/ShaleApps/{{SERVICE_NAME}}/internal/config"
	"net/http"
)

type Handler struct {
	SvcConfig *config.SvcConfig
}

func (h Handler) EmailDriverController(w http.ResponseWriter, r *http.Request) {
	// Controller logic here
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "OK",
		"message": "Load driver verified",
	})
}
