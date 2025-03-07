package main

import (
	"encoding/json"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type Server struct {
	router *http.ServeMux
	port   string
}

func ServeAPI(port string) *Server {
	server := &Server{
		router: http.NewServeMux(),
		port:   port,
	}

	// Serve static files
	server.router.Handle("/", http.FileServer(http.Dir("static")))
	server.router.Handle("/videos/", http.StripPrefix("/videos/", http.FileServer(http.Dir("videos"))))

	// API endpoints
	server.router.HandleFunc("/api/videos", server.handleListVideos())
	server.router.HandleFunc("/dowland", server.handleDownloadVideo())
	server.router.HandleFunc("/api/delete", server.handleDeleteVideo())

	return server
}

func (s *Server) Run() {
	log.Printf("Server starting on http://localhost%s\n", s.port)
	if err := http.ListenAndServe(s.port, s.router); err != nil {
		log.Fatal(err)
	}
}

func (s *Server) handleListVideos() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		videos, err := listVideoFiles()
		if err != nil {
			http.Error(w, "Error listing videos", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(videos)
	}
}

func listVideoFiles() ([]string, error) {
	var videos []string

	err := filepath.Walk("videos", func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && filepath.Ext(path) == ".mp4" {
			// Get just the filename
			videos = append(videos, filepath.Base(path))
		}
		return nil
	})

	if err != nil && !os.IsNotExist(err) {
		return nil, err
	}

	return videos, nil
}

func (s *Server) handleDownloadVideo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var body struct {
			VideoString string `json:"videoString"`
		}

		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		if body.VideoString == "" {
			http.Error(w, "Video URL is required", http.StatusBadRequest)
			return
		}

		getVideo(body.VideoString) // Ignore the return value since it's not used

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "success",
			"message": "Video downloaded successfully",
		})
	}
}

func (s *Server) handleDeleteVideo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		filename := r.URL.Query().Get("filename")
		if filename == "" {
			http.Error(w, "Filename is required", http.StatusBadRequest)
			return
		}

		err := os.Remove(filepath.Join("videos", filename))
		if err != nil {
			if os.IsNotExist(err) {
				http.Error(w, "Video not found", http.StatusNotFound)
			} else {
				http.Error(w, "Error deleting video", http.StatusInternalServerError)
			}
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "success",
			"message": "Video deleted successfully",
		})
	}
}
