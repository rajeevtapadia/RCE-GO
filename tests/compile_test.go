package docker

import (
	"rce-go/docker"
	"rce-go/utils"
	"strconv"
	"strings"
	"testing"
)

func TestCompileJS(t *testing.T) {
	var data utils.PayLoad
	data.Language = "javascript"
	data.Code = strconv.Quote("console.log(\"js\")")
	got := string(docker.Run(&data))
	want := "js"

	if got != want {
		t.Errorf("js test failed\n wanted: %s\n    got: %s", want, got)
	}
}

func TestCompilePython(t *testing.T) {
	var data utils.PayLoad
	data.Language = "python"
	data.Code = strconv.Quote(`print('python test is passing')`)

	got := docker.Run(&data)

	want := "python test is passing"

	gotStr := strings.TrimSpace(string(got))

	if gotStr != want {
		t.Errorf("python test failed\n wanted: %s\n got: %s", want, gotStr)
	}
}

func TestCompileCpp(t *testing.T) {
	var data utils.PayLoad
	data.Language = "cpp"
	data.Code = strconv.Quote(`#include <iostream>
int main() {
    std::cout << "cpp test is passing" << std::endl;
    return 0;
}`)

	got := docker.Run(&data)

	want := "cpp test is passing"

	gotStr := strings.TrimSpace(string(got))

	if gotStr != want {
		t.Errorf("cpp test failed\n wanted: %s\n got: %s", want, gotStr)
	}
}

func TestCompileC(t *testing.T) {
	var data utils.PayLoad
	data.Language = "c"
	data.Code = strconv.Quote(`#include <stdio.h>
int main() {
    printf("c test is passing");
    return 0;
}`)

	got := docker.Run(&data)

	want := "c test is passing"

	gotStr := strings.TrimSpace(string(got))

	if gotStr != want {
		t.Errorf("c test failed\n wanted: %s\n got: %s", want, gotStr)
	}
}
