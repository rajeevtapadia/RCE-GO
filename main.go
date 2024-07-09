package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"


	"github.com/docker/docker/client"
	"rce-go/docker"
)

func main() {
	ctx := context.Background()
	dockerCli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer dockerCli.Close()

	// upon starting pull all the required docker images
	docker.PullAllContainers(ctx, dockerCli)

	http.HandleFunc("POST /execute", func(w http.ResponseWriter, r *http.Request) {
		content, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(content))
		// fmt.Fprint(w, string(content))

		docker.StartNodeContainer(ctx, dockerCli)
	})

	log.Fatal(http.ListenAndServe(":4000", nil))
}
