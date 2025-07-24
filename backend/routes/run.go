package routes

import (
	"backend/sandbox"
	"encoding/json"
	"io"
	"net/http"
)

type RunRequest struct {
	Code string `json:"code"`
}

type RunResponse struct {
	Output string `json:"output"`
	Error  string `json:"error,omitempty"`
}

func RunHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only post allowed", http.StatusMethodNotAllowed)
		return
	}

	body, _ := io.ReadAll(r.Body)
	var req RunRequest
	json.Unmarshal(body, &req)

	output, err := sandbox.RunPythonCode(req.Code)
	var resp RunResponse
	if err != nil {
		resp.Error = err.Error()
	}
	resp.Output = output

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
