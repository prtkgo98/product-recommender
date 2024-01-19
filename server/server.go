package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"../recommendation"
)

type Server struct {
	RecommendationSystem *recommendation.RecommendationSystem
}

func NewServer() *Server {
	return &Server{
		RecommendationSystem: recommendation.NewRecommendationSystem(),
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/recommendations", s.handleRecommendations)

	log.Println("Server started on http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return err
	}

	return nil
}

var handleRecommendationsRequestBody struct {
	UserID    int `json:"user_id"`
	Page      int `json:"page"`
	PageCount int `json:"page_count"`
}

func (s *Server) handleRecommendations(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &handleRecommendationsRequestBody)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	userID := handleRecommendationsRequestBody.UserID
	page := handleRecommendationsRequestBody.Page
	count := handleRecommendationsRequestBody.PageCount

	recommendations, err := s.RecommendationSystem.GetRecommendations(userID, count, page)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(recommendations)
	if err != nil {
		http.Error(w, "Failed to encode recommendations", http.StatusInternalServerError)
		return
	}
}
