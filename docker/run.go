package docker

import (
	"fmt"
	"rce-go/utils"
)

func Run(data *utils.PayLoad) {
	var command string
	var image string
	switch data.Language {
	case "javascript":
		command = fmt.Sprintf("echo %s > index.js && node index.js", data.Code)
		image = utils.NodeImage
	case "python":
		command = "to be implemented"
		image = utils.PythonImage
	case "cpp":
		command = "to be implemented"
		image = "gcc:14"
	}

	fmt.Println(command)
	fmt.Println(image)

	StartContainer(command, image)
}
