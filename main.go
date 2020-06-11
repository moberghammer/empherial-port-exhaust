package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	if len(os.Args) == 1 {
		PrintUsageAndExit()
	}
	switch os.Args[1] {
	case "forward":
		CreateForwardServer(os.Args[2:])
	case "sleep":
		duration, err := time.ParseDuration(os.Args[2])
		if err != nil {
			panic(err)
		}
		CreateSleepServer(duration)
	default:
		PrintUsageAndExit()
	}
}

func PrintUsageAndExit() {
	fmt.Println("Usage: ")
	fmt.Println(" forward url")
	fmt.Println(" sleep 30s ")
	os.Exit(0)
}

func CreateSleepServer(sleepTime time.Duration) {
	fmt.Println(time.Now(), "Creating sleep server with sleep time: ", sleepTime)
	handlerSleep := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now(), "Starting to sleep for ", sleepTime)
		time.Sleep(sleepTime)
		fmt.Println(time.Now(), "Done sleeping for ", sleepTime)
	}
	http.HandleFunc("/", handlerSleep)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func CreateForwardServer(url []string) {
	fmt.Println(time.Now(), "Creating forward server with target url: ", url)
	handlerForward := func(w http.ResponseWriter, r *http.Request) {
		targetUrl := url[rand.Int()%len(url)]
		fmt.Println(time.Now(), " Requesting ", targetUrl)
		resp, err := http.Get(targetUrl)
		if err != nil {
			fmt.Println(time.Now(), " Request failed ", err)
		} else {
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(time.Now(), " Failed to read body ", err)
			} else {
				fmt.Println(time.Now(), " Request done ", body)
			}
			err = resp.Body.Close()
			if err != nil {
				fmt.Println(time.Now(), " Failed to close body")
			}
		}
	}
	http.HandleFunc("/", handlerForward)
	log.Fatal(http.ListenAndServe(":80", nil))
}
