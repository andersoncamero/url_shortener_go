package main

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

type URLShortener struct {
	urls  map[string]string
	mutex sync.RWMutex
}

type HomeResponse struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}

type ShortenResponse struct {
	ShortenURL string `json:"short_url"`
}

type StatsResponse struct {
	OriginalURL string `json:"original_url"`
	Clicks      int    `json:"clicks"`
}

type ShortenRequest struct {
	URL string `json:"url"`
}

func generateShortURL() string {
	bytes := make([]byte, 4)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func (u *URLShortener) HandleRedirection() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		shortUrl := r.URL.Path[1:]

		u.mutex.RLock()
		originalURL, ok := u.urls[shortUrl]
		u.mutex.RUnlock()

		if !ok {
			http.Error(w, "URL no encontrada", http.StatusNotFound)
			return
		}

		http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
	}
}

func (u *URLShortener) ShortenHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var req = ShortenRequest{}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(HomeResponse{
				Message: "Invalid request",
				Status:  false,
			})
			return
		}

		shortCode := generateShortURL()

		u.mutex.Lock()
		u.urls[shortCode] = req.URL
		u.mutex.Unlock()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(ShortenResponse{
			ShortenURL: "http://localhost:8080/" + shortCode,
		})
	}
}

func (u *URLShortener) StatusURLHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		click := 0
		shortUrl := r.URL.Path[1:]
		fmt.Println(shortUrl)

		u.mutex.RLock()
		for i, v := range u.urls {
			if u.urls[i] == v {
				fmt.Println("Found")
				click++
			}
		}
		u.mutex.RUnlock()

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(StatsResponse{
			OriginalURL: "http://localhost:8080/" + shortUrl,
			Clicks:      click,
		})
	}
}

func NewURLShortener() *URLShortener {
	return &URLShortener{
		urls: make(map[string]string),
	}
}

func main() {

	shortener := NewURLShortener()
	router := mux.NewRouter()

	router.HandleFunc("/{shortURL}", shortener.HandleRedirection()).Methods(http.MethodGet)
	router.HandleFunc("/stats/{shortURL}", shortener.StatusURLHandler()).Methods(http.MethodGet)
	router.HandleFunc("/shorten", shortener.ShortenHandler()).Methods(http.MethodPost)

	log.Println("Starting server on port", 8080)
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}
