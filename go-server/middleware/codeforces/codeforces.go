package codeforces

import (
	"encoding/json"
	"fmt"
	"go-server/middleware"
	"io"
	"log"
	"net/http"
)

type apiResponse struct {
	Status   int
	Contests []contestData `json:"result"`
}

type contestData struct {
	Id                  int
	Name                string
	Type                string
	Phase               string
	Frozen              bool
	DurationSeconds     int64
	StartTimeSeconds    int64
	RelativeTimeSeconds int64
}

type Codeforces struct {
	Contests []middleware.Contest
}

func (s Codeforces) Fetch() []middleware.Contest {
	resp, err := http.Get("https://codeforces.com/api/contest.list")
	var contests []middleware.Contest
	if err != nil {
		fmt.Println("No response from API")
		return contests
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return contests
	}

	var apiRes apiResponse
	json.Unmarshal(body, &apiRes)

	for _, c := range apiRes.Contests {
		if c.Phase != "BEFORE" {
			break
		}
		p := middleware.Contest{
			Name:      c.Name,
			StartTime: c.StartTimeSeconds,
			Duration:  c.DurationSeconds,
		}
		contests = append(contests, p)
	}

	s.Contests = contests
	return contests
}
