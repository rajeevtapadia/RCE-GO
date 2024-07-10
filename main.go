package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"rce-go/docker"

	"github.com/docker/docker/client"
)

func main() {
	fmt.Println("start..")
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
		fmt.Println(strconv.Quote(string(content)))
		// fmt.Fprint(w, string(content))

		docker.StartNodeContainer(ctx, dockerCli, string(content))
	})

	log.Fatal(http.ListenAndServe(":4000", nil))
	fmt.Println("server started on 4000")
}
