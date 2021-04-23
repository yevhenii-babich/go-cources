package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

var serverStarted bool

func runServer() {
	if serverStarted {
		return
	}
	go func() {
		serverStarted = true
		main()
		serverStarted = false
	}()
}

func TestGet(t *testing.T) {
	runServer()
	client := http.Client{Timeout: 15 * time.Second}
	respose, err := client.Get("http://localhost:8080/todos/")
	if err != nil {
		t.Fatal(err)
		return
	}

	dest := make([]Todo, 0)

	err = json.NewDecoder(respose.Body).Decode(&dest)
	if err != nil {
		t.Fatal(err)
		return
	}

	fmt.Printf("%+v \n", dest)
}

func TestPost(t *testing.T) {
	runServer()
	client := http.Client{Timeout: 15 * time.Second}
	request := &Todo{
		Name: "Try Go",
		Done: false,
	}
	rawData, err := json.Marshal(request)
	if err != nil {
		t.Fatal(err)
		return
	}

	req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/todos/",
		bytes.NewBuffer(rawData))
	if err != nil {
		t.Fatal(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)

	res, err := ioutil.ReadAll(resp.Body)
	str := string(res)
	fmt.Println(str)
}
