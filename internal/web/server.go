package web

import (
	"encoding/json"
	"log"

	"faceit-parser/service"
	"net/http"
)

type request struct {
	Usernames []string `json:"usernames"`
}

type response struct {
	Valid   []string `json:"valid"`
	Invalid []string `json:"invalid"`
}

func (s *Server) parseByNicknameHandler(w http.ResponseWriter, r *http.Request) {

	var req *request

	req.Usernames = make([]string, 0)
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var resp *response

	resp.Valid = make([]string, 0)
	resp.Invalid = make([]string, 0)

	for _, username := range req.Usernames {
		service.GetFaceitByLink()
	}

	data, err := json.Marshal(resp)
	if err != nil {
		log.Println("failed to marshal: ", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if _, err = w.Write(data); err != nil {
		log.Println("failed to write: ", err)
		http.Error(w, "failed to write", http.StatusInternalServerError)
		return
	}
}
