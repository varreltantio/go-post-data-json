package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type Data struct {
	Water float64 `json:"water"`
	Wind  float64 `json:"wind"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/posts"

	for {
		// generate random water and wind values
		water := float64(rand.Intn(100-1) + 1)
		wind := float64(rand.Intn(100-1) + 1)

		// determine water and wind status
		waterStatus := ""
		if water < 5 {
			waterStatus = "aman"
		} else if water >= 5 && water <= 8 {
			waterStatus = "siaga"
		} else {
			waterStatus = "bahaya"
		}

		windStatus := ""
		if wind < 6 {
			windStatus = "aman"
		} else if wind >= 6 && wind <= 15 {
			windStatus = "siaga"
		} else {
			windStatus = "bahaya"
		}

		// create data object
		data := Data{
			Water: water,
			Wind:  wind,
		}

		// marshal data to json format
		payload, err := json.Marshal(data)
		if err != nil {
			fmt.Println("error:", err)
			return
		}

		// send post request
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json")
		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		// print result to terminal
		fmt.Println(string(payload))
		fmt.Printf("status water : %v \n", waterStatus)
		fmt.Printf("status wind : %v \n", windStatus)

		// wait for 15 seconds
		time.Sleep(15 * time.Second)
	}
}
