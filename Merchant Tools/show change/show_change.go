package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

//jsonChange is the struct that defines the json response from the raida
type jsonChange struct {
	Server  string   `json:"server"`
	Status  string   `json:"status"`
	Owner   string   `json:"owner"`
	D1      []string `json:"d1"`
	D5      []string `json:"d5"`
	D25     []string `json:"d25"`
	D100    []string `json:"d100"`
	Message string   `json:"message"`
	Version string   `json:"version"`
	Time    string   `json:"time"`
}

func main() {
	denomination := 5
	change := showChange("2", denomination)
	fmt.Println(change)
}

func showChange(sn string, denom int, server ...int) [4][]string {
	rand.Seed(time.Now().UTC().UnixNano())
	changeArray := [4][]string{}
	serverArray := [3]int{}
	if len(server) > 0 {
		serverArray[0] = server[0]
	} else {
		serverArray[0] = rand.Intn(24)
	}

	if len(server) > 1 {
		serverArray[1] = server[1]
	} else {
		serverArray[1] = rand.Intn(24)
	}

	if len(server) > 2 {
		server[2] = server[2]
	} else {
		serverArray[2] = rand.Intn(24)
	}

	for serverArray[0] == serverArray[1] {
		serverArray[1] = rand.Intn(24)

		for serverArray[1] == serverArray[2] {
			serverArray[2] = rand.Intn(24)

			for serverArray[2] == serverArray[0] {
				serverArray[0] = rand.Intn(24)
			}

		}

	}

	fmt.Printf("%v", serverArray)
	done := make(chan [2]string)
	//var changeURL string

	for i := 0; i < 3; i++ {
		go sendRequest(done, denom, sn, serverArray[i])
		fmt.Printf("sent request %d\r\n", i)
	}

	var change []jsonChange

	for i := 0; i < 3; i++ {
		results := <-done
		fmt.Printf("time Elapsed for request %d: %v\r\n", i+1, results[1])
		change = append(change, readChange(results[0]))

	}

	var d1 []string
	var d5 []string
	var d25 []string
	var d100 []string

	if len(change[0].D1) > 0 || len(change[1].D1) > 0 {
		d1 = intersection(change[0].D1, change[1].D1)
		d1 = intersection(change[2].D1, d1)
	} else {
		d1 = change[0].D1
	}

	if len(change[0].D5) > 0 || len(change[1].D5) > 0 {
		d5 = intersection(change[0].D5, change[1].D5)
		d5 = intersection(change[2].D5, d5)
	} else {
		d5 = change[0].D5
	}

	if len(change[0].D25) > 0 || len(change[1].D25) > 0 {
		d25 = intersection(change[0].D25, change[1].D25)
		d25 = intersection(change[2].D25, d25)
	} else {
		d25 = change[0].D25
	}

	if len(change[0].D100) > 0 || len(change[1].D100) > 0 {
		d100 = intersection(change[0].D100, change[1].D100)
		d100 = intersection(change[2].D100, d100)
	} else {
		d100 = change[0].D100
	}

	changeArray[0] = d1
	changeArray[1] = d5
	changeArray[2] = d25
	changeArray[3] = d100

	return changeArray
}

//send the Request for the change
func sendRequest(done chan [2]string, denom int, sn string, serverID int) {
	var responseText [2]string

	changeURL := fmt.Sprintf("https://raida%d.cloudcoin.global/service/show_change?sn=%s&denomination=%d", serverID, sn, denom)

	start := time.Now()
	response, _ := http.Get(changeURL)

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	//	Request := fmt.Sprintf(" RAIDA %d: %v", serverID, changeURL)
	elapsed := time.Since(start)
	elapsedString := fmt.Sprintf("%v", elapsed)
	responseText[0] = string(body)
	responseText[1] = elapsedString
	done <- responseText
}

//read the change and return useable data
func readChange(returnResponse string) jsonChange {
	raw := []byte(returnResponse)
	var change jsonChange
	_ = json.Unmarshal(raw, &change)
	return change
}

//get the intersection of two Arrays
func intersection(a, b []string) (c []string) {
	m := make(map[string]bool)

	for _, item := range a {
		m[item] = true
	}

	for _, item := range b {
		if _, ok := m[item]; ok {
			c = append(c, item)
		}
	}
	return
}
