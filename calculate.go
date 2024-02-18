package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type Task struct {
	Number            int       `json:"Number"`
	Status            string    `json:"Status"`
	N                 int       `json:"N"`
	D                 float64   `json:"D"`
	N1                float64   `json:"N1"`
	I                 float64   `json:"I"`
	TTL               float64   `json:"TTL"`
	Iteration         int       `json:"Iteration"`
	Result            float64   `json:"Result"`
	TaskTakeDatetime  time.Time `json:"TaskTakeDatetime"`
	TaskStartDatetime time.Time `json:"TaskStartDatetime"`
	TaskEndDatetime   time.Time `json:"TaskEndDatetime"`
}

var N = 8
var err error

var tokens = make(chan struct{}, N)
var TaskList = make([]Task, 0)

func calculate(w http.ResponseWriter, r *http.Request) {

	tokens <- struct{}{}

	var task Task
	task.N, err = strconv.Atoi(r.URL.Query().Get("n"))
	if err != nil {
		fmt.Println(err)
	}

	task.D, err = strconv.ParseFloat(r.URL.Query().Get("d"), 64)
	if err != nil {
		fmt.Println(err)
	}
	task.N1, err = strconv.ParseFloat(r.URL.Query().Get("n1"), 64)
	if err != nil {
		fmt.Println(err)
	}
	task.I, err = strconv.ParseFloat(r.URL.Query().Get("I"), 64)
	if err != nil {
		fmt.Println(err)
	}
	task.TTL, err = strconv.ParseFloat(r.URL.Query().Get("TTL"), 64)
	if err != nil {
		fmt.Println(err)
	}

	// for i := 0; i < task.n; i++ {
	// 	task.n1 += task.d
	// 	time.Sleep(time.Duration((task.I)*1000) * time.Millisecond)
	// }

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(task)
	if err != nil {
		fmt.Println(err)
	}
	<-tokens
}
