package main

import (
	"context"
	"fmt"
	docker "github.com/fsouza/go-dockerclient"
	"go/types"
	"log"
	"net/http"
)

func main() {
	createMetricsServer()
}

func createMetricsServer() {
	handlerMetrics := func(w http.ResponseWriter, r *http.Request) {
		containerPids := getContainerpids()
		for _, pid := range containerPids {
			fmt.Println("Container pid: ", pid)
		}
	}
	http.HandleFunc("/metrics", handlerMetrics)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func getContainerpids() []int {
	client, err := docker.NewClientFromEnv()
	if err != nil {
		log.Fatalln("Failed to create docker client", err)
	}
	containers, err := client.ContainerList(context.Background(), types.ContainerListOptions{})
	return comtaiers

}
