package docker

import (
	"fmt"
	"rce-go/utils"
)

func Run(data *utils.PayLoad)[]byte {
	var command string
	var image string
	switch data.Language {
	case "javascript":
		command = fmt.Sprintf("echo %s > index.js && node index.js", data.Code)
		image = utils.NodeImage
	case "python":
		command = fmt.Sprintf("echo %s > main.py && python main.py", data.Code)
		image = utils.PythonImage
	case "cpp":
		command = fmt.Sprintf("echo -e %s > main.cpp && g++ main.cpp -o main && ./main", data.Code)
		image = utils.CppImage
	case "c":
		command = fmt.Sprintf("echo -e %s > main.c && gcc main.c -o main && ./main", data.Code)
		image = utils.CImage
	}

	fmt.Println(command)
	fmt.Println(image)

	return StartContainer(command, image)
}
