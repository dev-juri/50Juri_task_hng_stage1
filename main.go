package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

type Response struct {
	SlackName     string    `json:"slack_name"`
	CurrentDay    string    `json:"current_day"`
	UtcTime       time.Time `json:"utc_time"`
	Track         string    `json:"track"`
	GithubFileUrl string    `json:"github_file_url"`
	GithubRepoUrl string    `json:"github_repo_url"`
	StatusCode    int       `json:"status_code"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api", apiQueryHandler)
	http.ListenAndServe(":443", mux)
}

func apiQueryHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		u, err := url.Parse(r.URL.String())

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		params := u.Query()
		slackName, track := params.Get("slack_name"), params.Get("track")

		if len(slackName) < 1 || len(track) < 1 {
			http.Error(w, "slack_name and track cannot be empty", http.StatusBadRequest)
			return
		}

		response := Response{
			SlackName:     slackName,
			CurrentDay:    time.Now().UTC().Weekday().String(),
			UtcTime:       time.Now().UTC(),
			Track:         track,
			GithubFileUrl: "https://github.com/dev-juri/50Juri_task_hng_stage1/blob/main/main.go",
			GithubRepoUrl: "https://github.com/dev-juri/50Juri_task_hng_stage1",
			StatusCode:    http.StatusOK,
		}

		js, err := json.Marshal(response)
		js = append(js, '\n')

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(js)
	} else {
		http.Error(w, "Invalid request", http.StatusMethodNotAllowed)
		return
	}
}
