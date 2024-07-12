package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"rce-go/docker"
	"rce-go/utils"
	"strconv"
)

func main() {
	fmt.Println("start..")

	// upon starting pull all the required docker images
	docker.PullAllContainers()

	http.HandleFunc("POST /execute", func(w http.ResponseWriter, r *http.Request) {
		var data utils.PayLoad
		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, "Invalid payload", 400)
		}
		if !data.IsValid() {
			http.Error(w, "Unsupported language", 400)
		}
		data.Code = strconv.Quote(data.Code)
		fmt.Println(data.Code)

		fmt.Fprint(w, data.Language)

		out := docker.Run(&data)
		fmt.Println(string(out))
	})

	log.Fatal(http.ListenAndServe(":4000", nil))
	fmt.Println("server started on 4000")
}
