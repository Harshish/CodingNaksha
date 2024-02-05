package middleware

import (
	"encoding/json"
	"net/http"
)

var PlatformMap map[string]ContestInterface

func GetAllContests(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		return
	}
	//w.Write([]byte("Getting All Contests"))
	var res Response
	for k, v := range PlatformMap {
		var cpp ContestsPerPlatform
		cpp.Platform = k
		cpp.Contests = v.Fetch()
		res.Result = append(res.Result, cpp)
	}

	json.NewEncoder(w).Encode(res)
}
