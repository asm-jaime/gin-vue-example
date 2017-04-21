package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func TestGet(t *testing.T) { // {{{
	nums := 10
	config := Config{}
	config.SetDefault()
	vars := Vars{dataState: *NewDataState()}
	vars.dataState.SetRnd(10)

	router := NewRouter(&vars, &config)

	type Res struct {
		Msg   string `json:"msg"`
		Dates []Data `json:"body"`
	}
	res := Res{}
	// start make requests
	getData, err := http.NewRequest("GET", "/api/data", nil)
	wg := &sync.WaitGroup{}
	for count := 0; count < nums; count++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			response := httptest.NewRecorder()
			router.ServeHTTP(response, getData)
			err = json.Unmarshal(response.Body.Bytes(), &res)
			if len(res.Dates) < 1 {
				t.Error("error, empty get dates")
			}
		}()
	}
	if err != nil {
		t.Errorf("error put point: %v", err)
	}
	wg.Wait()
} // }}}

func TestPost(t *testing.T) { // {{{
	nums := 5
	config := Config{}
	config.SetDefault()
	vars := Vars{dataState: *NewDataState()}
	router := NewRouter(&vars, &config)

	// start make requests
	data := Data{}
	wg := &sync.WaitGroup{}
	for count := 0; count < nums; count++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			data.SetRnd()
			jdata, _ := json.Marshal(data)
			fmt.Println(string(jdata))
			postData, _ := http.NewRequest("POST", "/api/data", bytes.NewBuffer(jdata))
			postData.Header.Set("X-Custom-Header", "myvalue")
			postData.Header.Set("Content-Type", "application/json")
			response := httptest.NewRecorder()
			router.ServeHTTP(response, postData)
		}()
	}
	wg.Wait()

	// fmt.Println(vars.dataState.Dates)
	if len(vars.dataState.Dates) != nums {
		t.Error("error, count data does not overlap")
	}
} // }}}
