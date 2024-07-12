package utils

type PayLoad struct {
	Code     string `json:"code"`
	Language string `json:"language"`
	Input    string `json:"stdin"`
}

var supportedLangs = []string{
	"javascript",
	"python",
	"c",
	"cpp",
}

func (pl PayLoad) IsValid() bool {
	for _, lang := range supportedLangs {
		if lang == pl.Language {
			return true
		}
	}
	return false
}

// image names
const NodeImage = "node:20-alpine"
const PythonImage = "python:3.9.19-slim"
