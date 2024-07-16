package api

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
)

type ResponseError struct {
	Error string `json:"error"`
}

func SendErr(ctx context.Context, w http.ResponseWriter, code int, err error) {
	l := ctx.Value("logger").(*slog.Logger)

	l.Error("api error", "error", err, "code", code)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	err = json.NewEncoder(w).Encode(ResponseError{Error: err.Error()})
	if err != nil {
		l.Error("encode response", "error", err)
	}
}

func SendJSON(ctx context.Context, w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		SendErr(ctx, w, http.StatusInternalServerError, err)
		return
	}
}
