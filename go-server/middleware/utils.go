package middleware

type Response struct {
	Code   int                   `json:"code"`
	Result []ContestsPerPlatform `json:"result"`
}

type ContestsPerPlatform struct {
	Platform string    `json:"platform"`
	Contests []Contest `json:"contests"`
}

type Contest struct {
	Name      string `json:"name"`
	StartTime int64  `json:"startTime"`
	Duration  int64  `json:"duration"`
}

type ContestInterface interface {
	Fetch() []Contest
}
