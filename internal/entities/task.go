package entities

import "time"

type Task struct {
	Number            int       `json:"Number"`
	Status            string    `json:"Status"`
	N                 int       `json:"n"`
	D                 float64   `json:"d"`
	N1                float64   `json:"n1"`
	I                 float64   `json:"I"`
	TTL               float64   `json:"TTL"`
	Iteration         int       `json:"Iteration"`
	Result            float64   `json:"Result"`
	TaskTakeDatetime  time.Time `json:"TaskTakeDatetime"`
	TaskStartDatetime time.Time `json:"TaskStartDatetime"`
	TaskEndDatetime   time.Time `json:"TaskEndDatetime"`
}
