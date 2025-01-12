package main

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServe struct {
	listenAddr string
}

func ServeAPI(listenAddr string) *APIServe {
	return &APIServe{
		listenAddr: listenAddr,
	}
}

func (s *APIServe) Run() {

	router := mux.NewRouter()

	router.HandleFunc("/video", httpFuncHandler(s.DowlandVideo))
	http.ListenAndServe(s.listenAddr, router)

}

func httpFuncHandler(f APIFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}

func (s *APIServe) DowlandVideo(w http.ResponseWriter, r *http.Request) error {

	if r.Method != "POST" {
		return WriteJSON(w, http.StatusMethodNotAllowed, "WRONG METHOD")
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {

		return WriteJSON(w, http.StatusBadRequest, "Unable to read request body")
	}
	defer r.Body.Close()

	var videoStr VideoURLBody
	if err := json.Unmarshal(body, &videoStr); err != nil {
		return WriteJSON(w, http.StatusBadRequest, "Invalid JSON format")
	}

	getVideo(videoStr.VideoString)

	return WriteJSON(w,http.StatusOK,"Video Saved.")
}
