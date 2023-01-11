package logger

import "time"

type APILogData struct {
	IP       string `json:"ip"`
	Endpoint string `json:"endpoint"`
	//RequestBody  string    `json:"requestBody"`
	//ResponseBody string    `json:"responseBody"`
	StatusCode int       `json:"statusCode"`
	ErrMessage *string   `json:"errMessage"`
	Took       int64     `json:"took"`
	StartAt    time.Time `json:"startAt"`
	CreatedAt  time.Time `json:"createdAt"`
}
