package resp

import (
	"encoding/json"
	"net/http"

	"github.com/goji/pkg/logger"
)

func xJSONResponse(ctx http.ResponseWriter, response *APIResponse) {
	ctx.WriteHeader(response.Code)
	ctx.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(ctx).Encode(response); err != nil {
		logger.Error("json encoding error", err)
		http.Error(ctx, err.Error(), http.StatusInternalServerError)
	}
}
