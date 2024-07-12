package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rce-go/utils"
	"strconv"

	"github.com/docker/docker/client"
)



func main() {
	fmt.Println("start..")
	

	// upon starting pull all the required docker images
	// docker.PullAllContainers(ctx, dockerCli)

	http.HandleFunc("POST /execute", func(w http.ResponseWriter, r *http.Request) {
		var data utils.PayLoad
		err = json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Invalid payload", 400)
		}
		if data.IsValid() == false {
			http.Error(w, "Unsupported language", 400)
		}
		data.Code = strconv.Quote(data.Code)
		fmt.Println(data.Code)

		fmt.Fprint(w, data.Language)

		// docker.StartNodeContainer(ctx, dockerCli, string(content))
	})

	log.Fatal(http.ListenAndServe(":4000", nil))
	fmt.Println("server started on 4000")
}
